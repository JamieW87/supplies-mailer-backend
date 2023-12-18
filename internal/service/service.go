package service

import (
	"one-stop/internal/config"
	"one-stop/internal/store/postgres"
	http_out "one-stop/internal/transport/outbound/http"
)

type Service struct {
	env      *config.Environment
	db       *postgres.PostgresStore
	msClient *http_out.MailerSendClient
}

var validCategory = map[string]bool{
	"roofing":   true,
	"brickwork": true,
}

func NewService(env *config.Environment, db *postgres.PostgresStore, msc *http_out.MailerSendClient) *Service {
	return &Service{
		env:      env,
		db:       db,
		msClient: msc,
	}
}

func (s *Service) IsValidCategory(category string) bool {
	_, valid := validCategory[category]
	return valid
}
