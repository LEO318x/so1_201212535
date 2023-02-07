package models

type Tag struct {
	Id         int64  `json:"id"`
	Numero1    int64  `json:"num1"`
	Numero2    int64  `json:"num2"`
	Operacion  string `json:"operador"`
	Resultado  int64  `json:"resultado"`
	Fecha_hora string `json:"fecha_hora"`
}
