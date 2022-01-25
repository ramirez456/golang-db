package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ramirez456/go-db/pkg/product"
	"github.com/ramirez456/go-db/pkg/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()
	/*
	//crer las bases de datos

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil{
		log.Fatalf("Producto.Migrate %v", err)
	}

	storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

	if err := serviceInvoiceHeader.Migrate(); err != nil{
		log.Fatalf("InvoiceHeader.Migrate %v", err)
	}

	storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("InvoiItem.Migrate %v", err)
	}
	*/
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	m := &product.Model{
		Name: "curso de Base de datos con Go",
		Price: 50,
		Observations: "Este curso",
	}
	if err := serviceProduct.Create(m); err != nil{
		log.Fatalf("Producto.Create %v", err)
	}
	fmt.Printf("%+v\n",m)
}