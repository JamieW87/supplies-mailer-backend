package service

import "context"

func (s *Service) RetrieveSupplierInfo(ctx context.Context, category string) ([]string, error) {

	return s.db.GetSupplierEmailsForCategory(ctx, category)

}
