package data

import (
	"fmt"

	"github.com/galileoluna/challenge1/models"
)

func ExisteCurso(curso models.Curso) (int, []models.Curso, error) {
	Cursos := []models.Curso{}
	var cant int
	rows, err := db.Query(`SELECT * FROM Curso where idcurso = $1;`, curso.NOMBRE)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()
	for rows.Next() {
		cant++
		var id_DB int
		var nombre_DB string
		var CursoActual models.Curso
		err = rows.Scan(&id_DB, &nombre_DB)
		if err != nil {
			return 0, Cursos, err
		}
		CursoActual = models.NewCurso(id_DB, nombre_DB)
		Cursos = append(Cursos, CursoActual)
	}
	return cant, Cursos, err
}
func InsertCurso(nuevoCurso models.Curso) (int, error) {
	//Create
	var CursoID int
	var existe int

	existe, _, _ = ExisteCurso(nuevoCurso)

	if existe == 0 {
		err := db.QueryRow(`INSERT INTO CURSO(NOMBRE) VALUES($1) RETURNING IDCurso`, nuevoCurso.NOMBRE).Scan(&CursoID)

		if err != nil {
			return 0, err
		}
		fmt.Printf("Last inserted ID: %v\n", CursoID)
		return CursoID, err
	} else {
		fmt.Println("Ya existe ese curso...")
		return 0, nil
	}

}

func GetCursantesCurso(cursoABuscar models.Curso) ([]string, error) {
	Cursantes := []string{}
	rows, err := db.Query(`SELECT NOMBRE FROM ALUMNO INNER JOIN ASISTENCIA ON ALUMNO.DNIALUMNO = ASISTENCIA.DNIALUMNO WHERE ASISTENCIA.NOMBRECLASE=$1;`, cursoABuscar.NOMBRE)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var nombre_DB string
		err = rows.Scan(&nombre_DB)
		if err != nil {
			return Cursantes, err
		}
		nombre_DB = nombre_DB
		Cursantes = append(Cursantes, nombre_DB)
	}
	return Cursantes, err

}

func GetCursos() ([]models.Curso, error) {
	Cursos := []models.Curso{}
	var cant int

	rows, err := db.Query(`SELECT * FROM Curso;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		cant++
		var id_DB int
		var nombre_DB string
		var CursoActual models.Curso

		err = rows.Scan(&id_DB, &nombre_DB)
		if err != nil {
			return Cursos, err
		}
		CursoActual = models.NewCurso(id_DB, nombre_DB)
		Cursos = append(Cursos, CursoActual)
	}
	return Cursos, err
}
func YaEsCursante(nuevoCursante models.Alumno, curso models.Curso) bool {
	//Retrieve
	var nombre string

	var dnialumno int32

	err := db.QueryRow(`SELECT * FROM ASISTENCIA where DNIALUMNO = $1 and NOMBRECLASE=$2`, nuevoCursante.Dni, curso.NOMBRE).Scan(&dnialumno, &nombre)
	if err == nil {
		return false
	}
	return true
}
func CantidadAlumnosPorCurso() ([]models.CantidadAsistencia, error) {
	Listado := []models.CantidadAsistencia{}

	rows, err := db.Query(`SELECT DISTINCT NOMBRE, COUNT(NOMBRE) FROM CURSO LEFT JOIN ASISTENCIA ON CURSO.NOMBRE=ASISTENCIA.NOMBRECLASE GROUP BY CURSO.NOMBRE;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var i int
	for rows.Next() {

		var nombre_DB string
		var cant_DB int
		var CursoActual models.CantidadAsistencia
		err = rows.Scan(&nombre_DB, &cant_DB)
		if err != nil {
			return Listado, err
		}
		i++
		CursoActual = models.NewCantidadAsistencia(i, nombre_DB, cant_DB)
		Listado = append(Listado, CursoActual)
	}
	return Listado, err
}
func AgregarAlumnoACurso(nuevoCursante models.Alumno, curso models.Curso) (int, error) {
	//Create
	var IdAsistencia int
	var existe bool
	existe = YaEsCursante(nuevoCursante, curso)

	if existe == true {
		err := db.QueryRow(`INSERT INTO ASISTENCIA(DNIALUMNO,NOMBRECLASE) VALUES($1,$2) RETURNING IDASISTENCIA`, nuevoCursante.Dni, curso.NOMBRE).Scan(&IdAsistencia)

		if err != nil {
			return 0, err
		}

		fmt.Printf("Last inserted ID: %v\n", IdAsistencia)
		return IdAsistencia, err
	} else {
		fmt.Println("Ya existe ese")
		return 0, nil
	}

}
