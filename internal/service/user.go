package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/discreto13/go-gin-microservice/internal/core"
	"github.com/google/uuid"
)

const (
	kDataLayout = "2006-01-02"
)

type UserStorage interface {
	Insert(ctx context.Context, userService *core.User) error
	GetById(ctx context.Context, id string) (*core.User, error)
	GetAll(ctx context.Context) ([]*core.User, error)
	Delete(ctx context.Context, id string) (bool, error)
}

func NewUserService(storage UserStorage) *userService {
	return &userService{storage}
}

type userService struct {
	storage UserStorage
}

func (s *userService) Create(ctx context.Context, input *core.CreateUser) (*core.User, error) {
	// Validate input values
	bdate, err := time.Parse(kDataLayout, input.Birthday)

	if err != nil {
		return nil, fmt.Errorf("failed to parse birthday: %w", err)
	}

	if !bdate.Before(time.Now().AddDate(-18, 0, 0)) {
		return nil, errors.New("younger than 18yo")
	}

	// [TODO] ensure unique email by DB
	// for _, e := range s.users {
	// 	if e == input.Email {
	// 		return nil, errors.New("younger than 18yo")
	// 	}
	// }

	// Insert new user
	user := &core.User{
		Id:       uuid.New().String(),
		Name:     input.Name,
		Email:    input.Email,
		Birthday: bdate.Format(kDataLayout),
	}

	if err := s.storage.Insert(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to insert new user: %w", err)
	}

	return user, nil
}

func (s *userService) GetById(ctx context.Context, id string) (*core.User, error) {
	return s.storage.GetById(ctx, id)
}

func (s *userService) GetAll(ctx context.Context) ([]*core.User, error) {
	return s.storage.GetAll(ctx)
}

func (s *userService) Delete(ctx context.Context, id string) (bool, error) {
	return s.storage.Delete(ctx, id)
}
