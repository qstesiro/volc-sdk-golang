package rtc_v20230801_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/volcengine/volc-sdk-golang/service/rtc/v20230801"
)

func GetRoomOnlineUsers() {
	instance := rtc_v20230801.NewInstance()

	param := &rtc_v20230801.GetRoomOnlineUsersQuery{
		AppID:  "66aeef78e091820121ab8847",
		RoomID: "123",
	}

	resp, statusCode, err := instance.GetRoomOnlineUsers(context.Background(), param)

	if err != nil {
		if resp != nil && resp.ResponseMetadata.Error != nil {
			errStr, _ := json.Marshal(resp.ResponseMetadata.Error)
			log.Printf("statusCode: %d, error: %v", statusCode, string(errStr))
			// 网关返回的错误
			if resp.ResponseMetadata.Error.CodeN != nil && *resp.ResponseMetadata.Error.CodeN != 0 {
				switch *resp.ResponseMetadata.Error.CodeN {
				// InvalidAccessKey
				case 100009:
					log.Printf("请求的AK不合法")
				// SignatureDoesNotMatch
				case 100010:
					log.Printf("签名结果不正确")
				}
			} else {
				// 服务端返回的错误
				switch resp.ResponseMetadata.Error.Code {
				case "InvalidParameter":
					log.Printf("请求的参数错误, 请根据具体Error中的Message提示调整参数")
				}
			}
		} else {
			log.Fatalf("statusCode: %d, error: %v", statusCode, err)
		}
	} else {
		b, _ := json.Marshal(resp)
		v := bytes.Buffer{}
		json.Indent(&v, b, "", "  ")
		log.Printf("success %v", v.String())
	}
}
