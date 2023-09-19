package paciente

//import "time"

type Paciente struct {
	ID        int       `json:"id"`
	Nombre    string    `json:"nombre"`
	Apellido  string    `json:"apellido"`
	Domicilio string    `json:"domicilio"`
	DNI       int       `json:"dni"`
	FechaAlta string `json:"fecha_alta"`
}

type RequestPaciente struct {
	Nombre    string    `json:"nombre"`
	Apellido  string    `json:"apellido"`
	Domicilio string    `json:"domicilio"`
	DNI       int       `json:"dni"`
	FechaAlta string `json:"fecha_alta"`
}

type RequestUpdatePacienteSubject struct {
	key   string `query:"key"`
	value string `query:"value"`
}
