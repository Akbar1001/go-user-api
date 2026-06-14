package repository

import (
	"context"

	"go-user-api/db/sqlc/generated"
)

type UserRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(
	queries *sqlc.Queries,
) *UserRepository {
	return &UserRepository{
		queries: queries,
	}
}