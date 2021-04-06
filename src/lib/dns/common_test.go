package dns

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
