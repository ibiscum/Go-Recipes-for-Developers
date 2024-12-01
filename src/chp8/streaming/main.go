package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Result struct {
	Err   error
	Email string
}

func buildResult(rows *sql.Rows, result *Result) error {
	return rows.Scan(&result.Email)
}

func StreamResults(
	ctx context.Context,
	db *sql.DB,
	query string,
	args ...any,
) (<-chan Result, error) {
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	output := make(chan Result)
	go func() {
		defer rows.Close()
		defer close(output)
		var result Result
		for rows.Next() {
			// Check context cancellation
			if result.Err = ctx.Err(); result.Err != nil {
				// Context canceled. return
				output <- result
				return
			}
			// Set result fields
			result.Err = buildResult(rows, &result)
			output <- result
		}
		// If there was an error, return it
		if result.Err = rows.Err(); result.Err != nil {
			output <- result
		}
	}()
	return output, nil
}

const createStmt = `create table if not exists users (
email TEXT,
name TEXT)`

func initDb(conn *sql.DB) {
	_, err := conn.ExecContext(context.Background(), createStmt)
	if err != nil {
		panic(err)
	}

	// If there are fewer than 1000 rows, add some
	var n int
	conn.QueryRowContext(context.Background(), `select count(*) from users`).Scan(&n)
	for i := 0; i < 1000-n; i++ {
		name := fmt.Sprintf("User-%d", i)
		email := fmt.Sprintf("User%d@example.com", i)
		conn.ExecContext(context.Background(), `insert into users (name,email) values (?,?)`, name, email)
	}
}

func ProcessResult(result Result) error {
	fmt.Println(result)
	return nil
}

func main() {
	db, err := sql.Open("sqlite", "chp8.db")
	if err != nil {
		panic(err)
	}
	initDb(db)

	// Setup a cancelable context
	cancelableCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Call the streaming API
	results, err := StreamResults(cancelableCtx, db, "SELECT EMAIL FROM USERS")
	if err != nil {
		return
	}
	// Collect and process results
	for result := range results {
		if result.Err != nil {
			// Handle error in the result
			continue
		}
		// Process the result
		if err := ProcessResult(result); err != nil {
			// Processing error. Cancel streaming results
			cancel()
			// Expect to receive at least one more message from the channel,
			// because the streaming gorutine sends the error
			for range results {
			}
		}
	}
}
