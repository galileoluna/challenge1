package utils

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func MostrarCursantes(cursantes []string, nombre_curso string) {

	t := table.NewWriter()
	t.SetTitle(nombre_curso)

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Nombre"})

	for i := 0; i < len(cursantes); i++ {

		t.AppendRow([]interface{}{i, cursantes[i]})
	}
	t.Render()
}
