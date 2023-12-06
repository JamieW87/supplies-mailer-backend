package service

import "one-stop/internal/config"

type Service struct {
	env *config.Environment
}

func NewService(env *config.Environment) *Service {
	return &Service{env: env}
}
