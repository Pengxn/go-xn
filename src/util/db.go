package util

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"go-xn/src/config"
)

var orm *xorm.Engine

// SetEngine initialize database connection
// and return *xorm.Engine
func SetEngine() *xorm.Engine {
	orm, err := xorm.NewEngine("mysql", config.DBUrl())

	PanicIf(err)

	orm.TZLocation = time.Local

	return orm
}
