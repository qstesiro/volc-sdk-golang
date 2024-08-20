package rtc_v20231101

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/volcengine/volc-sdk-golang/service/rtc/v20231101"
)

func StartDetection() {
	instance := rtc_v20231101.NewInstance()

	callback := "http://demo.qstesiro.top/rtc"
	callbackType := int32(1)
	param := &rtc_v20231101.StartDetectionBody{
		AppID:        "66aeef78e091820121ab8847",
		RoomID:       "123",
		Callback:     &callback,
		CallbackType: &callbackType,
	}

	resp, statusCode, err := instance.StartDetection(context.Background(), param)

	if err != nil {
		if resp != nil && resp.ResponseMetadata.Error != nil {
			errStr, _ := json.Marshal(resp.ResponseMetadata.Error)
			log.Printf("statusCode: %d, error: %v\n", statusCode, string(errStr))
			// 网关返回的错误
			if resp.ResponseMetadata.Error.CodeN != nil && *resp.ResponseMetadata.Error.CodeN != 0 {
				switch *resp.ResponseMetadata.Error.CodeN {
				// InvalidAccessKey
				case 100009:
					log.Println("请求的AK不合法")
				// SignatureDoesNotMatch
				case 100010:
					log.Println("签名结果不正确")
				}
			} else {
				// 服务端返回的错误
				switch resp.ResponseMetadata.Error.Code {
				case "InvalidParameter":
					log.Println("请求的参数错误, 请根据具体Error中的Message提示调整参数")
				}
			}
		} else {
			log.Fatalf("statusCode: %d, error: %v\n", statusCode, err)
		}
	} else {
		b, _ := json.Marshal(resp)
		v := bytes.Buffer{}
		json.Indent(&v, b, "", "  ")
		log.Printf("success %v", v.String())
	}
}
