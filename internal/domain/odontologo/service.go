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
	Create(ctx context.Context, requestOdontologo RequestOdontologo) (Odontologo, error)
	GetAll(ctx context.Context) ([]Odontologo, error)
	GetByID(ctx context.Context, id int) (Odontologo, error)
	Update(ctx context.Context, requestOdontologo RequestOdontologo, id int) (Odontologo, error)
	Delete(ctx context.Context, id int) error
}

// NewService creates a new product service.
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, requestOdontologo RequestOdontologo) (Odontologo, error) {
	odontologo := requestToOdontologo(requestOdotologo)
	response, err := s.repository.Create(ctx, odontologo)
	if err != nil {
		log.Println("Error en service odontologo: Método Create")
		return Odontologo{}, errors.New("Error en service odontologo: Método Create")
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
