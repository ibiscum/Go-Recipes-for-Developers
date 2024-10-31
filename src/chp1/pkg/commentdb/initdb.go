package commentdb

import (
	"context"
	"database/sql"
)

func InitDB(conn *sql.DB) {
	_, err := conn.ExecContext(context.Background(), `create table if not exists comments (
  email TEXT,
  comment TEXT)`)
	if err != nil {
		panic(err)
	}
}
