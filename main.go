package main

import (
	_ "github.com/lib/pq"
	"github.com/ramirez456/go-db/pkg/invoiceheader"
	"github.com/ramirez456/go-db/pkg/invoiceitem"
	"github.com/ramirez456/go-db/pkg/product"
	"github.com/ramirez456/go-db/pkg/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()

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
}