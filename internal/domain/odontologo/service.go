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

// NewService creates a new odontologo service.
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}


// Update updates an odontologo.
func (s *service) Update(ctx context.Context, requestOdontologo RequestOdontologo, id int) (Odontologo, error) {
	odontologo := requestToodontologo(requestOdontologo)
	odontologo.ID = id
	response, err := s.repository.Update(ctx, odontologo)
	if err != nil {
		log.Println("log de error en service de odontologo", err.Error())
		return Odontologo{}, errors.New("error en servicio. Metodo Update")
	}

	return response, nil
}

func requestToOdontologo(requestOdontologo RequestOdontologo) Odontologo {
	var odontologo Odontologo
	odontologo.Nombre = requestodOdontologo.Nombre
	odontologo.Apellido = requestodOdontologo.Apellido
	odontologo.Matricula = requestodOdontologo.Matricula

	return odontologo
}