package odontologo

import (
	"context"
	"database/sql"
)

var (
	ErrNotFound  = errors.New("product not found")
)

type repository struct {
	db *sql.DB
}

func NewRepositoryMySql(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}


// Update updates an odontologo.
func (r *repository) Update(ctx context.Context, ondontologo Odontologo) (Odontologo, error) {
	statement, err := r.db.Prepare(QueryUpdateOdontologo)
	if err != nil {
		return Odontologo{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		odontologo.Nombre,
		odontologo.Apellido,
		odontologo.Matricula,
		odontologo.ID,
	)

	if err != nil {
		return Odontologo{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Odontologo{}, err
	}

	if rowsAffected < 1 {
		return Odontologo{}, ErrNotFound
	}

	return odontologo, nil
}