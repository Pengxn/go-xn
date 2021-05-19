package httplib

import (
	"net/http"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSetCookie(t *testing.T) {
	client := New().SetCookie(&http.Cookie{
		Name:     "test",
		Value:    "test",
		SameSite: http.SameSiteDefaultMode,
	})

	Convey("Test SetCookie.", t, func() {
		So(client.cookie, ShouldResemble, []*http.Cookie{
			{
				Name:     "test",
				Value:    "test",
				SameSite: http.SameSiteDefaultMode,
			},
		})
	})
}

func TestSetHeader(t *testing.T) {
	client := NewClient(NewHttpConfig()).SetHeader(map[string]string{
		"test": "test",
	})

	Convey("Test SetHeader.", t, func() {
		So(client.header, ShouldResemble, map[string]string{"test": "test"})
	})
}

func TestSetHeaderFace(t *testing.T) {
	client := NewTimeoutClient(50).SetHeaderFace(map[string]string{
		"test": "test",
	})

	Convey("Test SetHeader.", t, func() {
		So(client.config.Timeout, ShouldEqual, time.Second*50)
		So(client.header, ShouldResemble, map[string]string{"test": "test"})
		So(client.keepCase, ShouldEqual, true)
	})
}

func TestSetParams(t *testing.T) {
	client := New().SetParams(map[string][]string{
		"test": {"test"},
	})

	Convey("Test SetParams.", t, func() {
		So(client.params, ShouldResemble, map[string][]string{
			"test": {"test"},
		})
	})
}

func TestSetHost(t *testing.T) {
	client := New().SetHost("https://huiyifyj.cn")

	Convey("Test SetHost.", t, func() {
		So(client.host, ShouldResemble, "https://huiyifyj.cn")
	})
}

func TestSetContentType(t *testing.T) {
	client := New().SetContentType(JSON)

	Convey("Test SetContentType.", t, func() {
		So(client.contentType, ShouldResemble, JSON)
	})
}
