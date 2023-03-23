package db

import (
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/Pengxn/go-xn/src/config"
)

func TestDBUrl(t *testing.T) {
	Convey("Test if DB Url is correct.", t, func() {
		Convey("Test DBUrl when database is MySQL", func() {
			patch := ApplyGlobalVar(&config.Config.DB, config.DBConfig{Type: "MySQL"})
			defer patch.Reset()

			want := ":@tcp(:)/?charset=utf8"
			dbType, dsn := getDBUrl()

			So(dbType, ShouldEqual, "mysql")
			So(dsn, ShouldEqual, want)
		})

		Convey("Test DBUrl when database is PostgreSQL", func() {
			patch := ApplyGlobalVar(&config.Config.DB, config.DBConfig{Type: "PostgreSQL"})
			defer patch.Reset()

			want := "dbname= user= password= host= port="
			dbType, dsn := getDBUrl()

			So(dbType, ShouldEqual, "postgres")
			So(dsn, ShouldEqual, want)
		})

		Convey("Test DBUrl when database is SQLite", func() {
			patch := ApplyGlobalVar(&config.Config.DB, config.DBConfig{Type: "SQLite3"})
			defer patch.Reset()

			want := "file:?cache=shared&mode=rwc"
			dbType, dsn := getDBUrl()

			So(dbType, ShouldEqual, "sqlite3")
			So(dsn, ShouldEqual, want)
		})

		Convey("Test DBUrl when database is not supported", func() {
			patch := ApplyGlobalVar(&config.Config.DB, config.DBConfig{})
			defer patch.Reset()

			dbType, dsn := getDBUrl()

			So(dbType, ShouldEqual, "")
			So(dsn, ShouldEqual, "")
		})
	})
}
