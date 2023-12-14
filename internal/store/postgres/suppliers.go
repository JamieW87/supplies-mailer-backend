package postgres

import (
	"context"
	"fmt"

	"one-stop/internal/model"
)

func (pg PostgresStore) GetSupplierEmailsForCategory(ctx context.Context, category string) ([]string, error) {

	var Supplier model.Supplier

	var emails []string
	err := pg.Db.NewSelect().Model(Supplier).Column("email").Where("category = ?", category).Scan(ctx, &emails)
	if err != nil {
		return nil, fmt.Errorf("error retrieving emails: %w", err)
	}

	return emails, nil

}
