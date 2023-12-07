package convey

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSplit(t *testing.T) {
	c.Convey("基础用例", t, func() {
		var (
			s      = "a:b:c"
			sep    = ":"
			except = []string{"a", "b", "c"}
		)

		got := Split(s, sep)
		c.So(got, c.ShouldResemble, except)
	})

	c.Convey("不包含分隔符用例", t, func() {
		var (
			s      = "a:b:c"
			sep    = "|"
			except = []string{"a:b:c"}
		)

		got := Split(s, sep)
		c.So(got, c.ShouldResemble, except)
	})
}
