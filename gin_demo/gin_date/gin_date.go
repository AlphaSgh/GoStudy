package gin_date

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func SomeDataStruct() {
	r := gin.Default()
	//1.json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJson", "status": 200})
	})
	//2.struct
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 124
		c.JSON(200, msg)
	})
	//3.xml
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "abc"})
	})
	//4.yaml
	r.GET("/someYMAL", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "zhangsan"})
	})
	//5.protobuf
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		lable := "lable"
		data := &protoexample.Test{
			Label: &lable,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})
	r.Run(":8080")
}
