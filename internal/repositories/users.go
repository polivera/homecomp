package repositories

import (
	"golang.org/x/net/context"

	"homecomp/internal/database"
)

type UserRepo interface {
	CreateUser(data UserRow) error
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
	query := "insert into users(email, password) values(?, ?)"
	stmt, err := ur.db.Prepare(ur.ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Email, data.Password)
	return err
}
