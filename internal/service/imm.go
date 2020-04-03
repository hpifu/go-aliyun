package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func PercentEncode(str string) string {
	str = url.QueryEscape(str)
	str = strings.ReplaceAll(str, "+", "%20")
	str = strings.ReplaceAll(str, "*", "%2A")
	str = strings.ReplaceAll(str, "%7E", "~")
	return str
}

func Signature(methods string, params map[string]string, accessKeySecret string) string {
	var kvs []string
	for k, v := range params {
		kvs = append(kvs, fmt.Sprintf("%v=%v", PercentEncode(k), PercentEncode(v)))
	}
	sort.Strings(kvs)
	str := strings.Join(kvs, "&")
	toSign := methods + "&" + PercentEncode("/") + "&" + PercentEncode(str)
	h := hmac.New(sha1.New, []byte(accessKeySecret+"&"))
	h.Write([]byte(toSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func MakePopParams(methods string, params map[string]string, accessKeyID string, accessKeySecret string) map[string]string {
	if _, ok := params["AccessKeyId"]; !ok {
		params["AccessKeyId"] = accessKeyID
	}
	if _, ok := params["Format"]; !ok {
		params["Format"] = "JSON"
	}
	if _, ok := params["Version"]; !ok {
		params["Version"] = "2017-09-06"
	}
	if _, ok := params["Timestamp"]; !ok {
		params["Timestamp"] = time.Now().Format(time.RFC3339)
	}
	if _, ok := params["SignatureMethod"]; !ok {
		params["SignatureMethod"] = "HMAC-SHA1"
	}
	if _, ok := params["SignatureVersion"]; !ok {
		params["SignatureVersion"] = "1.0"
	}
	if _, ok := params["Format"]; !ok {
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, rand.Uint64())
		params["SignatureNonce"] = hex.EncodeToString(buf)
	}
	params["Signature"] = Signature(methods, params, accessKeySecret)
	return params
}

func imm(endpoint string, params map[string]string) ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, 500*time.Millisecond)
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		},
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return nil, err
	}

	if params != nil {
		values := &url.Values{}
		for k, v := range params {
			values.Add(k, v)
		}
		req.URL.RawQuery = values.Encode()
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

type IMMReq struct {
	Endpoint   string            `form:"message"`
	Params     map[string]string `form:"params"`
	Credential string            `form:"credential"`
}

func (s *Service) IMM(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	params := map[string]string{}
	for k, v := range c.Request.URL.Query() {
		params[k] = v[0]
	}

	//imm()

	return req, &EchoRes{
		Message: req.Message,
	}, http.StatusOK, nil
}
