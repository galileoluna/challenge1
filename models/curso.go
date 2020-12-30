package models

type Curso struct {
	ID     int
	NOMBRE string
}

func NewCurso(id_input int, nombre_input string) Curso {
	return Curso{
		ID:     id_input,
		NOMBRE: nombre_input,
	}
}
