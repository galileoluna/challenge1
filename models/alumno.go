package models

type Alumno struct {
	Nombre   string
	Apellido string
	Dni      int
}

func NewAlumno(nombre_input string, apellido_input string, dni_input int) Alumno {

	return Alumno{
		Nombre:   nombre_input,
		Apellido: apellido_input,
		Dni:      dni_input,
	}
}
