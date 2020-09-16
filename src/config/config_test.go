package config

import (
	"testing"

	. "github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDBUrl(t *testing.T) {
	ApplyFuncSeq(getDBConfiguration, []OutputCell{
		{Values: Params{&DBConfiguration{Type: "MySQL"}}},
		{Values: Params{&DBConfiguration{Type: "PostgreSQL"}}},
		{Values: Params{&DBConfiguration{Type: "SQLite3"}}},
		{Values: Params{&DBConfiguration{}}},
	})

	Convey("Test DBUrl when database is MySQL", t, func() {
		want := ":@tcp(:)/?charset=utf8"
		dbType, dsn := DBUrl()

		So(dbType, ShouldEqual, "mysql")
		So(dsn, ShouldEqual, want)
	})

	Convey("Test DBUrl when database is PostgreSQL", t, func() {
		want := "dbname= user= password= host= port="
		dbType, dsn := DBUrl()

		So(dbType, ShouldEqual, "postgres")
		So(dsn, ShouldEqual, want)
	})

	Convey("Test DBUrl when database is SQLite", t, func() {
		want := "file:?cache=shared&mode=rwc"
		dbType, dsn := DBUrl()

		So(dbType, ShouldEqual, "sqlite3")
		So(dsn, ShouldEqual, want)
	})

	Convey("Test DBUrl when database is not supported", t, func() {
		dbType, dsn := DBUrl()

		So(dbType, ShouldEqual, "")
		So(dsn, ShouldEqual, "")
	})
}
