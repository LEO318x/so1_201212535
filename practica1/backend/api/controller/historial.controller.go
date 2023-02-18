package controller

import (
	"api/db"
	"api/models"
	"fmt"
	"log"
	"time"
)

func GetHistorial() ([]models.Tag, error) {
	historial := []models.Tag{}
	bd, err := db.Conexion()
	if err != nil {
		log.Printf("Error con la base de datos" + err.Error())
	} else {
		err = bd.Ping()
		if err != nil {
			log.Printf("Error al conectarse a la base de datos. Por favor verifica las credenciales, el error es: " + err.Error())
		}
	}

	res, err := bd.Query("SELECT * FROM so1.log	ORDER BY id DESC")
	for res.Next() {
		var tag models.Tag
		// Para cada fila, escaneamos el resultamo y lo agregamos en nuestra estructura
		err = res.Scan(&tag.Id, &tag.Numero1, &tag.Numero2, &tag.Operacion, &tag.Resultado, &tag.Fecha_hora)

		if err != nil {
			panic(err.Error()) // Error
		}
		// Agregamos cada resultado en nuestro arreglo
		historial = append(historial, tag)

	}

	return historial, nil
}

func GetHistorial2() ([]models.Tag, error) {
	historial := []models.Tag{}
	bd, err := db.Conexion()
	if err != nil {
		log.Printf("Error con la base de datos" + err.Error())
	} else {
		err = bd.Ping()
		if err != nil {
			log.Printf("Error al conectarse a la base de datos. Por favor verifica las credenciales, el error es: " + err.Error())
		}
	}

	hoy := time.Now()
	pquery := fmt.Sprintf("%s%s%s%s%s", "SELECT * FROM so1.log WHERE fecha_hora >= ", "\"", hoy.Format("2006-01-02"), "\"", " ORDER BY id DESC")
	fmt.Println(pquery)
	res, err := bd.Query(pquery)
	for res.Next() {
		var tag models.Tag
		// Para cada fila, escaneamos el resultamo y lo agregamos en nuestra estructura
		err = res.Scan(&tag.Id, &tag.Numero1, &tag.Numero2, &tag.Operacion, &tag.Resultado, &tag.Fecha_hora)

		if err != nil {
			panic(err.Error()) // Error
		}
		// Agregamos cada resultado en nuestro arreglo
		historial = append(historial, tag)

	}

	return historial, nil
}
