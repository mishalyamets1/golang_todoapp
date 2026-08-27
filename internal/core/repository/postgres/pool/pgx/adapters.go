package core_pgx_pool

import (
	"errors"

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
		if errors.Is(err, pgx.ErrNoRows) {
			return core_postgres_pool.ErrNoRows
		}
		return err
	}
	return nil
}
type pgxRow struct {
	pgx.Row
}
type pgxCommandTag struct {
	pgconn.CommandTag
}