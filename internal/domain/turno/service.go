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
	//Create(ctx context.Context, RequestTurno turno) (Turno, error)
	//CreateByPaciente(ctx context.Context, RequestTurnoByPaciente turno) (Turno, error)
	//GetByID(ctx context.Context, id int) (Turno, error)
	GetByPacienteID(ctx context.Context, id int) ([]Turno, error)
	Update(ctx context.Context, requestTurno RequestTurno , id int) (Turno, error)
	UpdateSubject(ctx context.Context, id int, request RequestUpdateTurnoSubject,) (Turno, error)
	//Delete(ctx context.Context, id int) error
}

// TurnoService creates a new turno service.
func TurnoService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// Update updates an turno.
func (s *service) Update(ctx context.Context, requestTurno RequestTurno, id int) (Turno, error) {
	turno := requestToTurno(requestTurno)
	turno.ID = id
	response, err := s.repository.Update(ctx, turno)
	if err != nil {
		log.Println("log de error en service de turno", err.Error())
		return Turno{}, errors.New("error en turno. Metodo Update")
	}

	return response, nil
}


// GetByPacienteID returns a list of turnos according to paciente's ID.
func (s *service) GetByPacienteID(ctx context.Context, id int) ([]Turno, error) {
	turnos, err := s.repository.GetByPacienteID(ctx, id)
	if err != nil {
		log.Println("Error on turnos service", err.Error())
		return []Turno{}, ErrEmptyList
	}

	return turnos, nil
}

func (s *service) UpdateSubject(ctx context.Context, id int, request RequestUpdateTurnoSubject  ) (Turno, error) {

	response, err := s.repository.UpdateSubject(ctx, id, request)
	if err != nil {
		log.Println("log de error en service de turnos", err.Error())
		return Turno{}, errors.New("error en servicio. Metodo UpdateName")
	}
	return response, nil
}

func requestToTurno(requestTurno RequestTurno) Turno {
	var turno Turno
	turno.Paciente = requestTurno.Paciente
	turno.Odontologo = requestTurno.Odontologo
	turno.FechaHora = requestTurno.FechaHora
	turno.Descripcion = requestTurno.Descripcion

	return turno
}
