package credential

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DELETECredentialReq struct {
	Filename string `uri:"filename" form:"filename"`
}

type DELETECredentialRes struct{}

func (s *Service) DELETECredential(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	req := &DELETECredentialReq{}

	if err := c.BindUri(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}

	err := s.cs.Del(req.Filename)
	if err != nil {
		return nil, nil, http.StatusNoContent, nil
	}

	return req, nil, http.StatusAccepted, nil
}
