package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"one-stop/internal/model"
)

func (pg PostgresStore) InsertUser(ctx context.Context, user *model.User) error {

	//SELECT user_id FROM users WHERE email = 'user_email@example.com';

	//If not exists ==
	_, err := pg.Db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}

	return nil
}

func (pg PostgresStore) InsertUserCategory(ctx context.Context, userId uuid.UUID, category string) error {

	var cat model.Category

	err := pg.Db.NewSelect().Model(&cat).Where("name = ?", category).Scan(ctx)
	if err != nil {
		return fmt.Errorf("error retrieving user category: %w", err)
	}

	userCat := &model.UserCategory{
		UserId:     userId,
		CategoryId: cat.Id,
	}

	_, err = pg.Db.NewInsert().Model(userCat).Exec(ctx)
	if err != nil {
		return fmt.Errorf("error inserting user category: %w", err)
	}

	return nil
}
