package endpoint

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DELETEEndpointReq struct {
	Category string `uri:"category" json:"category"`
	Endpoint string `form:"endpoint" json:"endpoint"`
}

func (s *Service) DELETEEndpoint(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	req := &DELETEEndpointReq{}

	if err := c.BindUri(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}
	if err := c.Bind(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}

	err := s.es.Del(req.Category, req.Endpoint)
	if err != nil {
		return req, nil, http.StatusNoContent, nil
	}

	return req, nil, http.StatusAccepted, nil
}
