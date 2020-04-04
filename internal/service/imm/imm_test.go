package imm

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHamc(t *testing.T) {
	Convey("test hmac", t, func() {
		mac := hmac.New(sha256.New, []byte("123456"))
		mac.Write([]byte("hello world"))
		So(hex.EncodeToString(mac.Sum(nil)), ShouldEqual, "83b3eb2788457b46a2f17aaa048f795af0d9dabb8e5924dd2fc0ea682d929fe5")
	})
}

func TestSignature(t *testing.T) {
	Convey("test signature", t, func() {
		So(Signature("POST", map[string]string{
			"Project":          "test-project",
			"RegionId":         "cn-shanghai",
			"AccessKeyId":      "testid",
			"Format":           "JSON",
			"SignatureMethod":  "HMAC-SHA1",
			"SignatureVersion": "1.0",
			"SignatureNonce":   "d1ac7371108dc53541c9d0f29e5396c7",
			"Timestamp":        "2019-02-22T09:30:54Z",
			"Action":           "GetProject",
			"Version":          "2017-09-06",
		}, "testsecret"), ShouldEqual, "NPzJnV5HAdj4jkShTWKa9WwOZxU=")
	})
}
