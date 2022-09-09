package common

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func Connect(url string, l *Logger) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		l.Log("Error connecting to database")
		return nil
	}

	return conn
}
