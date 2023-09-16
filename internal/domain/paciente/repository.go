package paciente

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	Create(ctx context.Context, paciente Paciente) (Paciente, error)
	//GetAll(ctx context.Context) ([]Paciente, error)
	GetByID(ctx context.Context, id int) (Paciente, error)
	//Update(ctx context.Context, paciente Paciente) (Paciente, error)
	//Delete(ctx context.Context, id int) error
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

	statement, err := r.db.Prepare(QueryInsertPaciete)

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

func (r *repository) UpdateSubject(ctx context.Context, id int, request RequestUpdatePacienteSubject) (Paciente, error) {
	statement, err := r.db.Prepare(QueryUpdateSubject + request.key + " = ? WHERE ID = ?")
	if err != nil {
		return Paciente{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		request.value,
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
