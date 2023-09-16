package turno

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

// TurnoRepository creates a new repository.
func TurnoRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
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
	rows, err := r.db.Query(QueryGetTurnByPacienteId, id)

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
			&turno.Descripcion
		)
		if err != nil {
			return []Turno{}, err
		}

		turnos = append(turnos, turno)
	}

	if err := rows.Err(); err != nil {
		return []Turnos{}, err
	}

	return turnos, nil
}