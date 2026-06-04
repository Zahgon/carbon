package gorm

import (
	"gorm.io/gorm"
)

const (
	driverMySQL  = "mysql"
	driverPgSQL  = "postgres"
	driverSQLite = "sqlite"
)

var (
	dsn string
	dia gorm.Dialector
	db  *gorm.DB
	err error
)

func connect(driver string) *gorm.DB { _ = "STUB: not implemented"; return nil }
