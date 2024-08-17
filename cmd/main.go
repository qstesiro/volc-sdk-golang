package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	// "github.com/volcengine/volc-sdk-golang/example/rtc/v20231101"
	"github.com/volcengine/volc-sdk-golang/service/rtc/v20230801"
	"github.com/volcengine/volc-sdk-golang/service/rtc/v20231101"
)

func main() {
	GetRoomOnlineUsers()
	ListDetectionTasks()
	// server()
}

func server() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong\n")
	})
	r.POST("/rtc/room", Callback)
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

func GetRoomOnlineUsers() {
	instance := rtc_v20230801.NewInstance()
	param := &rtc_v20230801.GetRoomOnlineUsersQuery{
		AppID:  "66aeef78e091820121ab8847",
		RoomID: "123",
	}
	resp, statusCode, err := instance.GetRoomOnlineUsers(context.Background(), param)
	if err != nil {
		log.Fatalf("error = %v\n", err)
	}
	PrintResponse(resp, statusCode)
}

func ListDetectionTasks() {
	instance := rtc_v20231101.NewInstance()
	userID := "123"
	param := &rtc_v20231101.ListDetectionTaskQuery{
		AppID:  "66aeef78e091820121ab8847",
		RoomID: "123",
		UserID: &userID,
	}
	resp, statusCode, err := instance.ListDetectionTask(context.Background(), param)
	if err != nil {
		log.Fatalf("error = %v\n", err)
	}
	PrintResponse(resp, statusCode)
}

func PrintResponse(r interface{}, code int) {
	b, _ := json.Marshal(r)
	v := bytes.Buffer{}
	json.Indent(&v, b, "", "  ")
	log.Printf("code = %d\n%v\n", code, v.String())
}
