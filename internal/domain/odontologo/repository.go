package odontologo

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound  = errors.New("odontologo not found")
	ErrEmptyList = errors.New("the list is empty")
	ErrStatement = errors.New("error preparing statement")
	ErrExec      = errors.New("error exect statement")
	ErrLastId    = errors.New("error getting last id")
)

type Repository interface {
	Create(ctx context.Context, odontologo Odontologo) (Odontologo, error)
	GetByID(ctx context.Context, id int) (Odontologo, error)
	Update(ctx context.Context, odontologo Odontologo) (Odontologo, error)
	Delete(ctx context.Context, id int) error
	UpdateSubject(ctx context.Context, id int, request RequestUpdateOdontologoSubject) (Odontologo, error)
}

type repository struct {
	db *sql.DB
}

func NewRepositoryMySql(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create crea un nuevo odontologo.
func (r *repository) Create(ctx context.Context, odontologo Odontologo) (Odontologo, error) {

	statement, err := r.db.Prepare(QueryInsertOdontologo)

	if err != nil {
		return Odontologo{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		odontologo.Nombre,
		odontologo.Apellido,
		odontologo.Matricula,
	)

	if err != nil {
		return Odontologo{}, ErrExec
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return Odontologo{}, ErrLastId
	}

	odontologo.ID = int(lastId)

	return odontologo, nil
}

// GetByID retorna un odontologo por ID
func (r *repository) GetByID(ctx context.Context, id int) (Odontologo, error) {
	row := r.db.QueryRow(QueryGetOdontologoById, id)

	var odontologo Odontologo
	err := row.Scan(
		&odontologo.ID,
		&odontologo.Nombre,
		&odontologo.Apellido,
		&odontologo.Matricula,
	)

	if err != nil {
		return Odontologo{}, err
	}

	return odontologo, nil
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

// Update actualiza alguno de los campos de odontologo
func (r *repository) UpdateSubject(ctx context.Context, id int, request RequestUpdateOdontologoSubject) (Odontologo, error) {

	var query string = QueryUpdateOdontologoSubject + request.Key + " = ? WHERE ID = ?"

	statement, err := r.db.Prepare(query)
	if err != nil {
		return Odontologo{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		request.Value,
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

// Delete elimina odontologo
func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteOdontologo, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil

}
