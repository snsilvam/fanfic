package storage

import (
	"database/sql"
	"fmt"
)

const (
	//This const created the fields to table post
	mySQLMigrateUser = `CREATE TABLE IF NOT EXISTS Users( 
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY, 
		name VARCHAR(50) NOT NULL, 
		age INT NOT NULL,
		gender  VARCHAR(25),
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP
		
	)`
)

//This is
type MySQLUser struct {
	db *sql.DB
}

//
func NewMySQLUser(db *sql.DB) *MySQLUser {
	return &MySQLUser{db}
}

//The method prepare of db package, allows create a new table into our data base
func (p *MySQLUser) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateUser)
	if err != nil {
		return err
	}
	//What is stmt?
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migrate successfully")
	return nil
}
