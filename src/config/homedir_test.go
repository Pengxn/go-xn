package config

import (
	"os/user"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHomeDir(t *testing.T) {
	Convey("Test Home Directory Path", t, func() {
		home, _ := user.Current()

		So(HomeDir(), ShouldEqual, home.HomeDir)
	})
}
