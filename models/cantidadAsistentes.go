package models

type CantidadAsistencia struct {
	ID       int
	NOMBRE   string
	CANTIDAD int
}

func NewCantidadAsistencia(id_input int, nombre_input string, cantidad_input int) CantidadAsistencia {
	return CantidadAsistencia{
		ID:       id_input,
		NOMBRE:   nombre_input,
		CANTIDAD: cantidad_input,
	}
}
