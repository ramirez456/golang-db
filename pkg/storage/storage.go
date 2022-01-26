package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)
//sigleton creacion
var(
	db *sql.DB
	once sync.Once
)

func NewPostgresDB(){
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://maxhoustonramirezmartel:ebcnemj987@localhost:5432/mongo-go?sslmode=disable")
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