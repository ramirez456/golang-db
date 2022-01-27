package storage

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
	"time"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/go-sql-driver/mysql"
)
//sigleton creacion
var(
	db *sql.DB
	once sync.Once
)
func loadEnv(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func NewPostgresDB(){
	loadEnv()
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	url := "postgres://"+dbUser+":"+dbPassword+"@localhost:5432/"+dbDatabase+"?sslmode=disable"
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", url)
		if err != nil {
			log.Fatalf("can't open db: %v",err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't open db: %v",err)
		}
		fmt.Println("Conectado a postgres")
	})
}

func NewMySQLDB(){
	loadEnv()
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	url := dbUser+":"+dbPassword+"@/"+dbDatabase
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", url)
		if err != nil {
			log.Fatalf("can't open db: %v",err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't open db: %v",err)
		}
		fmt.Println("Conectado a postgres")
	})
}

func Pool() *sql.DB{
	return db
}

func stringToNull(s string) sql.NullString{
	null := sql.NullString{String: s}
	if null.String != ""{
		null.Valid = true
	}
	return null
}
func timeToNull(t  time.Time) sql.NullTime{
	null := sql.NullTime{Time: t}
	if null.Time.IsZero(){
		null.Valid = true
	}
	return null
}