package parameter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GETParametersReq struct {
	Category    string `uri:"category" json:"category"`
	SubCategory string `uri:"subCategory" json:"subCategory"`
}

func (s *Service) GETParameters(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	req := &GETParameterReq{}

	if err := c.BindUri(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}

	res, err := s.ps.List(req.Category, req.SubCategory)
	if err != nil {
		return req, nil, http.StatusNoContent, nil
	}

	return req, res, http.StatusOK, nil
}
