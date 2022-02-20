package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewMySQLDB() {
	//The function Do receive a anonymously function
	once.Do(func() {
		var err error
		// The first argument of Open is the name of driver ("mysql")
		// and the second argument is the string of connection, where put the credential of acces to DB
		db, err = sql.Open("mysql", "root:7752930*Nic0@tcp(localhost:3306)/fanfic")
		if err != nil {
			log.Fatalf("The app can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("The app can't do ping: %v", err)
		}

		fmt.Println("successfully connection")
	})
}
