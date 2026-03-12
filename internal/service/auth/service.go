package auth

import (
	"context"
	"errors"
	"time"

	"github.com/EdwardShiroki/webstore/internal/domain/user"
	"github.com/google/uuid"
)

type Service struct {
	repo user.Repository
}

func NewService(userRepo user.Repository) *Service {
	return &Service{repo: userRepo}
}

func (s *Service) Register(ctx context.Context, login, password string) (*user.User, error) {
	existingUser, err := s.repo.GetByLogin(ctx, login)
	if err != nil && !errors.Is(err, user.ErrUserNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with this login already exists")
	}
	passwordHash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	newUser := &user.User{
		ID:           uuid.New(),
		Login:        login,
		PasswordHash: passwordHash,
		Role:         "user",
		CreatedAt:    time.Now().Format(time.RFC3339),
	}
	err = s.repo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
