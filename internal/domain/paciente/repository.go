package paciente

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrEmptyList = errors.New("the list is empty")
	ErrNotFound  = errors.New("paciente not found")
	ErrStatement = errors.New("error preparing statement")
	ErrExec      = errors.New("error exect statement")
	ErrLastId    = errors.New("error getting last id")
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	Create(ctx context.Context, paciente Paciente) (Paciente, error)
	GetByID(ctx context.Context, id int) (Paciente, error)
	Update(ctx context.Context, paciente Paciente) (Paciente, error)
	Delete(ctx context.Context, id int) error
	UpdateSubject(ctx context.Context, id int, request RequestUpdatePacienteSubject) (Paciente, error)
}

// NewRepositoryMySql crea un nuevo repositorio.
func NewRepositoryMySql(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create crea un nuevo paciente.
func (r *repository) Create(ctx context.Context, paciente Paciente) (Paciente, error) {

	statement, err := r.db.Prepare(QueryInsertPaciente)

	if err != nil {
		return Paciente{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		paciente.Nombre,
		paciente.Apellido,
		paciente.Domicilio,
		paciente.DNI,
		paciente.FechaAlta,
	)

	if err != nil {
		return Paciente{}, ErrExec
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return Paciente{}, ErrLastId
	}

	paciente.ID = int(lastId)

	return paciente, nil
}

// GetByID returns a paciente by its ID.
func (r *repository) GetByID(ctx context.Context, id int) (Paciente, error) {
	row := r.db.QueryRow(QueryGetPacienteById, id)

	var paciente Paciente
	err := row.Scan(
		&paciente.ID,
		&paciente.Nombre,
		&paciente.Apellido,
		&paciente.Domicilio,
		&paciente.DNI,
		&paciente.FechaAlta,
	)

	if err != nil {
		return Paciente{}, err
	}

	return paciente, nil
}

func (r *repository) Update(ctx context.Context, paciente Paciente) (Paciente, error) {
	statement, err := r.db.Prepare(QueryUpdatePaciente)
	if err != nil {
		return Paciente{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		paciente.Nombre,
		paciente.Apellido,
		paciente.Domicilio,
		paciente.DNI,
		paciente.FechaAlta,
		paciente.ID,
	)

	if err != nil {
		return Paciente{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Paciente{}, err
	}

	if rowsAffected < 1 {
		return Paciente{}, ErrNotFound
	}

	return paciente, nil
}

func (r *repository) UpdateSubject(ctx context.Context, id int, request RequestUpdatePacienteSubject) (Paciente, error) {
	statement, err := r.db.Prepare(QueryUpdateSubject + request.Key + " = ? WHERE ID = ?")
	if err != nil {
		return Paciente{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		request.Value,
		id,
	)

	if err != nil {
		return Paciente{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Paciente{}, err
	}

	if rowsAffected < 1 {
		return Paciente{}, ErrNotFound
	}

	pacienteActualizado, err := r.GetByID(ctx, id)
	if err != nil {
		return Paciente{}, err
	}

	return pacienteActualizado, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeletePaciente, id)
	if err != nil {
		return ErrExec
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
