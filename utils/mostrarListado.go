package utils

import (
	"os"

	"github.com/galileoluna/challenge1/models"
	"github.com/jedib0t/go-pretty/table"
)

func MostrarListado(cursantes []models.CantidadAsistencia) {
	t := table.NewWriter()
	t.SetTitle("Listado de Cursos")
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Nombre", "Cantidad Alumnos"})
	for i := 0; i < len(cursantes); i++ {
		t.AppendRow([]interface{}{cursantes[i].ID, cursantes[i].NOMBRE, cursantes[i].CANTIDAD})
	}
	t.Render()
}
