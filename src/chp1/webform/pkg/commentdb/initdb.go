package commentdb

import (
	"context"
	"database/sql"
)

const createStmt = `create table if not exists comments (
email TEXT,
comment TEXT)`

func InitDB(conn *sql.DB) {
	_, err := conn.ExecContext(context.Background(), createStmt)
	if err != nil {
		panic(err)
	}
}
