// Code generated by protoc-gen-volcengine-sdk
// source: VodWorkflowService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func Test_StartWorkflow(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodStartWorkflowRequest{
		Vid:               "your Vid",
		TemplateId:        "your TemplateId",
		Input:             nil,
		Priority:          0,
		CallbackArgs:      "your CallbackArgs",
		EnableLowPriority: false,
	}

	resp, status, err := instance.StartWorkflow(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_RetrieveTranscodeResult(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodRetrieveTranscodeResultRequest{
		Vid:        "your Vid",
		ResultType: "your ResultType",
	}

	resp, status, err := instance.RetrieveTranscodeResult(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetWorkflowExecution(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetWorkflowExecutionStatusRequest{
		RunId:           "your RunId",
		NeedTasksDetail: "your NeedTasksDetail",
	}

	resp, status, err := instance.GetWorkflowExecution(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
