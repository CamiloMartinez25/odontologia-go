package odontologo

import (
	"context"
	"errors"
	"log"
)

type service struct {
	repository Repository
}

type Service interface {
	//Create(ctx context.Context, requestOdontologo RequestOdontologo) (Odontologo, error)
	//GetAll(ctx context.Context) ([]Odontologo, error)
	//GetByID(ctx context.Context, id int) (Odontologo, error)
	Update(ctx context.Context, requestOdontologo RequestOdontologo, id int) (Odontologo, error)
	//Delete(ctx context.Context, id int) error
	UpdateName(ctx context.Context, id int, nombreNuevo string) (Odontologo, error)
}

// NewService creates a new odontologo service.
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, requestOdontologo RequestOdontologo) (Odontologo, error) {
	odontologo := requestToOdontologo(requestOdontologo)
	response, err := s.repository.Create(ctx, odontologo)
	if err != nil {
		log.Println("Error en service Odontologo: Método Create")
		return Odontologo{}, errors.New("Error en service Odontologo: Método Create")
	}

	return response, nil
}

// Update updates an odontologo.
func (s *service) Update(ctx context.Context, requestOdontologo RequestOdontologo, id int) (Odontologo, error) {
	odontologo := requestToOdontologo(requestOdontologo)
	odontologo.ID = id
	response, err := s.repository.Update(ctx, odontologo)
	if err != nil {
		log.Println("log de error en service de odontologo", err.Error())
		return Odontologo{}, errors.New("error en servicio. Metodo Update")
	}

	return response, nil
}


func (s *service) UpdateSubject(ctx context.Context, id int, request RequestUpdateOdontologoSubject) (Odontologo, error) {

	response, err := s.repository.UpdateSubject(ctx, id, request)
	if err != nil {
		log.Println("log de error en service de odontologo", err.Error())
		return Odontologo{}, errors.New("error en servicio. Metodo UpdateName")
	}
	return response, nil
}

func requestToOdontologo(requestOdontologo RequestOdontologo) Odontologo {
	var odontologo Odontologo
	odontologo.Nombre = requestOdontologo.Nombre
	odontologo.Apellido = requestOdontologo.Apellido
	odontologo.Matricula = requestOdontologo.Matricula

	return odontologo
}