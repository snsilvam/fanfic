package main

import (
	"fmt"
	"log"

	"github.com/snsilvam/fanfic/pkg/categorys"
	"github.com/snsilvam/fanfic/storage"
)

func main() {
	fmt.Println("Hello team of fanfic")
	storage.NewMySQLDB()
	//Create a new category into data base mysql
	storageCategorys := storage.NewMySQLCategorys(storage.Pool())
	serviceCategorys := categorys.NewService(storageCategorys)
	m := &categorys.ModelCategorys{
		Name: "Humor",
	}
	if err := serviceCategorys.Create(m); err != nil {
		log.Fatalf("Category.Create: %v", err)
	}
	fmt.Printf("%+v\n", m)
}
