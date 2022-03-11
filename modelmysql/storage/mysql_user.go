package storage

import (
	"database/sql"
	"fmt"
)

const (
	//This const created the fields to table Users
	mySQLMigrateUsers = `CREATE TABLE IF NOT EXISTS Users( 
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY, 
		name VARCHAR(40) NULL,
		age INT NULL,
		gender CHAR(8) NULL DEFAULT'another',
		nickname VARCHAR(40) NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP
		
	)`
)

//db *sql.DB declare one variable db, this variable we help to created the user table in data base
type MySQLUsers struct {
	db *sql.DB
}

//NewMySQLUsers is the db's constructor
func NewMySQLUsers(db *sql.DB) *MySQLUsers {
	return &MySQLUsers{db}
}

//The function MySQLUsers's Migrate, allows create the table in mysql
func (p *MySQLUsers) Migrate() error {
	//The db package's prepare method, allows create a new table into our data base
	//What is stmt?
	stmt, err := p.db.Prepare(mySQLMigrateUsers)
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
