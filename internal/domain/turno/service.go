package turno


import (
	"context"
	"errors"
	"log"
)


var (
	ErrEmptyList = errors.New("the list is empty")
)

type service struct {
	repository Repository
}

type Service interface {
	Create(ctx context.Context, RequestTurno turno) (Turno, error)
	CreateByPaciente(ctx context.Context, RequestTurnoByPaciente turno) (Turno, error)
	GetByID(ctx context.Context, id int) (Turno, error)
	GetByPacienteID(ctx context.Context, id int) ([]Turno, error)
	Update(ctx context.Context, RequestTurno turno, id int) (Turno, error)
	UpdatePatch(ctx context.Context, RequestTurno turno, id int) (Turno, error)
	Delete(ctx context.Context, id int) error
}

// TurnoService creates a new turno service.
func TurnoSefvice(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// GetByPacienteID returns a list of turnos according to paciente's ID.
func (s *service) GetByPacienteID(ctx context.Context) ([]Turno, error) {
	turnos, err := s.repository.GetByPacienteID(ctx)
	if err != nil {
		log.Println("Error on turnos service", err.Error())
		return []Turno{}, ErrEmptyList
	}

	return turnos, nil
}