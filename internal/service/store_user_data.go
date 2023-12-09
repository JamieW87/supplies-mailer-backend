package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"one-stop/internal/model"
)

func (s *Service) StoreUserData(ctx context.Context, name, email, phone string) (uuid.UUID, error) {

	now := time.Now()

	user := &model.User{
		Id:        uuid.New(),
		Name:      name,
		Email:     email,
		Phone:     phone,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.db.InsertUser(ctx, user)
	if err != nil {
		return uuid.UUID{}, err
	}
	return user.Id, err
}

func (s *Service) InsertUserCategory(ctx context.Context, userId uuid.UUID, category string) error {

	return s.db.InsertUserCategory(ctx, userId, category)
}
