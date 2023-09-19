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
	GetByID(ctx context.Context, id int) (Odontologo, error)
	Update(ctx context.Context, requestOdontologo RequestOdontologo, id int) (Odontologo, error)
	Delete(ctx context.Context, id int) error
	UpdateSubject(ctx context.Context, id int, request RequestUpdateOdontologoSubject) (Odontologo, error)
}

// NewService creates a new odontologo service.
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// Create creates an odontologo
func (s *service) Create(ctx context.Context, requestOdontologo RequestOdontologo) (Odontologo, error) {
	odontologo := requestToOdontologo(requestOdontologo)
	log.Println(odontologo)
	response, err := s.repository.Create(ctx, odontologo)
	if err != nil {
		log.Println("Error en service Odontologo: Método Create")
		log.Println(err)
		return Odontologo{}, errors.New("Error en service Odontologo: Método Create")
	}

	return response, nil
}

// Get return an odontologo by ID
func (s *service) GetByID(ctx context.Context, id int) (Odontologo, error) {
	odontologo, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("Error en el service Odontologo: Método GetByID", err.Error())
		return Odontologo{}, errors.New("Error en service Odontologo: Metodo GetByID")
	}

	return odontologo, nil
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

// Update actualiza alguno de los campos de odontologo
func (s *service) UpdateSubject(ctx context.Context, id int, request RequestUpdateOdontologoSubject) (Odontologo, error) {
	log.Println(request)
	response, err := s.repository.UpdateSubject(ctx, id, request)
	if err != nil {
		log.Println("log de error en service de odontologo", err.Error())
		return Odontologo{}, errors.New("error en servicio. Metodo UpdateName")
	}
	return response, nil
}

// Delete elimina odontologo
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("Error en el service Odontologo: Método Delete", err.Error())
		return errors.New("Error en service: Método Delete")
	}

	return nil
}

func requestToOdontologo(requestOdontologo RequestOdontologo) Odontologo {
	var odontologo Odontologo
	odontologo.Nombre = requestOdontologo.Nombre
	odontologo.Apellido = requestOdontologo.Apellido
	odontologo.Matricula = requestOdontologo.Matricula

	return odontologo
}
