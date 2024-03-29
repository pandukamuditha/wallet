// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package repositories

import (
	"database/sql"
	"time"
)

type Auth struct {
	UserID   int64
	Username string
	Password string
	Claims   sql.NullString
}

type Transaction struct {
	ID          int64
	FromWallet  int64
	ToWallet    int64
	Amount      int32
	CreatedDate time.Time
	UpdatedDate time.Time
}

type User struct {
	ID    int64
	Fname string
	Lname string
}

type Wallet struct {
	ID      int64
	UserID  int64
	Balance int32
}
