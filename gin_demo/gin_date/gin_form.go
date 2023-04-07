package gin_date

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func GinForm() {
	r := gin.Default()
	r.POST("/loginForm", func(c *gin.Context) {
		var form LoginForm
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadGateway, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8080")
}
