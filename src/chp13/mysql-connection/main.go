package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	// Import the mysql driver
	_ "github.com/go-sql-driver/mysql"
)

/*
   You can run this using mysql docker image.

   1. Start a mysql container using:

     docker run -p 13306:3306 -e MYSQL_ROOT_PASSWORD=rootpwd -e MYSQL_DATABASE=testdb  -d mysql:9.1


   2. Install mysql CLI tool for your platform

   3. Run:

      mysql -h localhost -P 13306 -p -u root --protocol=tcp

     Enter the root password when asked: "rootpwd"

   4. Create a new user

     create user 'myuser'@'%' identified by 'mypassword';

   5. Grant privileges to the new user

     grant all on testdb.* to 'myuser'@'%';


   6. Run this go program

*/

func main() {
	// Use mysql driver name and driver specific connection string
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(localhost:13306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Check if database connection succeeded, with 5 second timeout
	ctx, cancel := context.WithTimeout(context.
		Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Success!")
}
