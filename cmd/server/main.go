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

	"one-stop/internal/config"
	"one-stop/internal/service"
	http2 "one-stop/internal/transport/inbound/http"
)

func main() {
	log := logrus.New()
	env, err := config.Get()
	if err != nil {
		log.Error("Could not retrieve environment", err)
		panic(err)
	}

	r := gin.New()
	svc := service.NewService(env)
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
