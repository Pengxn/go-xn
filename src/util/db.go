package util

import (
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"xorm.io/xorm"

	"github.com/Pengxn/go-xn/src/config"
)

// DBEngine initialize database connection engine
// and return *xorm.Engine
func DBEngine() *xorm.Engine {
	orm, err := xorm.NewEngine("mysql", config.DBUrl())
	if err != nil {
		log.Fatalln("Can't connect your DB.", err.Error())
	}

	return orm
}
