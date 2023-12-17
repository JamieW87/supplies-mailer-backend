package service

import (
	"fmt"

	"one-stop/internal/model"
)

func (s *Service) SendSupplierEmail(suppliers []model.SendSupplierInfo, name, email, category string) error {

	for i, _ := range suppliers {
		err := s.awsClient.SendEmail(suppliers[i].Email, suppliers[i].Name, name, email, category)
		if err != nil {
			return fmt.Errorf("error sending email: %w", err)
		}
	}
	return nil
}
