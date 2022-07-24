package storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/discreto13/go-gin-microservice/internal/core"
)

var ErrDuplicatedUserId = errors.New("user id already exist")
var ErrUserNotFound = errors.New("user not found")

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	storage := &UserStorage{
		db: db,
	}

	return storage
}

func (s *UserStorage) Insert(ctx context.Context, newUser *core.User) error {
	_, err := s.db.Exec("INSERT INTO users(id,name,email,birthday) VALUES($1,$2,$3,$4)",
		newUser.Id, newUser.Name, newUser.Email, newUser.Birthday)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStorage) GetById(ctx context.Context, id string) (*core.User, error) {
	var user core.User
	err := s.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.Id, &user.Name, &user.Email, &user.Birthday)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserStorage) GetAll(ctx context.Context) ([]*core.User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var usersList []*core.User
	for rows.Next() {
		var user core.User
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.Birthday)
		usersList = append(usersList, &user)
	}

	return usersList, nil
}
