package xorm

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"xorm.io/xorm"
)

const (
	driverMySQL  = "mysql"
	driverPgSQL  = "postgres"
	driverSQLite = "sqlite"
)

var (
	dsn string
	db  *xorm.Engine
	err error
)

func connect(driver string) *xorm.Engine { _ = "STUB: not implemented"; return nil }
