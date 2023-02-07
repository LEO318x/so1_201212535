package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Tag struct {
	Num1 int64
	Num2 int64
}

func Conexion() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: No se puede cargar el archivo .env")
	}

	db, err := sql.Open("mysql", os.Getenv("USERDB")+":"+os.Getenv("PASSDB")+"@tcp("+os.Getenv("HOSTDB")+":"+os.Getenv("PORTDB")+")/"+os.Getenv("NAMEDB"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM log")

	for res.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = res.Scan(&tag.Num1, &tag.Num2)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Println(tag.Num1)
		log.Println(tag.Num2)
	}

}
