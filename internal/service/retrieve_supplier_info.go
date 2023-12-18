package service

import (
	"context"

	"one-stop/internal/model"
)

func (s *Service) RetrieveSupplierInfo(ctx context.Context, category string) ([]model.SendSupplierInfo, error) {

	return s.db.GetSupplierEmailsForCategory(ctx, category)
}
