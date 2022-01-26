package main

import (
	_ "github.com/lib/pq"
	"github.com/ramirez456/go-db/pkg/invoice"
	"github.com/ramirez456/go-db/pkg/invoiceheader"
	"github.com/ramirez456/go-db/pkg/invoiceitem"
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
	/*
	// Para crear productos

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
	*/
	/*
	//Get All

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	ms, err := serviceProduct.GetAll()
	if err != nil{
		log.Fatalf("Producto.Get All %v", err)
	}
	fmt.Println(ms)*/
/*
//Get by id
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m, err := serviceProduct.GetByID(1)
	switch {
		case errors.Is(err, sql.ErrNoRows):
			fmt.Println("No hay producto con este ID")
		case err != nil:
			log.Fatalf("product get by ID %v",err)
		default:
			fmt.Println(m)
	}
*/
/*
 // Update

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	m := &product.Model{
		ID: 2,
		Name: "sin fecha",
	}
	err := serviceProduct.Update(m)
	if err != nil {
		log.Fatalf("product update by id %v",err)
	}
	fmt.Println("Create success") /*

 */
/*
//Delete

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(2)
	if err != nil {
		log.Fatalf("Product delete by id %v",err)
	}
	fmt.Println("Delete success")
*/
	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewPsqlInvoiceItem( storage.Pool())

	storageInvoice := storage.NewPsqlInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
		)
	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Wings Houston Ramirez Martel",
		},
		Items: invoiceitem.Models{
			 &invoiceitem.Model{ProductID: 1},
			 &invoiceitem.Model{ProductID: 99},
		},
	}
	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}
}