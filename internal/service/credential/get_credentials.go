package credential

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) GETCredentials(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	req := &GETCredentialReq{}

	cs, err := s.cs.List()
	if err != nil {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("list failed. err: [%v]", err)
	}

	return req, cs, http.StatusOK, nil
}
