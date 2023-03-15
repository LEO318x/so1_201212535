package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Conexion() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: No se puede cargar el archivo .env")
	}

	return sql.Open("mysql", os.Getenv("USERDB")+":"+os.Getenv("PASSDB")+"@tcp("+os.Getenv("HOSTDB")+":"+os.Getenv("PORTDB")+")/"+os.Getenv("NAMEDB"))
	//return sql.Open("mysql", "root:root@tcp(localhost:3306)/sopes")
}
