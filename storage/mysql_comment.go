package storage

import (
	"database/sql"
	"fmt"
)

const (
	//This const created the fields to table post
	mySQLMigrateComment = `CREATE TABLE IF NOT EXISTS comments( 
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY, 
		title VARCHAR(50) NOT NULL, 
		body VARCHAR(200) NOT NULL,
		view INT,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP
	)`
)

//This is
type MySQLComment struct {
	db *sql.DB
}

//
func NewMySQLComment(db *sql.DB) *MySQLComment {
	return &MySQLComment{db}
}

//The method prepare of db package, allows create a new table into our data base
func (p *MySQLComment) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateComment)
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
