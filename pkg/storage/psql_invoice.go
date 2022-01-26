package storage

import (
	"database/sql"
	"fmt"
	"github.com/ramirez456/go-db/pkg/invoice"
	"github.com/ramirez456/go-db/pkg/invoiceheader"
	"github.com/ramirez456/go-db/pkg/invoiceitem"
)

type PsqlInvoice struct {
	db *sql.DB
	storageHeader invoiceheader.Storage
	storageItems invoiceitem.Storage
}

func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice{
	return &PsqlInvoice{
		db: db,
		storageHeader: h,
		storageItems: i,
	}
}

func (p *PsqlInvoice) Create(m *invoice.Model) error{
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	if err:= p.storageHeader.CreateTx(tx, m.Header); err != nil{
		tx.Rollback()
		return err
	}
	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil{
		tx.Rollback()
		return err
	}
	fmt.Printf("Items creados %d \n", len(m.Items))
	return tx.Commit()
}