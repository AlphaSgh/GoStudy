package gin_middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("middleware execing now")
		c.Set("request", "middleware")
		//c.Next()
		status := c.Writer.Status()
		fmt.Println("middleware execed now", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
func AllTest() {
	r := gin.Default()
	r.Use(MiddleWare())
	{
		r.GET("/ce", func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			c.JSON(200, gin.H{"request": req})
		})
	}
	r.Run(":8080")
}
