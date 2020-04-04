package credential

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GETCredentialReq struct {
	Filename string `uri:"filename" form:"filename"`
}

func (s *Service) GETCredential(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	req := &GETCredentialReq{}

	if err := c.BindUri(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}

	credential, err := s.cs.Get(req.Filename)
	if err != nil {
		return nil, nil, http.StatusNoContent, nil
	}

	return req, credential, http.StatusOK, nil
}
