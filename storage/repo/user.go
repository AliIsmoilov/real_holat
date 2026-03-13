package repo

import (
	"context"

	"github.com/google/uuid"
)

type UserI interface {
	Create(ctx context.Context, req User) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type User struct {
	Id       uuid.UUID
	Email    string
	Password string // TODO: implement password hashing
	Role     string
}
