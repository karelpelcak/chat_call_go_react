package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Nelze načíst .env, použij default proměnné prostředí")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	log.Println(dsn)

	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln("Chyba při připojení k DB:", err)
	}

	fmt.Println("Připojeno k databázi!")

	var version string
	err = DB.Get(&version, "SELECT version()")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("PostgreSQL verze:", version)
}
