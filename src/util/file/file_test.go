package path

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsExist(t *testing.T) {
	Convey("Check if path exists.", t, func() {
		Convey("Pass a file name that exists", func() {
			So(IsExist("file.go"), ShouldEqual, true)
		})
		Convey("Pass a directory name that exists", func() {
			So(IsExist("testdata"), ShouldEqual, true)
		})
		Convey("Pass a directory name that does not exist", func() {
			So(IsExist("test"), ShouldEqual, false)
		})
	})
}

func TestIsFile(t *testing.T) {
	Convey("Check if given path is a file.", t, func() {
		Convey("Pass a file name that exists", func() {
			So(IsFile("file.go"), ShouldEqual, true)
		})
		Convey("Pass a directory name that exists", func() {
			So(IsFile("testdata"), ShouldEqual, false)
		})
		Convey("Pass a file name that does not exist", func() {
			So(IsFile("test.txt"), ShouldEqual, false)
		})
	})
}

func TestIsDir(t *testing.T) {
	Convey("Check if given path is a directory.", t, func() {
		Convey("Pass a file name", func() {
			So(IsDir("file.go"), ShouldEqual, false)
		})
		Convey("Pass a directory name", func() {
			So(IsDir("testdata"), ShouldEqual, true)
		})
		Convey("Pass a invalid path", func() {
			So(IsDir("foo"), ShouldEqual, false)
		})
	})
}
