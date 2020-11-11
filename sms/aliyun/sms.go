package aliyun

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/xm5646/powernotify/errno"
)

func (c *AliAPIClient) SendSMS(tempCode string, signName string, phone string, args map[string]string) *errno.Err  {
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = signName
	request.TemplateCode = tempCode
	temStr, _ := json.Marshal(args)
	request.TemplateParam = string(temStr)
	response, err := c.Instance.SendSms(request)
	if err != nil {
		fmt.Println("无法发送短信，调用短信网关失败。")
		return errno.New(&errno.Errno{
			Code:    500,
			Message: "调用API网关失败,暂时无法发送短信",
		}, err)
	} else if response.Code != "OK" {
		fmt.Printf("发送短信失败, %s\n", response.Message)
		return errno.New(&errno.Errno{
			Code:    500,
			Message: response.Message,
		}, nil)
	}

	return nil
}
