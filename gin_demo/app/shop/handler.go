package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func goodsHandler(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("shop goodsHandler"))
}
func checkoutHandler(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("shop checkoutHandler"))
}
