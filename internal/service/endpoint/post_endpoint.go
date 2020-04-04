package endpoint

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type POSTEndpointReq struct {
	Category string `uri:"category" json:"category"`
	Endpoint string `form:"endpoint" json:"endpoint"`
}

type POSTEndpointRes struct {
	Rid string `json:"rid"`
}

func (s *Service) POSTEndpoint(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	req := &POSTEndpointReq{}

	if err := c.BindUri(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}
	if err := c.Bind(req); err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("bind failed. err: [%v]", err)
	}

	err := s.es.Put(req.Category, req.Endpoint)
	if err != nil {
		return req, nil, http.StatusBadRequest, fmt.Errorf("put parameters failed. err: [%v]", err)
	}

	return req, &POSTEndpointRes{
		Rid: rid,
	}, http.StatusOK, nil
}
