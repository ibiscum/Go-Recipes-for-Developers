package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
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
	ID        uint64
	Name      string
	LastLogin time.Time
	AvatarURL string
}

type UpdateUserRequest struct {
	Name      *string
	LastLogin *time.Time
	AvatarURL *string
}

type UserSearchRequest struct {
	Ids            []uint64
	Name           *string
	LoggedInBefore *time.Time
	LoggedInAfter  *time.Time
	AvatarURL      *string
}

func main() {
	// Open the sqlite database using the given local file ./database.db
	db, err := sql.Open("pgx", "postgres://postgres:mysecretpassword@localhost:15432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// You don't need to ping an embedded database
	db.Exec(`CREATE TABLE user_info (
user_id int not null,
user_name varchar(32) not null,
last_login timestamp null,
avatar_url varchar(128) null
)`)

	{
		// Add some users to the database
		users := make([]User, 0)
		for i := 0; i < 100; i++ {
			t := time.Now().Add(time.Duration(-rand.Intn(1000)) * time.Minute)
			users = append(users, User{
				ID:        uint64(i),
				Name:      fmt.Sprintf("User-%d", i),
				AvatarURL: fmt.Sprintf("http://@example.com/%d_avatar", i),
				LastLogin: t,
			})
		}

		if err := AddUsers(db, users); err != nil {
			panic(err)
		}
		fmt.Println("Inserted 100 users")
	}

	after := time.Now().Add(-time.Hour)
	rows, err := db.Query(`SELECT user_id, user_name, last_login, avatar_url FROM user_info WHERE last_login > $1`, after)
	if err != nil {
		panic(err)
	}
	// Close the rows object when done
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		// Retrieve data from this row
		var user User
		// avatar column is nullable, so we pass a *string instead of string
		var avatarURL *string
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.LastLogin,
			&avatarURL); err != nil {
			panic(err)
		}
		// avatar URL can be nil in the db
		if avatarURL != nil {
			user.AvatarURL = *avatarURL
		}
		users = append(users, user)
	}
	// Check if there was an error during iteration
	if err := rows.Err(); err != nil {
		panic(err)
	}
	fmt.Println(users)

	// Update one user
	now := time.Now()
	urlString := "https://example.org/avatar.jpg"
	update := UpdateUserRequest{
		LastLogin: &now,
		AvatarURL: &urlString,
	}

	err = UpdateUser(context.Background(), db, 1, &update)
	if err != nil {
		panic(err)
	}

	users, err = SearchUsers(context.Background(), db, &UserSearchRequest{
		Ids: []uint64{1, 2, 3, 4, 5},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Users with id=1,2,3,4,5", users)

	str := "User-1"
	users, err = SearchUsers(context.Background(), db, &UserSearchRequest{
		Name: &str,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Users with name=User-1", users)
}

func AddUsers(db *sql.DB, users []User) error {
	stmt, err := db.Prepare(`INSERT INTO user_info (user_id,user_name,last_login,avatar_url) VALUES ($1,$2,$3,$4)`)
	if err != nil {
		return err
	}
	// Close the prepared statement when done
	defer stmt.Close()
	for _, user := range users {
		// Run the prepared statement with different arguments
		_, err := stmt.Exec(user.ID, user.Name, user.LastLogin, user.AvatarURL)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateUser(ctx context.Context, db *sql.DB, userId uint64, req *UpdateUserRequest) error {
	query := strings.Builder{}
	args := make([]interface{}, 0)
	// Start building the query. Be mindful of spaces to separate uery clauses
	query.WriteString("UPDATE user_info SET ")
	if req.Name != nil {
		args = append(args, *req.Name)
		fmt.Fprintf(&query, "user_name=$%d", len(args))
	}
	if req.LastLogin != nil {
		if len(args) > 0 {
			query.WriteString(",")
		}
		args = append(args, *req.LastLogin)
		fmt.Fprintf(&query, "last_login=$%d", len(args))
	}
	if req.AvatarURL != nil {
		if len(args) > 0 {
			query.WriteString(",")
		}
		args = append(args, *req.AvatarURL)
		fmt.Fprintf(&query, "avatar_url=$%d", len(args))
	}
	args = append(args, userId)
	fmt.Fprintf(&query, " WHERE user_id=$%d", len(args))
	_, err := db.ExecContext(ctx, query.String(), args...)
	return err
}

func SearchUsers(ctx context.Context, db *sql.DB, req *UserSearchRequest) ([]User, error) {
	query := strings.Builder{}
	where := strings.Builder{}
	args := make([]interface{}, 0)
	// Start building the query. Be mindful of spaces to separate	query clauses
	query.WriteString("SELECT user_id, user_name, last_login,avatar_url FROM user_info ")

	if len(req.Ids) > 0 {
		// Add this to the WHERE clause with an AND
		if where.Len() > 0 {
			where.WriteString(" AND ")
		}
		// Build an IN clause.
		// We have to add one argument for each id
		where.WriteString("user_id IN (")
		for i, id := range req.Ids {
			if i > 0 {
				where.WriteString(",")
			}
			args = append(args, id)
			fmt.Fprintf(&where, "$%d", len(args))
		}
		where.WriteString(")")
	}
	if req.Name != nil {
		if where.Len() > 0 {
			where.WriteString(" AND ")
		}
		args = append(args, *req.Name)
		fmt.Fprintf(&where, "user_name=$%d", len(args))
	}
	if req.LoggedInBefore != nil {
		if where.Len() > 0 {
			where.WriteString(" AND ")
		}
		args = append(args, *req.LoggedInBefore)
		fmt.Fprintf(&where, "last_login<$%d", len(args))
	}
	if req.LoggedInAfter != nil {
		if where.Len() > 0 {
			where.WriteString(" AND ")
		}
		args = append(args, *req.LoggedInAfter)
		fmt.Fprintf(&where, "last_login>$%d", len(args))
	}
	if req.AvatarURL != nil {
		if where.Len() > 0 {
			where.WriteString(" AND ")
		}
		args = append(args, *req.AvatarURL)
		fmt.Fprintf(&where, "avatar_url=$%d", len(args))
	}
	if where.Len() > 0 {
		query.WriteString(" WHERE ")
		query.WriteString(where.String())
	}
	rows, err := db.QueryContext(ctx, query.String(), args...)
	if err != nil {
		return nil, err
	}
	users := make([]User, 0)
	for rows.Next() {
		// Retrieve data from this row
		var user User
		// avatar column is nullable, so we pass a *string instead of string
		var avatarURL *string
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.LastLogin,
			&avatarURL); err != nil {
			return nil, err
		}
		// avatar URL can be nil in the db
		if avatarURL != nil {
			user.AvatarURL = *avatarURL
		}
		users = append(users, user)
	}
	// Check if there was an error during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
