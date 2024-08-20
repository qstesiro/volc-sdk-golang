package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/volcengine/volc-sdk-golang/example/rtc/v20230801"
	"github.com/volcengine/volc-sdk-golang/example/rtc/v20231101"
)

// 不同类型的单元测试
// https://stackoverflow.com/questions/19998250/proper-package-naming-for-testing-with-the-go-language
func main() {
	rtc_v20230801.GetRoomOnlineUsers()
	rtc_v20231101.StartDetection()
	<-time.After(time.Second)
	rtc_v20231101.ListDetectionTasks()
	<-time.After(time.Hour)
	rtc_v20231101.StopDetection()
	// server()
}

func server() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong\n")
	})
	r.POST("/rtc", Callback)
	r.Run(":8080")
}

func Callback(ctx *gin.Context) {
	var m map[string]interface{}
	ctx.BindJSON(&m)
	PrintCallback(m)
}

func PrintCallback(m interface{}) {
	b, _ := json.Marshal(m)
	v := bytes.Buffer{}
	json.Indent(&v, b, "", "  ")
	log.Printf("success %v", v.String())
}
