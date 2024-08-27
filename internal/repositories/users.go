package repositories

import (
	"fmt"

	"golang.org/x/net/context"

	"homecomp/internal/database"
)

const (
	UserTableName   string = "users"
	UserFieldID     string = "id"
	UserFieldEmail  string = "email"
	UserFieldPasswd string = "password"
)

type UserRepo interface {
	CreateUser(data UserRow) error
	GetUserByEmail(email string) UserRow
}

type UserRow struct {
	ID       uint32
	Email    string
	Password string
}

type userRepo struct {
	db  database.DBCon
	ctx context.Context
}

func NewUserRepo(ctx context.Context, db database.DBCon) UserRepo {
	return &userRepo{
		db:  db,
		ctx: ctx,
	}
}

func (ur *userRepo) CreateUser(data UserRow) error {
	query := fmt.Sprintf("insert into %s(%s, %s) values(?, ?)", UserTableName, UserFieldEmail, UserFieldPasswd)
	stmt, err := ur.db.Prepare(ur.ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Email, data.Password)
	return err
}

func (ur *userRepo) GetUserByEmail(email string) UserRow {
	var id uint32
	query := fmt.Sprintf("select %s from %s where %s = ?", UserFieldID, UserTableName, UserFieldEmail)
	ur.db.GetDB().QueryRow(query, email).Scan(&id)

	return UserRow{
		ID:    id,
		Email: email,
	}
}
