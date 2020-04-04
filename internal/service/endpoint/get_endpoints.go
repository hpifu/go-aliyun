package endpoint

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GETEndpointsReq struct {
	Category string `uri:"category" json:"category"`
}

func (s *Service) GETEndpoints(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	req := &GETEndpointsReq{}

	if err := c.BindUri(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}

	res, err := s.es.List(req.Category)
	if err != nil {
		return req, nil, http.StatusNoContent, nil
	}

	return req, res, http.StatusOK, nil
}
