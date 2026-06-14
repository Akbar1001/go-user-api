package repository

import (
	"context"
	"time"

	sqlc "go-user-api/db/sqlc/generated"
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

// Create user
func (r *UserRepository) CreateUser(
	name string,
	dob time.Time,
) (sqlc.User, error) {

	return r.queries.CreateUser(
		context.Background(),
		sqlc.CreateUserParams{
			Name: name,
			Dob:  dob,
		},
	)
}

// Get User
func (r *UserRepository) GetUser(
	id int32,
) (sqlc.User, error) {

	return r.queries.GetUser(
		context.Background(),
		id,
	)
}

//List Users
func (r *UserRepository) ListUsers() ([]sqlc.User, error) {

	return r.queries.ListUsers(
		context.Background(),
	)
}

// Delete User
func (r *UserRepository) DeleteUser(
	id int32,
) error {

	return r.queries.DeleteUser(
		context.Background(),
		id,
	)
}

// Update User
func (r *UserRepository) UpdateUser(
	id int32,
	name string,
	dob time.Time,
) (sqlc.User, error) {

	return r.queries.UpdateUser(
		context.Background(),
		sqlc.UpdateUserParams{
			ID:   id,
			Name: name,
			Dob:  dob,
		},
	)
}