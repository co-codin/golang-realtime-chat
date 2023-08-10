package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertId int
	query := "INSERT INTO users(username, password, email) VALUES ($1, $2, $3) returning id"
	
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertId)
	
	if err != nil {
		return &User{}, err
	}

	return nil, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return nil, nil
}