package storage

import (
	"database/sql"
	"fmt"
)

const (
	//This const created the fields to table post
	mySQLCreatePost = `CREATE TABLE IF NOT EXISTS post( 
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY, 
		title VARCHAR(50) NOT NULL, 
		body VARCHAR(200) NOT NULL,
		view INT,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP
		CONSTRAINT post_user_id_fk FOREIGN KEY (user_id) REFERENCES user (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
)

//This is
type MySQLPost struct {
	db *sql.DB
}

//
func NewMySQLPost(db *sql.DB) *MySQLPost {
	return &MySQLPost{db}
}

//The method prepare of db package, allows create a new table into our data base
func (p *MySQLPost) Migrate() error {
	stmt, err := p.db.Prepare(mySQLCreatePost)
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
