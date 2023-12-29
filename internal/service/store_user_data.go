package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"one-stop/internal/model"
)

func (s *Service) StoreUserData(ctx context.Context, name, email, phone, category, postcode string) (uuid.UUID, error) {

	now := time.Now()

	user := &model.User{
		Id:        uuid.New(),
		Name:      name,
		Email:     email,
		Phone:     phone,
		Postcode:  postcode,
		CreatedAt: now,
		UpdatedAt: now,
	}

	userId, err := s.db.InsertUser(ctx, user)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error inserting user: %w", err)
	}

	err = s.db.InsertUserCategory(ctx, userId, category)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error inserting user category: %w", err)
	}
	return userId, err
}
