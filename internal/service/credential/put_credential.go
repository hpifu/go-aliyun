package credential

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hpifu/go-aliyun/internal/store"
)

type POSTCredentialReq struct {
	Filename string `uri:"filename" form:"filename"`
	store.Credential
}

type POSTCredentialRes struct {
	Filename string `json:"filename"`
	Rid      string `json:"rid"`
}

func (s *Service) POSTCredential(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	req := &POSTCredentialReq{}

	if err := c.Bind(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}

	err := s.cs.Put(req.Filename, req.AccessKeyID, req.AccessKeySecret)
	if err != nil {
		return nil, nil, http.StatusNoContent, nil
	}

	return req, &POSTCredentialRes{
		Filename: req.Filename,
		Rid:      rid,
	}, http.StatusOK, nil
}
