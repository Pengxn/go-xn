package config

import (
	"testing"

	. "github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDBUrl(t *testing.T) {
	Convey("Test DBUrl when database is MySQL", t, func() {
		ApplyFunc(getDBConfiguration, func() *DBConfiguration {
			return &DBConfiguration{Type: "MySQL"}
		})

		want := ":@tcp(:)/?charset=utf8"
		dbType, dsn := DBUrl()

		So(dbType, ShouldEqual, "mysql")
		So(dsn, ShouldEqual, want)
	})

	Convey("Test DBUrl when database is PostgreSQL", t, func() {
		ApplyFunc(getDBConfiguration, func() *DBConfiguration {
			return &DBConfiguration{Type: "PostgreSQL"}
		})

		want := "dbname= user= password= host= port="
		dbType, dsn := DBUrl()

		So(dbType, ShouldEqual, "postgres")
		So(dsn, ShouldEqual, want)
	})

	Convey("Test DBUrl when database is SQLite", t, func() {
		ApplyFunc(getDBConfiguration, func() *DBConfiguration {
			return &DBConfiguration{Type: "SQLite3"}
		})

		want := "file:?cache=shared&mode=rwc"
		dbType, dsn := DBUrl()

		So(dbType, ShouldEqual, "sqlite3")
		So(dsn, ShouldEqual, want)
	})

	Convey("Test DBUrl when database is not supported", t, func() {
		ApplyFunc(getDBConfiguration, func() *DBConfiguration {
			return &DBConfiguration{}
		})

		dbType, dsn := DBUrl()

		So(dbType, ShouldEqual, "")
		So(dsn, ShouldEqual, "")
	})
}
