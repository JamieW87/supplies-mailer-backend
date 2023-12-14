package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun/migrate"

	"one-stop/internal/config"
	"one-stop/internal/service"
	"one-stop/internal/store/postgres"
	http2 "one-stop/internal/transport/inbound/http"
	http_out "one-stop/internal/transport/outbound/http"
	"one-stop/migrations"
)

func main() {
	ctx := context.Background()
	log := logrus.New()

	log.Info("Gathering config")
	env, err := config.Get()
	if err != nil {
		log.Error("Could not retrieve environment", err)
		panic(err)
	}

	db := postgres.NewPostgresClient(env)

	log.Info("Running migrations")
	migrator := migrate.NewMigrator(db.Db, migrations.Migrations)
	if err := migrator.Init(ctx); err != nil {
		panic(err)
	}
	mig, err := migrator.Migrate(ctx)
	if err != nil {
		panic(err)
	}
	if mig != nil {
		log.Info(fmt.Sprintf("Migration applied: %s", mig.String()))
	}
	awsClient, err := http_out.NewAwsClient(env)
	if err != nil {
		log.Error("error creating aws client")
		panic(err)
	}

	svc := service.NewService(env, db, awsClient)

	log.Info("Running server")
	r := gin.New()
	http2.Register(r, env, log, svc)
	port := fmt.Sprintf(":%d", env.Port)

	h := &http.Server{
		Addr:    port,
		Handler: r,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	log.Info(fmt.Sprintf("Server running on : %d\n", env.Port))

	go func() {
		err := h.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Error("An error occurred", err)
			log.Error("Terminating program", nil)
			os.Exit(1)
		}
	}()
	<-stop

	log.Info("Shutting down server...")

	timeout := time.Duration(5) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = h.Shutdown(ctx)
	if err != nil {
		log.Error("Failed to shutdown gracefully", err)
	} else {
		log.Info("Server shutdown gracefully")
	}
}
