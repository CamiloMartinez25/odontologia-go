package odontologo

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("odontologo not found")
)

type Repository interface {
	//Create(ctx context.Context, odontologo Odontologo) (Odontologo, error)
	//GetAll(ctx context.Context) ([]Odontologo, error)
	GetByID(ctx context.Context, id int) (Odontologo, error)
	Update(ctx context.Context, odontologo Odontologo) (Odontologo, error)
	//Delete(ctx context.Context, id int) error
	UpdateName(ctx context.Context, id int, nombreNuevo string) (Odontologo, error)
}

type repository struct {
	db *sql.DB
}

func NewRepositoryMySql(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

// Update updates an odontologo.
func (r *repository) Update(ctx context.Context, odontologo Odontologo) (Odontologo, error) {
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

func (r *repository) UpdateName(ctx context.Context, id int, nombreNuevo string) (Odontologo, error) {

	statement, err := r.db.Prepare(QueryUpdateOdontologoNombre)
	if err != nil {
		return Odontologo{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		nombreNuevo,
		id,
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

	odontologoActualizado, err := r.GetByID(ctx, id)
	if err != nil {
		return Odontologo{}, err
	}

	return odontologoActualizado, nil
}

func (r *repository) GetByID(ctx context.Context, id int) (Odontologo, error)
