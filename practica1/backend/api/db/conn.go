package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Conexion() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: No se puede cargar el archivo .env")
	}

	//return sql.Open("mysql", os.Getenv("USERDB")+":"+os.Getenv("PASSDB")+"@tcp("+os.Getenv("HOSTDB")+":"+os.Getenv("PORTDB")+")/"+os.Getenv("NAMEDB"))
	return sql.Open("mysql", "root:root@tcp(192.168.3.7:3306)/so1")
}
