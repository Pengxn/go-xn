package home

import (
	"errors"
	"os"
	"os/user"
	"testing"

	. "github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHomeDir(t *testing.T) {
	Convey("Test Home Directory Path", t, func() {
		home, _ := user.Current()

		So(HomeDir(), ShouldEqual, home.HomeDir)
	})
}

func TestDirWindows(t *testing.T) {
	Convey("Test dirWindows and error is nil.", t, func() {
		p := ApplyFunc(os.Getenv, func(_ string) string {
			return "C:\\Users"
		})
		defer p.Reset()

		home, err := dirWindows()
		So(home, ShouldEqual, "C:\\Users")
		So(err, ShouldEqual, nil)
	})

	Convey("Test dirWindows and error isn't nil.", t, func() {
		p := ApplyFunc(os.Getenv, func(_ string) string {
			return ""
		})
		defer p.Reset()

		home, err := dirWindows()
		So(home, ShouldEqual, "")
		So(err, ShouldResemble, errors.New("Can't find 'USERPROFILE' environment variable"))
	})
}

func TestDirUnix(t *testing.T) {
	Convey("Test dirUnix and error is nil.", t, func() {
		p := ApplyFunc(os.Getenv, func(_ string) string {
			return "~"
		})
		defer p.Reset()

		home, err := dirUnix()
		So(home, ShouldEqual, "~")
		So(err, ShouldEqual, nil)
	})

	Convey("Test dirUnix and error isn't nil.", t, func() {
		p := ApplyFunc(os.Getenv, func(_ string) string {
			return ""
		})
		defer p.Reset()

		home, err := dirUnix()
		So(home, ShouldEqual, "")
		So(err, ShouldResemble, errors.New("Can't find 'HOME' environment variable"))
	})
}
