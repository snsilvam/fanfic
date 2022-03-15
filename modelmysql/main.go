package main

import (
	"fmt"

	"github.com/snsilvam/fanfic/storage"
	//"log"
	//"github.com/snsilvam/fanfic/pkg/categorys"
	//"github.com/snsilvam/fanfic/pkg/post"
	//"github.com/snsilvam/fanfic/pkg/comments"
	//"github.com/snsilvam/fanfic/pkg/posts"
)

func main() {
	fmt.Println("Hello team of fanfic")
	storage.NewMySQLDB()

	//Add the table users into Data base mysql
	/* storageUsers := storage.NewMySQLUsers(storage.Pool())
	serviceUsers := users.NewService(storageUsers)
	if err := serviceUsers.Migrate(); err != nil {
		log.Fatalf("Error in post.Migrate: %v", err)
	} */
	//Add the table categorys into data base mysql
	/* storageCategorys := storage.NewMySQLCategorys(storage.Pool())
	serviceCategorys := categorys.NewService(storageCategorys)
	if err := serviceCategorys.Migrate(); err != nil {
		log.Fatalf("Error in post.Migrate: %v", err)
	} */
	//Add the table categorys into data base mysql
	/* storageTags := storage.NewMySQLTags(storage.Pool())
	serviceTags := tags.NewService(storageTags)
	if err := serviceTags.Migrate(); err != nil {
		log.Fatalf("Error in post.Migrate: %v", err)
	} */
	//Add the table posts into Data base
	/* storagePosts := storage.NewMySQLPosts(storage.Pool())
	servicePosts := posts.NewService(storagePosts)

	if err := servicePosts.Migrate(); err != nil {
		log.Fatalf("Error in post.Migrate: %v", err)
	} */
	//Add the table comments into Data base
	/* storageComments := storage.NewMySQLComments(storage.Pool())
	serviceComments := comments.NewService(storageComments)

	if err := serviceComments.Migrate(); err != nil {
		log.Fatalf("Error in post.Migrate: %v", err)
	} */
}
