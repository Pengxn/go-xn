package dns

import (
	"net/http"
	"reflect"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/Pengxn/go-xn/src/util/httplib"
)

func TestSetAction(t *testing.T) {
	dns := NewDNS()

	Convey("Test SetAction.", t, func() {
		So(dns.SetAction("test").action, ShouldEqual, "test")
	})
}

func TestDo(t *testing.T) {
	post := func(_ *httplib.Client, url string) (*http.Response, error) {
		return nil, nil
	}
	patch := ApplyMethod(reflect.TypeOf(&httplib.Client{}), "POST", post)
	defer patch.Reset()

	dns := DNSCommon{
		action: "action",
		region: "region", // optional param
	}
	response, err := dns.do(map[string]string{
		"Other": "other", // other params
	})
	Convey("Test do.", t, func() {
		So(response, ShouldEqual, nil)
		So(err, ShouldEqual, nil)
	})
}

func TestSortParams(t *testing.T) {
	params := map[string]string{
		"Timestamp": "1234",
		"SecretId":  "secret",
		"Action":    "test",
		"Signature": "fyj",
	}
	result := "Action=test&SecretId=secret&Signature=fyj&Timestamp=1234"

	Convey("Test sortParams.", t, func() {
		So(sortParams(params), ShouldEqual, result)
	})
}

func TestSignHmacSha256(t *testing.T) {
	result := "fxY/WK6EqoVmc0NMofRgT3jBf1Df46wI54dCJQ75ILo="
	Convey("Test signHmacSha256.", t, func() {
		So(signHmacSha256("fyj", "psh"), ShouldEqual, result)
	})
}
