package users_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mishalyamets1/golang_todoapp/internal/core/domain"
	core_errors "github.com/mishalyamets1/golang_todoapp/internal/core/errors"
)

func (r *UsersRepository) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT * FROM todoapp.users WHERE id=$1`

	row := r.pool.QueryRow(ctx, query, id)

	var UserModel UserModel

	err := row.Scan(
		&UserModel.ID,
		&UserModel.Version,
		&UserModel.FullName,
		&UserModel.PhoneNumber,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%d' : %w",
				id,
				core_errors.ErrNotFound,
			)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}
	userDomain := domain.NewUser(UserModel.ID, UserModel.Version, UserModel.FullName, UserModel.PhoneNumber)
	return userDomain, nil

}