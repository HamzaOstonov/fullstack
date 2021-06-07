package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}
	Method = &Postgres{db: db}
}
