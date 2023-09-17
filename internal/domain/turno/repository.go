package turno

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}
type Repository interface {
	Create(ctx context.Context, turno Turno) (Turno, error)
	GetAll(ctx context.Context) ([]Turno, error)
	//GetByID(ctx context.Context, id int) (Turno, error)
	GetByPacienteID(ctx context.Context, id int) (Turno, error)
	Update(ctx context.Context, turno Turno) (Turno, error)
	UpdateSubject(ctx context.Context, id int, request RequestUpdateTurnoSubject) (Odontologo, error)
	//Delete(ctx context.Context, id int) error

}

// TurnoRepository creates a new repository.
func TurnoRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

//Create crea un turno 
func (r *repository) Create(ctx context.Context, turno Turno) (Turno, error) {

	statement, err := r.db.Prepare(QueryInsertTurn)

	if err != nil {
		return Turno{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		turno.Paciente,
		turno.Odontologo,
		turno.FechaHora,
		turno.Descripcion
	)

	if err != nil {
		return Turno{}, ErrExec
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return Turno{}, ErrLastId
	}

	turno.ID = int(lastId)

	return turno, nil
}

// GetAll returns all turnos.
func (r *repository) GetAll(ctx context.Context) ([]Turno, error) {
	rows, err := r.db.Query(QueryGetAllTurnos)
	if err != nil {
		return []Turno{}, err
	}

	defer rows.Close()

	var turnos []Turno

	for rows.Next() {
		var turno Turno
		err := rows.Scan(
			&turno.ID,
			&turno.Paciente,
			&turno.Odontologo,
			&turno.FechaHora,
			&turno.Descripcion,
		)
		if err != nil {
			return []Turno{}, err
		}

		turnos = append(turnos, turno)
	}

	if err := rows.Err(); err != nil {
		return []Turno{}, err
	}

	return turnos, nil
}

// Update updates an turno.
func (r *repository) Update(ctx context.Context, turno Turno) (Turno, error) {
	statement, err := r.db.Prepare(QueryUpdateTurno)
	if err != nil {
		return Turno{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		turno.Paciente,
		turno.Odontologo,
		turno.FechaHora,
		turno.Descripcion,
	)

	if err != nil {
		return Turno{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Turno{}, err
	}

	if rowsAffected < 1 {
		return Turno{}, ErrNotFound
	}

	return turno, nil
}

// GetByPacienteID returns a list of turnos according to paciente's ID.
func (r *repository) GetByPacienteID(ctx context.Context, id int) ([]Turno, error) {
	rows, err := r.db.Query(QueryGetTurnByPacienteID, id)

	if err != nil {
		return []Turno{}, err
	}

	defer rows.Close()

	var turnos []Turno

	for rows.Next() {
		var turno Turno
		err := rows.Scan(
			&turno.ID,
			&turno.Paciente,
			&turno.Odontologo,
			&turno.FechaHora,
			&turno.Descripcion,
		)
		if err != nil {
			return []Turno{}, err
		}

		turnos = append(turnos, turno)
	}

	if err := rows.Err(); err != nil {
		return []Turno{}, err
	}

	return turnos, nil
}

func (r *repository) UpdateSubject(ctx context.Context, id int, request RequestUpdateTurnoSubject) (Turno, error) {

	statement, err := r.db.Prepare(QueryUpdateTurnoSubject + request.key + " = ? WHERE ID = ?")
	if err != nil {
		return Turno{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		request.value,
		id,
	)

	if err != nil {
		return Turno{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Turno{}, err
	}

	if rowsAffected < 1 {
		return Turno{}, ErrNotFound
	}

	turnoActualizado, err := r.GetByID(ctx, id)
	if err != nil {
		return Turno{}, err
	}

	return turnoActualizado, nil
}