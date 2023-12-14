package service

import "fmt"

func (s *Service) SendSupplierEmail(suppliers []string, category string) error {

	for i, _ := range suppliers {
		err := s.awsClient.SendEmail(suppliers[i], category)
		if err != nil {
			return fmt.Errorf("error sending email: %w", err)
		}
	}
	return nil
}
