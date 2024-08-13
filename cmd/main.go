package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/volcengine/volc-sdk-golang/example/rtc/v20230801"
)

func main() {
	// rtc_v20230801_test.GetRoomOnlineUsers()
	server()
}

func server() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong\n")
	})
	r.Run(":8080")
}
