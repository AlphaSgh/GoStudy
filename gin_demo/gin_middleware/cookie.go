package gin_middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CookieUse() {
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
		}
		fmt.Println("the value of cookie is ", cookie)
	})
	r.Run(":8080")
}
