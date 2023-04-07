package gin_middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	since := time.Since(start)
	fmt.Println("time:", since)
}
func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}
func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}
func TestMiddleWare() {
	r := gin.Default()
	r.Use(myTime)
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	r.Run(":8080")
}
