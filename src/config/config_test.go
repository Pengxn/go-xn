package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDBUrl(t *testing.T) {
	Convey("Test DB url", t, func() {
		want := "root:password@tcp(127.0.0.1:3306)/fyj?charset=utf8"
		dbType, dsn := DBUrl()

		So(dbType, ShouldEqual, "mysql")
		So(dsn, ShouldEqual, want)
	})
}
