package paciente

import (
	"context"
	"errors"
	"log"
)

type service struct {
	repository Repository
}

type Service interface {
	Create(ctx context.Context, requestPaciente RequestPaciente) (Paciente, error)
	GetByID(ctx context.Context, id int) (Paciente, error)
	//Update(ctx context.Context, requestPaciente RequestPaciente, id int) (Paciente, error)
	Delete(ctx context.Context, id int) error
	UpdateSubject(ctx context.Context, id int, request RequestUpdatePacienteSubject) (Paciente, error)
}

// NewService creates a new product service.
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, requestPaciente RequestPaciente) (Paciente, error) {
	paciente := requestToPaciente(requestPaciente)
	response, err := s.repository.Create(ctx, paciente)
	if err != nil {
		log.Println("Error en service Paciente: Método Create")
		return Paciente{}, errors.New("Error en service Paciente: Método Create")
	}

	return response, nil
}

// GetByID returns a paciente by its ID.
func (s *service) GetByID(ctx context.Context, id int) (Paciente, error) {
	paciente, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("log de error en service de paciente", err.Error())
		return Paciente{}, errors.New("error en servicio. Metodo GetByID")
	}

	return paciente, nil
}

// Update updates an paciente.
func (s *service) Update(ctx context.Context, requestPaciente RequestPaciente, id int) (Paciente, error) {
	paciente := requestToPaciente(requestPaciente)
	paciente.ID = id
	response, err := s.repository.Update(ctx, paciente)
	if err != nil {
		log.Println("log de error en service de paciente", err.Error())
		return Paciente{}, errors.New("error en servicio. Metodo Update")
	}

	return response, nil
}

func (s *service) UpdateSubject(ctx context.Context, id int, request RequestUpdatePacienteSubject) (Paciente, error) {
	response, err := s.repository.UpdateSubject(ctx, id, request)
	if err != nil {
		log.Println("log de error en service de paciente", err.Error())
		return Paciente{}, errors.New("error en servicio. Metodo UpdateSubjet")
	}
	return response, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("log de error borrado de paciente", err.Error())
		return errors.New("error en servicio. Metodo Delete")
	}

	return nil
}

func requestToPaciente(requestPaciente RequestPaciente) Paciente {
	var paciente Paciente
	paciente.Nombre = requestPaciente.Nombre
	paciente.Apellido = requestPaciente.Apellido
	paciente.Domicilio = requestPaciente.Domicilio
	paciente.DNI = requestPaciente.DNI
	paciente.FechaAlta = requestPaciente.FechaAlta

	return paciente
}
