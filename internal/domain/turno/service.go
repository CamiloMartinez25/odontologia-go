package turno

import (
	"context"
	"log"
)

type service struct {
	repository Repository
}

type Service interface {
	// Create(ctx context.Context, RequestTurno turno) (Turno, error)
	// CreateByPaciente(ctx context.Context, RequestTurnoByPaciente turno) (Turno, error)
	// GetByID(ctx context.Context, id int) (Turno, error)
	// GetByPacienteID(ctx context.Context, id int) ([]Turno, error)
	// Update(ctx context.Context, RequestTurno turno, id int) (Turno, error)
	// UpdatePatch(ctx context.Context, RequestTurno turno, id int) (Turno, error)
	// Delete(ctx context.Context, id int) error
}

// TurnoService creates a new turno service.
func TurnoService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// GetAll returns all turnos.
func (s *service) GetAll(ctx context.Context) ([]Turno, error) {
	turnos, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("log de error en service de turnos", err.Error())
		return []Turno{}, ErrEmptyList
	}

	return turnos, nil
}

// GetByPacienteID returns a list of turnos according to paciente's ID.
