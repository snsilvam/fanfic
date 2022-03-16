## Migrates to mysql of fanfic
	In order to migrate the tables, please add the following code in import 
	<!-- 	//"log"
		//"github.com/snsilvam/fanfic/pkg/categorys"
		//"github.com/snsilvam/fanfic/pkg/post"
		//"github.com/snsilvam/fanfic/pkg/comments"
		//"github.com/snsilvam/fanfic/pkg/posts" -->
    //Add the table users into Data base mysql, migrate the Users's table 
    <!-- storageUsers := storage.NewMySQLUsers(storage.Pool())
	serviceUsers := users.NewService(storageUsers)
	if err := serviceUsers.Migrate(); err != nil {
		log.Fatalf("Error in post.Migrate: %v", err)
	}  -->

	//Add the table categorys into data base mysql, migrate the  category's table 
	<!-- storageCategorys := storage.NewMySQLCategorys(storage.Pool())
	serviceCategorys := categorys.NewService(storageCategorys)
	if err := serviceCategorys.Migrate(); err != nil {
		log.Fatalf("Error in post.Migrate: %v", err)
	}  -->

	//Add the table tags into data base mysql, migrate the tags's table 
	<!-- storageTags := storage.NewMySQLTags(storage.Pool())
	serviceTags := tags.NewService(storageTags)
	if err := serviceTags.Migrate(); err != nil {
		log.Fatalf("Error in post.Migrate: %v", err)
	} -->

	//Add the table posts into data base mysql, migrate the post's table
	<!-- storagePosts := storage.NewMySQLPosts(storage.Pool())
	servicePosts := posts.NewService(storagePosts)
	if err := servicePosts.Migrate(); err != nil {
		log.Fatalf("Error in post.Migrate: %v", err)
	} -->

	//Add the table comments into data base mysql, migrate the comments's table
	<!-- storageComments := storage.NewMySQLComments(storage.Pool())
	serviceComments := comments.NewService(storageComments)
	if err := serviceComments.Migrate(); err != nil {
		log.Fatalf("Error in post.Migrate: %v", err)
	} -->