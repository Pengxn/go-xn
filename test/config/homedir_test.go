package config

import (
	"os/user"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"go-xn/src/config"
)

func TestHomeDir(t *testing.T) {
	Convey("Test Home Directory Path", t, func() {
		home, _ := user.Current()

		So(config.HomeDir(), ShouldEqual, home.HomeDir)
	})
}
