package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		r1 := gin.Default()
		//http://localhost:8992/sayHello/
		r1.GET("/sayHello/", func(c *gin.Context) {
			c.JSON(200, "hello1")
		})
		r1.Run(":8992")
	}()

	r := gin.Default()
	//http://localhost:8991/sayHello/
	r.GET("/sayHello/", func(c *gin.Context) {
		c.JSON(200, "hello")
	})
	r.Run(":8991")
}
