package powernotify

import (
	"fmt"
	"testing"
)

func TestMailSender_Send(t *testing.T) {
	receivers := make([]MailReceiver, 0, 1)
	re1 := NewMailReceiver("lixiaoming", "lixmsucc@163.com")
	receivers = append(receivers, re1)
	mails := make([]MailMessage, 0, 1)
	mail := NewMailMessage("测试邮件通知", "<html><body><h1>hello</h1></body></html>", HtmlType, nil)
	mails = append(mails, mail)
	mailConfig := NewMailConfig("smtp.qq.com", 465, "530107801@qq.com", "djymdeerfphobihh", true)
	sender := NewMailSender(mailConfig, receivers, nil, mails, nil)
	send, err := sender.Send()
	if err != nil {
		fmt.Printf("has error:%s", err.Error())
		return
	} else {
		fmt.Println(send)
	}

}
