package storage

import (
	"database/sql"
	"fmt"
	"github.com/ramirez456/go-db/pkg/invoiceheader"
)
const(
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_headers_id_pk PRIMARY KEY(id)
	)`
	psqlCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES($1) RETURNING id, created_at`
)
type PsqlInvoiceHeader struct {
	db *sql.DB
}

func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db  }
}
 func (p *PsqlInvoiceHeader) Migrate() error{
 	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
	 if err != nil {
		 return err
	 }
	 defer stmt.Close()

 	_,err =  stmt.Exec()
	 if err != nil {
		 return err
	 }
	 fmt.Println("Migrations execute de Invoice Header")
	 return nil
 }

 func (p *PsqlInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error{

 	stmt, err := tx.Prepare(psqlCreateInvoiceHeader)

	 if err != nil {
		 return err
	 }

	 defer stmt.Close()

 	return stmt.QueryRow(m.Client).Scan(&m.ID, &m.CreatedAt)
 }