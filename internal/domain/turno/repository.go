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