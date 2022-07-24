package storage

import (
	"context"
	"errors"

	"github.com/discreto13/go-gin-microservice/internal/core"
)

var ErrDuplicatedUserId = errors.New("user id already exist")
var ErrUserNotFound = errors.New("user not found")

type UserStorage struct {
	users map[string]*core.User
}

func NewUserStorage() *UserStorage {
	storage := &UserStorage{
		users: make(map[string]*core.User),
	}

	// FIXME: remove adding default users for testing purposes
	storage.users["1"] = &core.User{
		Id:       "1",
		Name:     "",
		Email:    "some.name@gmail.com",
		Birthday: "15.03.2022",
	}
	storage.users["12"] = &core.User{
		Id:       "12",
		Name:     "Some",
		Email:    "some.12name@gmail.com",
		Birthday: "15.03.2022",
	}
	storage.users["123"] = &core.User{
		Id:       "123",
		Name:     "Some Name",
		Email:    "some.123name@gmail.com",
		Birthday: "15.03.2022",
	}

	return storage
}

func (s *UserStorage) Insert(ctx context.Context, newUser *core.User) error {
	if _, ok := s.users[newUser.Id]; ok {
		return ErrDuplicatedUserId
	}

	s.users[newUser.Id] = newUser
	return nil
}

func (s *UserStorage) GetById(ctx context.Context, id string) (*core.User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *UserStorage) GetAll(ctx context.Context) ([]*core.User, error) {
	usersList := make([]*core.User, 0, len(s.users))
	for _, v := range s.users {
		usersList = append(usersList, v)
	}
	return usersList, nil
}
