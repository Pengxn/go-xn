package httplib

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewHttpConfig(t *testing.T) {
	Convey("Test NewHttpConfig", t, func() {
		So(NewHttpConfig(), ShouldResemble, &Config{
			Timeout:            10 * time.Second,
			DisableKeepAlives:  true,
			InsecureSkipVerify: true,
			Encode:             true,
			IsEncodeForGo:      false,
			Charset:            "UTF-8",
			Certificates:       nil,
		})
	})
}
