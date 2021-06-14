package home

import (
	"errors"
	"os"
	"os/user"
	"runtime"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
)

var ps = string(os.PathSeparator)

func TestHomeDir(t *testing.T) {
	Convey("Test Home Directory Path", t, func() {
		home, _ := user.Current()

		So(HomeDir(), ShouldEqual, home.HomeDir)
	})
}

func TestHome(t *testing.T) {
	Convey("Test home and return non-empty string.", t, func() {
		var homePath string
		if runtime.GOOS == "windows" {
			homePath = "C:\\Users"
		} else {
			homePath = "~/fyj"
		}

		p := ApplyFunc(os.Getenv, func(_ string) string {
			return homePath
		})
		defer p.Reset()

		So(home(), ShouldEqual, homePath)
	})

	Convey("Test home and return empty string.", t, func() {
		p := ApplyFunc(os.Getenv, func(_ string) string {
			return ""
		})
		defer p.Reset()
		p.ApplyFunc(user.Current, func() (*user.User, error) {
			return &user.User{}, errors.New("user.Current error")
		})

		So(home(), ShouldEqual, "")
	})
}

func TestConfigDir(t *testing.T) {
	Convey("Test home and return non-empty string.", t, func() {
		var homePath string
		if runtime.GOOS == "windows" {
			homePath = "C:\\Users"
		} else {
			homePath = "~/fyj"
		}

		p := ApplyFunc(os.Getenv, func(_ string) string {
			return homePath
		})
		defer p.Reset()

		So(ConfigDir("test"), ShouldEqual, homePath+ps+".config"+ps+"test")
	})
}

func TestCacheDir(t *testing.T) {
	Convey("Test home and return non-empty string.", t, func() {
		var homePath string
		if runtime.GOOS == "windows" {
			homePath = "C:\\Users"
		} else {
			homePath = "~/fyj"
		}

		p := ApplyFunc(os.Getenv, func(_ string) string {
			return homePath
		})
		defer p.Reset()

		So(CacheDir("test"), ShouldEqual, homePath+ps+".cache"+ps+"test")
	})
}

func TestDataDir(t *testing.T) {
	Convey("Test home and return non-empty string.", t, func() {
		var homePath string
		if runtime.GOOS == "windows" {
			homePath = "C:\\Users"
		} else {
			homePath = "~/fyj"
		}

		p := ApplyFunc(os.Getenv, func(_ string) string {
			return homePath
		})
		defer p.Reset()

		So(DataDir("test"), ShouldEqual, homePath+ps+".local"+ps+"share"+ps+"test")
	})
}
