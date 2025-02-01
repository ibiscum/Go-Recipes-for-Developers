package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/exp/rand"
)

/**
   You can run this example with a postgres container:

  docker run -p 15432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres:17


  The PGX driver uses $n for statement arguments instead of ?, so this
  example is adjusted to use $n convention.
*/

type User struct {
	Id        string
	Name      string
	Email     string
	LastLogin *time.Time
}

func main() {
	// Open the sqlite database using the given local file ./database.db
	db, err := sql.Open("pgx", "postgres://postgres:mysecretpassword@localhost:15432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// You don't need to ping an embedded database
	db.Exec(`CREATE TABLE users (
user_id varchar(32),
user_name varchar(128),
email varchar(128),
last_login timestamp)`)

	// Add some users to the database
	users := make([]User, 0)
	for i := 0; i < 100; i++ {
		t := time.Now().Add(time.Duration(-rand.Intn(1000)) * time.Minute)
		users = append(users, User{
			Id:        fmt.Sprint(i),
			Name:      fmt.Sprintf("User-%d", i),
			Email:     fmt.Sprintf("user%d@example.com", i),
			LastLogin: &t,
		})
	}

	// Add users in a single transaction
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 1. Start transaction
	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		panic(err)
	}

	if err := AddUsers(tx, users); err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()

	fmt.Println("Inserted 100 users in a single transaction")

	fmt.Println("First 10 users fetched by id")
	for i := 0; i < 10; i++ {
		user, err := GetUserByID(db, fmt.Sprint(i))
		if err != nil {
			panic(err)
		}
		fmt.Printf("User %+v\n", user)
	}

	t := time.Now().Add(-time.Hour)
	fmt.Printf("Users logged in within the last 10 hours (%v)\n", t)
	names, err := GetUserNamesLoggedInAfter(db, t)
	if err != nil {
		panic(err)
	}
	fmt.Println(names)
}

func AddUsers(tx *sql.Tx, users []User) error {
	stmt, err := tx.Prepare(`INSERT INTO users (user_id,user_name,email,last_login) VALUES ($1,$2,$3,$4)`)
	if err != nil {
		return err
	}
	// Close the prepared statement when done
	defer stmt.Close()
	for _, user := range users {
		// Run the prepared statement with different arguments
		_, err := stmt.Exec(user.Id, user.Name, user.Email, user.LastLogin)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetUserByID(db *sql.DB, id string) (*User, error) {
	var user User
	err := db.QueryRow(`SELECT user_id, user_name, last_login FROM
users WHERE user_id=$1`, id).
		Scan(&user.Id, &user.Name, &user.LastLogin)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserNamesLoggedInAfter(db *sql.DB, after time.Time) ([]string, error) {
	rows, err := db.Query(`SELECT users.user_name FROM users WHERE last_login > $1`, after)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	// Check if iteration produced any errors
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return names, nil
}
