package turno

import "time"

// Structure Turno describes a turno entity.
type Turno struct {
	ID          int        `json:"id"`
	Paciente    Paciente   `json:"paciente"`
	Odontologo  Odontologo `json:"odontologo"`
	FechaHora   time.Time  `json:"fechaHora"`
	Descripcion string     `json:"descripcion"`
}

type RequestTurno struct {
	Paciente    Paciente   `json:"paciente"`
	Odontologo  Odontologo `json:"odontologo"`
	FechaHora   time.Time  `json:"fechaHora"`
	Descripcion string     `json:"descripcion"`
}

type RequestTurnoByPaciente struct {
	Paciente    Paciente   `json:"paciente"`
	Odontologo  Odontologo `json:"odontologo"`
}

