package data

import (
	"fmt"

	"github.com/galileoluna/challenge1/models"
)

func ExisteAlumno(nuevoAlumno models.Alumno) (int, []models.Alumno, error) {
	//Retrieve
	var cant int

	alumnos := []models.Alumno{}

	rows, err := db.Query(`SELECT NOMBRE,APELLIDO,DNIALUMNO FROM ALUMNO where DNIALUMNO = $1`, nuevoAlumno.Dni)
	if err == nil {
		return 0, nil, nil
	}
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()
	for rows.Next() {
		cant++
		var NOMBRE string
		var APELLIDO string
		var DNIALUMNO int
		var alumnoActual models.Alumno

		err = rows.Scan(&NOMBRE, &APELLIDO, &DNIALUMNO)
		if err != nil {
			return 0, alumnos, err
		}
		alumnoActual = models.NewAlumno(NOMBRE, APELLIDO, DNIALUMNO)
		alumnos = append(alumnos, alumnoActual)
	}
	return cant, alumnos, err
}

func InsertAlumno(nuevoAlumno models.Alumno) (int, error) {
	//Creamos
	//var alumnoDNI

	var dni_al int
	var existeAlumno int

	existeAlumno, _, _ = ExisteAlumno(nuevoAlumno)

	if existeAlumno == 0 {
		err := db.QueryRow(`INSERT INTO ALUMNO(NOMBRE,APELLIDO,DNIALUMNO) VALUES($1,$2,$3) RETURNING DNIALUMNO`, nuevoAlumno.Nombre, nuevoAlumno.Apellido, nuevoAlumno.Dni).Scan(&dni_al)

		if err != nil {
			return 0, err
		}

		fmt.Printf("Last inserted ID: %v\n", dni_al)
		return dni_al, err

	} else {

		fmt.Println("Ya existe ese alumno...")
		return 0, nil

	}

}

func DeleteAlumno(viejoAlumno models.Alumno) (int, error) {
	//Delete
	//var existeAlumno bool
	//existeAlumno = ExisteAlumno(viejoAlumno)

	var noExisteAlumno int

	noExisteAlumno, _, _ = ExisteAlumno(viejoAlumno)

	if noExisteAlumno == 0 {
		fmt.Println("No existe ese usuario, pruebe nuevamente")
		return 0, nil
	} else {

		fmt.Println("Entra")
		res, err := db.Exec(`DELETE FROM ALUMNO WHERE DNIALUMNO = $1`, viejoAlumno.Dni)
		if err != nil {
			return 0, err
		}
		rowsDeleted, err := res.RowsAffected()
		if err != nil {
			return 0, err
		}

		return int(rowsDeleted), nil

	}

}

func GetAlumnos() ([]models.Alumno, error) {
	alumnos := []models.Alumno{}
	rows, err := db.Query(`SELECT NOMBRE,APELLIDO,DNIALUMNO FROM ALUMNO`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var NOMBRE string
		var APELLIDO string
		var DNIALUMNO int
		var alumnoActual models.Alumno

		err = rows.Scan(&NOMBRE, &APELLIDO, &DNIALUMNO)
		if err != nil {
			return alumnos, err
		}
		alumnoActual = models.NewAlumno(NOMBRE, APELLIDO, DNIALUMNO)
		alumnos = append(alumnos, alumnoActual)
	}
	return alumnos, err
}
