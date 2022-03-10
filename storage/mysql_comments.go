package storage

import (
	"database/sql"
	"fmt"
)

const (
	//This const created the fields to table post
	mySQLMigrateComments = `CREATE TABLE IF NOT EXISTS comments( 
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,  
		comment TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP,
		user_id INT NOT NULL,
		post_id INT NOT NULL,
		CONSTRAINT commets_users FOREIGN KEY (user_id) REFERENCES fanfic.users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
		CONSTRAINT posts_posts FOREIGN KEY (post_id) REFERENCES fanfic.posts (id) ON DELETE RESTRICT ON UPDATE CASCADE
	)`
)

//This is
type MySQLComments struct {
	db *sql.DB
}

//
func NewMySQLComments(db *sql.DB) *MySQLComments {
	return &MySQLComments{db}
}

//The method prepare of db package, allows create a new table into our data base
func (p *MySQLComments) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateComments)
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
