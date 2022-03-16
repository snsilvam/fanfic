package storage

import (
	"database/sql"
	"fmt"

	"github.com/snsilvam/fanfic/pkg/categorys"
)

const (
	//This const created the fields to table Category
	mySQLMigrateCategorys = `CREATE TABLE IF NOT EXISTS category( 
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY, 
		name VARCHAR(30) NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP
		
	)`
	mySQLCreateProduct = `INSERT INTO category(name, created_at) VALUES(?, ?)`
)

//db *sql.DB declare one variable db, this variable we help to created the user table in data base
type MySQLCategorys struct {
	db *sql.DB
}

//NewMySQLCategory is the db's constructor
func NewMySQLCategorys(db *sql.DB) *MySQLCategorys {
	return &MySQLCategorys{db}
}

//The function MySQLCategory's Migrate, allows create the table in mysql
func (c *MySQLCategorys) Migrate() error {
	//The db package's prepare method, allows create a new table into our data base
	//What is stmt?
	stmt, err := c.db.Prepare(mySQLMigrateCategorys)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migrate successfully")
	return nil
}

//The function Create implement the interface category.storage
func (c *MySQLCategorys) Create(m *categorys.ModelCategorys) error {
	stmt, err := c.db.Prepare(mySQLCreateProduct)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(
		m.Name,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = uint(id)
	fmt.Printf("Category successfuly created, with id: %d", m.ID)
	return nil
}
