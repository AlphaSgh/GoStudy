package gin_date

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func AsyncSync() {
	r := gin.Default()
	r.GET("/long_async", func(c *gin.Context) {
		copyContext := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("async " + copyContext.Request.URL.Path)
		}()
	})
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("sync " + c.Request.URL.Path)
	})
	r.Run(":8080")
}
