package markdown

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestToHTML(t *testing.T) {
	md := `# Hello, World!
> This is test.
`
	html := `<h1>Hello, World!</h1>
<blockquote>
<p>This is test.</p>
</blockquote>
`

	Convey("Test ToHTML.", t, func() {
		h, err := ToHTML([]byte(md))
		So(h, ShouldResemble, []byte(html))
		So(err, ShouldBeNil)
	})
}
