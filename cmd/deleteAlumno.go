/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"strconv"

	"github.com/galileoluna/challenge1/data"
	"github.com/galileoluna/challenge1/models"

	"github.com/spf13/cobra"
)

// deleteAlumnoCmd represents the deleteAlumno command
var deleteAlumnoCmd = &cobra.Command{
	Use:   "deleteAlumno",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var alumnoEliminado models.Alumno
		var dni int
		dni, _ = strconv.Atoi(args[2])
		alumnoEliminado = models.NewAlumno(args[0], args[1], dni)
		data.DeleteAlumno(alumnoEliminado)
	},
}

func init() {
	rootCmd.AddCommand(deleteAlumnoCmd)
}
