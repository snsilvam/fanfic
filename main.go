package main

import (
	"fmt"
	"log"

	"github.com/snsilvam/fanfic/pkg/wattpad/comment"

	//"github.com/snsilvam/fanfic/pkg/wattpad/post"
	"github.com/snsilvam/fanfic/storage"
)

func main() {
	fmt.Println("Hello team of fanfic")
	storage.NewMySQLDB()

	//Add a table into Data base
	/*storagePost := storage.NewMySQLPost(storage.Pool())
	servicePost := post.NewService(storagePost)

	if err := servicePost.Migrate(); err != nil {
		log.Fatalf("Error in post.Migrate: %v", err)
	}*/
	storagePost := storage.NewMySQLComment(storage.Pool())
	servicePost := comment.NewService(storagePost)

	if err := servicePost.Migrate(); err != nil {
		log.Fatalf("Error in post.Migrate: %v", err)
	}

}
