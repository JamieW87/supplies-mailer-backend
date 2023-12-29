package service

import (
	"context"
	"fmt"

	"one-stop/internal/model"
)

func (s *Service) SendSupplierEmail(ctx context.Context, suppliers []model.SendSupplierInfo, name, email, category, message string) error {

	for i, _ := range suppliers {
		err := s.msClient.SendEmail(ctx, suppliers[i].Email, suppliers[i].Name, name, email, category, message)
		if err != nil {
			return fmt.Errorf("error sending email: %w", err)
		}
	}
	return nil
}
