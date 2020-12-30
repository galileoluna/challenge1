package utils

import (
	"os"

	"github.com/galileoluna/challenge1/models"

	"github.com/jedib0t/go-pretty/table"
)

func MostrarAlumnos(cursantes []models.Alumno) {

	t := table.NewWriter()
	t.SetTitle("Alumnos")

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Nombre", "Apellido", "DNI"})

	for i := 0; i < len(cursantes); i++ {

		t.AppendRow([]interface{}{i, cursantes[i].Nombre, cursantes[i].Apellido, cursantes[i].Dni})
	}
	t.Render()
}
