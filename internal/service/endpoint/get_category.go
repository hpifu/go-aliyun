package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) GETCategory(rid string, c *gin.Context) (interface{}, interface{}, int, error) {
	return nil, []string{"imm", "ots", "oss"}, http.StatusOK, nil
}

