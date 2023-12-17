package postgres

import (
	"context"

	"one-stop/internal/model"
)

func (pg PostgresStore) GetSupplierEmailsForCategory(ctx context.Context, category string) ([]model.SendSupplierInfo, error) {

	var suppliers []model.SendSupplierInfo
	err := pg.Db.NewSelect().
		Column("s.name", "s.email").
		TableExpr("suppliers AS s").
		Join("JOIN supplier_categories AS sc ON s.id = sc.supplier_id").
		Join("JOIN categories AS c ON sc.category_id = c.id").
		Where("c.name = ?", category).
		Scan(ctx, &suppliers)

	return suppliers, err

}
