package repositories

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pandukamuditha/simple-blog/internal/models"
)

type UserRepository struct {
	DB *pgx.Conn
}

const createAccountQuery = `
INSERT INTO public."user" (
	fname, lname
) VALUES (
	$1, $2
) RETURNING id, fname, lname
`

type CreateUserParams struct {
	FName string `json:"firstName"`
	LName string `json:"lastName"`
}

func (r *UserRepository) CreateUser(ctx context.Context, args CreateUserParams) (*models.User, error) {
	row := r.DB.QueryRow(ctx, createAccountQuery, args.FName, args.LName)

	var i models.User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
	)

	if err != nil {
		return nil, err
	}

	return &i, nil
}

const getUserQuery = `
	SELECT 
		"id", "fname", "lname"
	FROM
		public."user"
	WHERE
		"id" = $1
`

func (r *UserRepository) GetUserById(ctx context.Context, userId int) (*models.User, error) {
	row := r.DB.QueryRow(ctx, getUserQuery, userId)

	var i models.User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
	)

	if err != nil {
		return nil, err
	}

	return &i, nil
}

