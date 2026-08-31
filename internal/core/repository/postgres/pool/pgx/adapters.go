package core_pgx_pool

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	core_postgres_pool "github.com/mishalyamets1/golang_todoapp/internal/core/repository/postgres/pool"
)

type pgxRows struct {
	pgx.Rows
}
func (r *pgxRow) Scan(dest... any) error {

	
	err := r.Row.Scan(dest...) 
	if err != nil {
		return mapErrors(err)
	}
	
	return nil
}
type pgxRow struct {
	pgx.Row
}
type pgxCommandTag struct {
	pgconn.CommandTag
}

func mapErrors(err error) error {
	const pgxViolatesForeignErrorCode = "23503"

	if errors.Is(err, pgx.ErrNoRows) {
		return core_postgres_pool.ErrNoRows

	}

	var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgxViolatesForeignErrorCode {
				return fmt.Errorf("%v : %w", err, core_postgres_pool.ErrViolatesForeignKey)
			}
		}
		return fmt.Errorf("%v: %w", err, core_postgres_pool.ErrUnknown)
} 