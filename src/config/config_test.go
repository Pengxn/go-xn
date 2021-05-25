package config

import (
	"testing"

	"github.com/go-ini/ini"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/Pengxn/go-xn/src/util/log"
)

var configTest *ini.File

func init() {
	configFile, err := ini.Load("example.ini")
	if err != nil {
		log.Errorf("ini.Load error: %v", err)
	}
	configTest = configFile
}

func TestGetDBConfig(t *testing.T) {
	patch := ApplyGlobalVar(&config, configTest)
	defer patch.Reset()

	Convey("Test GetDBConfig", t, func() {
		So(GetDBConfig(), ShouldResemble, &DBConfig{
			Type:     "MySQL",
			User:     "root",
			Password: "password",
			Port:     "3306",
			Name:     "fyj",
			Url:      "127.0.0.1",
		})
	})
}

func TestGetDNSConfig(t *testing.T) {
	patch := ApplyGlobalVar(&config, configTest)
	defer patch.Reset()

	Convey("Test GetDNSConfig", t, func() {
		So(GetDNSConfig(), ShouldResemble, &DNSConfig{
			SecretID:  "AKIDAF*** ***Z0Gno6C",
			SecretKey: "IdxgOu*** ***Fj2CFYJ",
		})
	})
}

func TestGetSentryConfig(t *testing.T) {
	patch := ApplyGlobalVar(&config, configTest)
	defer patch.Reset()

	Convey("Test GetSentryConfig", t, func() {
		So(GetSentryConfig(), ShouldResemble, &SentryConfig{
			DSN:   "https:/0eaj7***6gv4s@sentry.io/1234567",
			Debug: true,
		})
	})
}
