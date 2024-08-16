package main

import (
	"bytes"
	"encoding/json"
	"log"
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
	rtc := r.Group("/rtc")
	rtc.POST("/room", Callback)
	r.Run(":8080")
}

func Callback(ctx *gin.Context) {
	var m map[string]interface{}
	ctx.BindJSON(&m)
	Print(m)
}

func Print(m interface{}) {
	b, _ := json.Marshal(m)
	v := bytes.Buffer{}
	json.Indent(&v, b, "", "  ")
	log.Printf("success %v", v.String())
}
