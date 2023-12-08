package service

import (
	"one-stop/internal/config"
	"one-stop/internal/store/postgres"
)

type Service struct {
	env *config.Environment
	db  *postgres.PostgresStore
}

func NewService(env *config.Environment, db *postgres.PostgresStore) *Service {
	return &Service{
		env: env,
		db:  db}
}
