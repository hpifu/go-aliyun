package parameter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GETParameterReq struct {
	Filename    string `uri:"filename"`
	Category    string `uri:"category"`
	SubCategory string `uri:"subCategory"`
}

func (s *Service) GETParameter(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	req := &GETParameterReq{}

	if err := c.BindUri(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}

	params, err := s.ps.Get(req.Category, req.SubCategory, req.Filename)
	if err != nil {
		return nil, nil, http.StatusNoContent, nil
	}

	return req, params, http.StatusOK, nil
}
