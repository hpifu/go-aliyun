package store

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCredentialStore(t *testing.T) {
	Convey("test credential store", t, func() {
		cs, err := NewCredentialStore("data/credential")
		So(err, ShouldBeNil)

		Convey("test put", func() {
			So(cs.Put("test1", "b866b4af589873fb68137a2cec56969f", "dfab538cddfd3fab36a87d6ac1ea4a45"), ShouldBeNil)
		})

		Convey("test get", func() {
			c, err := cs.Get("test1")
			So(err, ShouldBeNil)
			So(c.AccessKeyID, ShouldEqual, "b866b4af589873fb68137a2cec56969f")
			So(c.AccessKeySecret, ShouldEqual, "dfab538cddfd3fab36a87d6ac1ea4a45")
		})

		Convey("test list", func() {
			fns, err := cs.List()
			So(err, ShouldBeNil)
			So(fns, ShouldResemble, []string{"test1"})
		})

		Convey("test del", func() {
			So(cs.Del("test1"), ShouldBeNil)
		})
	})
}
