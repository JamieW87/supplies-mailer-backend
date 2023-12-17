package service

import (
	"one-stop/internal/config"
	"one-stop/internal/store/postgres"
	http_out "one-stop/internal/transport/outbound/http"
)

type Service struct {
	env       *config.Environment
	db        *postgres.PostgresStore
	awsClient *http_out.AWSClient
}

var ValidCategories = map[string]bool{
	"roofing":   true,
	"brickwork": true,
}

func NewService(env *config.Environment, db *postgres.PostgresStore, awsC *http_out.AWSClient) *Service {
	return &Service{
		env:       env,
		db:        db,
		awsClient: awsC,
	}
}

func (s *Service) ValidateCategory(category string) bool {

	return false
}
