package odontologo

type Odontologo struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Matricula string `json:"matricula"`
}

type RequestOdontologo struct {
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Matricula string `json:"matricula"`
}

type RequestUpdateOdontologoSubject struct {
	Key string `form:"key" json:"key"`
	Value string `form:"value" json:"value"`
}