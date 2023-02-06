package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Conexion() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: No se puede cargar el archivo .env")
	}

	fmt.Println(os.Getenv("USERDB"))
	fmt.Println(os.Getenv("PASSDB"))
	fmt.Println(os.Getenv("NAMEDB"))
	fmt.Println(os.Getenv("HOSTDB"))
	fmt.Println(os.Getenv("PORTDB"))
}
