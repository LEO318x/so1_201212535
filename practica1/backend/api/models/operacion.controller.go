package models

type Operacion struct {
	Num1 int64
	Num2 int64
}

type ResultadoOperacion struct {
	Resultado int64 `json:"resultado"`
}
