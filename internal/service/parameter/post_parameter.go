package parameter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type POSTParameterReq struct {
	Category    string            `uri:"category" json:"category"`
	SubCategory string            `uri:"subCategory" json:"subCategory"`
	Filename    string            `form:"filename" json:"filename"`
	Params      map[string]string `form:"params" json:"params"`
}

type POSTParameterRes struct {
	Filename string `json:"filename"`
	Rid      string `json:"rid"`
}

func (s *Service) POSTParameter(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	req := &POSTParameterReq{}

	if err := c.BindUri(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}
	if err := c.Bind(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}

	err := s.ps.Put(req.Category, req.SubCategory, req.Filename, req.Params)
	if err != nil {
		return req, nil, http.StatusBadRequest, fmt.Errorf("put parameters failed. err: [%v]", err)
	}

	return req, &POSTParameterRes{
		Filename: req.Filename,
		Rid:      rid,
	}, http.StatusOK, nil
}
