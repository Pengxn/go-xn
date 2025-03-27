package model

import (
	"log"

	"xorm.io/xorm"

	"github.com/Pengxn/go-xn/src/util/db"
)

var orm *xorm.Engine = db.DBEngine()

func InitTables() {
	initTable(new(User))
	initTable(new(Article))
	initTable(new(Option))
	initTable(new(WebAuthnCredential))
}

func initTable(bean any) {
	if exist, err := orm.IsTableExist(bean); err != nil {
		log.Panicln("query table exist error", err)
	} else if exist {
		return
	}

	if err := orm.Sync(bean); err != nil {
		log.Panicln("create table error", err)
	}
}
