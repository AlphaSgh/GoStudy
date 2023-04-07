package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func postHandler(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("blog postHandler"))
}
func commentHandler(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("blog commentHandler"))
}
