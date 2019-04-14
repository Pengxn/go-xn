package util

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"go-xn/src/config"
)

// DBEngine initialize database connection engine
// and return *xorm.Engine
func DBEngine() *xorm.Engine {
	orm, err := xorm.NewEngine("mysql", config.DBUrl())

	if err != nil {
		panic(err)
	}

	if err = orm.Ping(); err != nil {
		log.Fatalln("Your DB can't work it normally.", err.Error())
	}

	orm.TZLocation = time.Local
	orm.ShowSQL(true)

	return orm
}
