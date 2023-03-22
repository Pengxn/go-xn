package db

import (
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/Pengxn/go-xn/src/config"
)

func TestDBUrl(t *testing.T) {
	patch := ApplyFuncSeq(config.GetDBConfig, []OutputCell{
		{Values: Params{config.DBConfig{Type: "MySQL"}}},
		{Values: Params{config.DBConfig{Type: "PostgreSQL"}}},
		{Values: Params{config.DBConfig{Type: "SQLite3"}}},
		{Values: Params{config.DBConfig{}}},
	})
	defer patch.Reset()

	Convey("Test if DB Url is correct.", t, func() {
		Convey("Test DBUrl when database is MySQL", func() {
			want := ":@tcp(:)/?charset=utf8"
			dbType, dsn := getDBUrl()

			So(dbType, ShouldEqual, "mysql")
			So(dsn, ShouldEqual, want)
		})

		Convey("Test DBUrl when database is PostgreSQL", func() {
			want := "dbname= user= password= host= port="
			dbType, dsn := getDBUrl()

			So(dbType, ShouldEqual, "postgres")
			So(dsn, ShouldEqual, want)
		})

		Convey("Test DBUrl when database is SQLite", func() {
			want := "file:?cache=shared&mode=rwc"
			dbType, dsn := getDBUrl()

			So(dbType, ShouldEqual, "sqlite3")
			So(dsn, ShouldEqual, want)
		})

		Convey("Test DBUrl when database is not supported", func() {
			dbType, dsn := getDBUrl()

			So(dbType, ShouldEqual, "")
			So(dsn, ShouldEqual, "")
		})
	})
}
