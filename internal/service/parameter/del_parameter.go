package parameter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DELETECredentialReq struct {
	Category    string `uri:"category" json:"category"`
	SubCategory string `uri:"subCategory" json:"subCategory"`
	Filename    string `uri:"filename" json:"filename"`
}

func (s *Service) DELETECredential(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	req := &DELETECredentialReq{}

	if err := c.BindUri(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}

	err := s.ps.Del(req.Category, req.SubCategory, req.Filename)
	if err != nil {
		return nil, nil, http.StatusNoContent, nil
	}

	return req, nil, http.StatusAccepted, nil
}
