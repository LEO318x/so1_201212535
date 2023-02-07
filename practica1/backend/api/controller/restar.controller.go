package controller

import (
	"api/db"
	"api/models"
	"log"
	"time"
)

func RealizarResta(num1 int64, num2 int64) []models.ResultadoOperacion {
	resultado := []models.ResultadoOperacion{}
	bd, err := db.Conexion()
	if err != nil {
		log.Printf("Error con la base de datos" + err.Error())
	} else {
		err = bd.Ping()
		if err != nil {
			log.Printf("Error al conectarse a la base de datos. Por favor verifica las credenciales, el error es: " + err.Error())
		}
	}
	res := num1 - num2
	fechahora := time.Now().Format("2006-01-02 15:04:05")
	_, err = bd.Exec("INSERT INTO log (numero1, numero2, operacion, resultado, fecha_hora) VALUES (?, ?, ?, ?, ?)", num1, num2, "-", res, fechahora)
	if err != nil {
		log.Printf("Hubo un error al insertar el registro en el historial")
	}
	var tag models.ResultadoOperacion
	tag.Resultado = res
	resultado = append(resultado, tag)
	return resultado
}
