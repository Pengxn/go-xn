package model

import (
	"xorm.io/xorm"

	"github.com/Pengxn/go-xn/src/util/db"
	"github.com/Pengxn/go-xn/src/util/log"
)

var orm *xorm.Engine = db.DBEngine()

func InitTables() {
	initTable(new(User))
	initTable(new(Article))
	initTable(new(Option))
}

func initTable(bean interface{}) {
	if exist, err := orm.IsTableExist(bean); err != nil {
		log.Panic("query table exist error: ", err)
	} else if exist {
		return
	}

	if err := orm.Sync(bean); err != nil {
		log.Panic("create table error: ", err)
	}
}
