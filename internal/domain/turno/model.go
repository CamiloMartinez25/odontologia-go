package turno

import "time"

// Structure Turno describes a turno entity.
type Turno struct {
	ID          int        `json:"id"`
	Paciente    int   `json:"paciente"`
	Odontologo  string `json:"odontologo"`
	FechaHora   time.Time  `json:"fechaHora"`
	Descripcion string     `json:"descripcion"`
}

type RequestTurno struct {
	Paciente    int   `json:"paciente"`
	Odontologo  string `json:"odontologo"`
	FechaHora   time.Time  `json:"fechaHora"`
	Descripcion string     `json:"descripcion"`
}

type RequestTurnoByPaciente struct {
	Paciente    int   `json:"paciente"`
	Odontologo  string `json:"odontologo"`
}

type RequestUpdateTurnoSubject struct {
	key string `query:"key"`
	value string `query:"value"`
}
