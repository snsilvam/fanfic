package storage

import (
	"database/sql"
	"fmt"
)

const (
	//This const created the fields to table posts
	mySQLCreatePosts = `CREATE TABLE IF NOT EXISTS posts( 
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY, 
		title VARCHAR(150) NOT NULL, 
		body TEXT NOT NULL,
		view INT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP,
		user_id INT NOT NULL,
		category_id INT NOT NULL,
		CONSTRAINT posts_users FOREIGN KEY (user_id) REFERENCES fanfic.users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
		CONSTRAINT posts_categorys FOREIGN KEY (category_id) REFERENCES fanfic.category (id) ON DELETE RESTRICT ON UPDATE CASCADE
	)`
)

//This is
type MySQLPosts struct {
	db *sql.DB
}

//
func NewMySQLPosts(db *sql.DB) *MySQLPosts {
	return &MySQLPosts{db}
}

//The method prepare of db package, allows create a new table into our data base
func (p *MySQLPosts) Migrate() error {
	stmt, err := p.db.Prepare(mySQLCreatePosts)
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
