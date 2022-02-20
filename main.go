package main

import (
	"fmt"

	"github.com/snsilvam/fanfic/storage"
)

func main() {
	fmt.Println("Hello team of fanfic")
	storage.NewMySQLDB()
}
