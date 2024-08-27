package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"homecomp/internal/configs"
)

type DBCon interface {
	GetDB() *sql.DB
	Query(ctx context.Context, query string) (*sql.Rows, error)
	Prepare(ctx context.Context, query string) (*sql.Stmt, error)
	Close()
}

type dbCon struct {
	db *sql.DB
}

func NewConnection(conf configs.DBConfig) (DBCon, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Pass, conf.Host, conf.Port, conf.Name),
	)
	if err != nil {
		return &dbCon{}, err
	}

	rows, err := db.Query("SELECT 1")
	if err != nil {
		return &dbCon{}, err
	}
	defer rows.Close()

	return &dbCon{
		db: db,
	}, err
}

func (dcn *dbCon) GetDB() *sql.DB {
	return dcn.db
}

func (dcn *dbCon) Query(ctx context.Context, query string) (*sql.Rows, error) {
	return dcn.db.QueryContext(ctx, query)
}

func (dcn *dbCon) Prepare(ctx context.Context, query string) (*sql.Stmt, error) {
	return dcn.db.PrepareContext(ctx, query)
}

func (dcn *dbCon) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	return dcn.db.QueryRowContext(ctx, query, args)
}

func (dcn *dbCon) Close() {
	err := dcn.db.Close()
	if err != nil {
		panic("cannot close connection to database")
	}
}
