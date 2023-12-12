package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"one-stop/internal/model"
)

func (pg PostgresStore) InsertUser(ctx context.Context, user *model.User) (uuid.UUID, error) {

	var selectUser model.User
	err := pg.Db.NewSelect().Model(&selectUser).Where("email = ?", user.Email).Scan(ctx)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			_, err = pg.Db.NewInsert().Model(user).Exec(ctx)
			if err != nil {
				return uuid.Nil, fmt.Errorf("error inserting user: %w", err)
			}
			return user.Id, nil
		default:
			return uuid.Nil, fmt.Errorf("error retrieving user: %w", err)
		}
	}
	return selectUser.Id, nil
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
