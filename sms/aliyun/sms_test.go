package aliyun

import (
	"fmt"
	"testing"
)

func TestAliAPIClient_SendSMS(t *testing.T) {
	ALIClient.Init("", "", "")
	args := make(map[string]string)
	args["code"] = "123456"
	err := ALIClient.SendSMS("", "", "", args)
	if err != nil {
		fmt.Println("短信发送失败")
		fmt.Printf("%d,%s", err.Code, err.Message)
	}
}