package turno


import (
	"context"
	"errors"
	"log"
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
func TurnoService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
