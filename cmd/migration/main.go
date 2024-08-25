package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"homecomp/internal/configs"
	"homecomp/internal/database"
)

const (
	commandUp   string = "up"
	commandDown string = "down"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("action is required (up/down)")
		return
	}

	conf, err := configs.NewConfig()
	if err != nil {
		panic(fmt.Sprintf("can't load configuration: %s", err.Error()))
	}
	db, err := database.NewConnection(conf.Database)
	if err != nil {
		panic(fmt.Sprintf("can't establish database connection: %s", err.Error()))
	}
	driver, err := mysql.WithInstance(db.GetDB(), &mysql.Config{})
	if err != nil {
		panic(fmt.Sprintf("cannot get a driver instance for migraions: %s", err.Error()))
	}
	m, err := migrate.NewWithDatabaseInstance("file://./migrations/", "homecomp", driver)
	if err != nil {
		panic(fmt.Sprintf("cannot get migration instance: %s", err.Error()))
	}

	switch os.Args[1] {
	case commandUp:
		if len(os.Args) > 2 {
			steps, err := strconv.Atoi(os.Args[2])
			if err != nil {
				panic("invalid steps argument, should be a number")
			}
			m.Steps(steps)
		}
		m.Up()
		fmt.Println("[success] - migration up complete")
	case commandDown:
		if len(os.Args) > 2 {
			steps, err := strconv.Atoi(os.Args[2])
			if err != nil {
				panic("invalid steps argument, should be a number")
			}
			m.Steps(steps * -1)
		}
		m.Down()
		fmt.Println("[success] - migration down complete")
	default:
		fmt.Println("unknown actions")
	}
}
