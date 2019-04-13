package util

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"go-xn/src/config"
)

// DBEngine initialize database connection engine
// and return *xorm.Engine
func DBEngine() *xorm.Engine {
	orm, err := xorm.NewEngine("mysql", config.DBUrl())

	PanicIf(err)

	orm.TZLocation = time.Local

	return orm
}
