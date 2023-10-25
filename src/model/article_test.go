package model

import (
	"github.com/go-ini/ini"
	"xorm.io/xorm"

	"github.com/Pengxn/go-xn/src/util/log"
)

func init() {
	CleanTable(testORM, &Article{})
	db := testORM.NewSession()
	defer db.Close()
	if num, err := db.Insert(
		&Article{ID: 1, Title: "One", Content: "1: Content One", Status: 1},
		&Article{ID: 2, Title: "Two", Content: "2: Content Two", Status: 0},
		&Article{ID: 3, Title: "Three", Content: "3: Content Three", Status: 1},
		&Article{ID: 4, Title: "Four", Content: "4: Content Four", Status: 2},
		&Article{ID: 5, Title: "Five", Content: "5: Content Five", Status: 1},
		&Article{ID: 6, Title: "Six", Content: "6: Content Six", Status: 1},
		&Article{ID: 7, Title: "Seven", Content: "7: Content Seven", Status: 0},
		&Article{ID: 8, Title: "Eight", Content: "8: Content Eight", Status: 1},
		&Article{ID: 9, Title: "Nine", Content: "9: Content Nine", Status: 1},
	); err != nil {
		log.Infof("Init data to article table error: %+v, num: %d", err, num)
	}
}

var testORM = func() *xorm.Engine {
	config, err := ini.LooseLoad("fyj.ini")
	if err != nil {
		log.Errorln("Fail to read fyj.ini file.", err)
	}

	dbType := config.Section("test").Key("dbType").String()
	uri := config.Section("test").Key("dbURI").String()
	if uri == "" || dbType == "" {
		dbType = "sqlite3"
		uri = "./test.db?cache=shared&mode=rwc"
	}

	testDB, err := xorm.NewEngine(dbType, uri)
	if err != nil {
		log.Fatalln("Fail to connect test database.")
	}

	return testDB
}()

// CleanTable that is used to drop tables for test.
func CleanTable(orm *xorm.Engine, modelTable interface{}) {
	if has, _ := orm.IsTableExist(modelTable); has {
		if err := orm.DropTables(modelTable); err != nil {
			log.Fatalf("Drop table error: %+v, model: %+v", err, modelTable)
		}
	}
	if err := orm.CreateTables(modelTable); err != nil {
		log.Fatalf("Create table error: %+v, model: %+v", err, modelTable)
	}
}
