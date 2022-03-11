package storage

import (
	"database/sql"
	"fmt"
)

const (
	//This const created the fields to table Tags
	mySQLMigrateTags = `CREATE TABLE IF NOT EXISTS tags( 
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY, 
		tag_name VARCHAR(40) NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP
		
	)`
)

//db *sql.DB declare one variable db, this variable we help to created the tags table in data base
type MySQLTags struct {
	db *sql.DB
}

//NewMySQLTags is the db's constructor
func NewMySQLTags(db *sql.DB) *MySQLTags {
	return &MySQLTags{db}
}

//The function MySQLTags's Migrate, allows create the table in mysql
func (p *MySQLTags) Migrate() error {
	//The db package's prepare method, allows create a new table into our data base
	//What is stmt?
	stmt, err := p.db.Prepare(mySQLMigrateTags)
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
