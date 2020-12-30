package utils

import (
	"os"

	"github.com/galileoluna/challenge1/models"
	"github.com/jedib0t/go-pretty/table"
)

func MostrarCursos(cursantes []models.Curso) {

	t := table.NewWriter()
	t.SetTitle("Cursos")

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Nombre"})

	for i := 0; i < len(cursantes); i++ {

		t.AppendRow([]interface{}{cursantes[i].ID, cursantes[i].NOMBRE})
	}
	t.Render()
}
