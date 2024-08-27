package repositories

import (
	"context"
	"fmt"

	"homecomp/internal/database"
)

const (
	UserTableName   string = "users"
	UserFieldID     string = "id"
	UserFieldEmail  string = "email"
	UserFieldPasswd string = "password"
)

type UserRepo interface {
	CreateUser(ctx context.Context, data UserRow) error
	GetUserByEmail(ctx context.Context, email string) UserRow
}

type UserRow struct {
	ID       uint32
	Email    string
	Password string
}

type userRepo struct {
	db database.DBCon
}

func NewUserRepo(db database.DBCon) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) CreateUser(ctx context.Context, data UserRow) error {
	query := fmt.Sprintf("insert into %s(%s, %s) values(?, ?)", UserTableName, UserFieldEmail, UserFieldPasswd)
	stmt, err := ur.db.Prepare(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Email, data.Password)
	return err
}

func (ur *userRepo) GetUserByEmail(ctx context.Context, email string) UserRow {
	var id uint32
	query := fmt.Sprintf("select %s from %s where %s = ?", UserFieldID, UserTableName, UserFieldEmail)
	ur.db.QueryRow(ctx, query, email).Scan(&id)

	return UserRow{
		ID:    id,
		Email: email,
	}
}
