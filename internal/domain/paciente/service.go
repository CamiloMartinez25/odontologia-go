package paciente

import (
	"context"
	"errors"
	"log"
)

type service struct {
	repository Repository
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

func requestToPaciente(requestPaciente RequestPaciente) Paciente {
	var paciente Paciente
	paciente.Nombre = requestPaciente.Nombre
	paciente.Apellido = requestPaciente.Apellido
	paciente.Domicilio = requestPaciente.Domicilio
	paciente.DNI = requestPaciente.DNI
	paciente.FechaAlta = requestPaciente.FechaAlta

	return paciente
}
