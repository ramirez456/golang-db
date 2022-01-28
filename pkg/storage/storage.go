package storage

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ramirez456/go-db/pkg/product"
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

type  Driver string

const(
	MySQL Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

func loadEnv(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func newPostgresDB(){
	loadEnv()
	dbUser := os.Getenv("PSQL_DB_USERNAME")
	dbPassword := os.Getenv("PSQL_DB_PASSWORD")
	dbDatabase := os.Getenv("PSQL_DB_DATABASE")
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

func newMySQLDB(){
	loadEnv()
	dbUser := os.Getenv("MYSQL_DB_USERNAME")
	dbPassword := os.Getenv("MYSQL_DB_PASSWORD")
	dbDatabase := os.Getenv("MYSQL_DB_DATABASE")
	dbHost := os.Getenv("MYSQL_DB_HOST")
	dbPort  := os.Getenv("MYSQL_DB_PORT")

	url := dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbDatabase+"?parseTime=true"
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", url)
		if err != nil {
			log.Fatalf("can't open db: %v",err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't open db: %v",err)
		}
		fmt.Println("Conectado a mysql")
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

func DAOProduct(driver Driver)(product.Storage, error){
	switch driver {
	case Postgres:
		return newPsqlProduct(db), nil
	case MySQL:
		return newMySQLProduct(db), nil
	default:
		return nil, fmt.Errorf("driver not implement")
	}
}