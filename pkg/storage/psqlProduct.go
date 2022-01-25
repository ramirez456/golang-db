package storage

import (
	"database/sql"
	"fmt"
	"github.com/ramirez456/go-db/pkg/product"
)
const(
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price float NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY(id)
	)`
	psqlCreateProduct =  `INSERT INTO 
	products(name, observations, price, created_at)
	VALUES($1, $2, $3, $4) RETURNING id`;
)
type PsqlProduct struct {
	db *sql.DB
}

func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db  }
}
 func (p *PsqlProduct) Migrate() error{
 	stmt, err := p.db.Prepare(psqlMigrateProduct)
	 if err != nil {
		 return err
	 }
	 defer stmt.Close()

 	_,err =  stmt.Exec()
	 if err != nil {
		 return err
	 }
	 fmt.Println("Migration execute Product")
	 return nil
 }
 func (p *PsqlProduct) Create(m *product.Model) error {

 	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}

	defer stmt.Close()
 	err = stmt.QueryRow(m.Name, stringToNull(m.Observations), m.Price, m.CreatedAt).Scan(&m.ID)
	 if err != nil {
		 return err
	 }
	 fmt.Println("Create execute Product")
	 return nil
 }