package store

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEndpointStore(t *testing.T) {
	Convey("test endpoint store", t, func() {
		es, err := NewEndpointStore("data/endpoint")
		So(err, ShouldBeNil)

		Convey("test put", func() {
			So(es.Put("imm", "https://imm.cn-shanghai.aliyuncs.com"), ShouldBeNil)
			So(es.Put("imm", "https://imm.cn-beijing.aliyuncs.com"), ShouldBeNil)
			So(es.Put("imm", "https://imm.cn-hangzhou.aliyuncs.com"), ShouldBeNil)
		})

		Convey("test list", func() {
			c, err := es.List("imm")
			So(err, ShouldBeNil)
			So(c, ShouldResemble, []string{
				"https://imm.cn-beijing.aliyuncs.com",
				"https://imm.cn-hangzhou.aliyuncs.com",
				"https://imm.cn-shanghai.aliyuncs.com",
			})
		})

		Convey("test del", func() {
			So(es.Del("imm", "https://imm.cn-beijing.aliyuncs.com"), ShouldBeNil)
		})
	})
}
