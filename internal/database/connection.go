package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"homecomp/internal/configs"
)

type DBCon interface {
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

func (dcn *dbCon) Query(query string) (*sql.Rows, error) {
	return dcn.db.Query(query)
}

func (dcn *dbCon) Close() {
	err := dcn.db.Close()
	if err != nil {
		panic("cannot close connection to database")
	}
}
