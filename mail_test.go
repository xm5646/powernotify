package powernotify

import (
	"fmt"
	"testing"
)

func TestMailSender_Send(t *testing.T) {
	mailReceiver := NewMailReceiver("lixiaoming", "lixmsucc@163.com")
	receiver := "lixiaoming@qq.com"
	cc := "xxx@qq.com"
	mail := NewMailMessage("测试邮件通知", "<html><body><h1>hello</h1></body></html>", HtmlType, nil)
	sender := &MailSender{}
	sender = sender.LoadConfig(NewMailConfig("smtp.qq.com", 465, "530107801@qq.com", "授权码或密码", true))
	sender = sender.AddReceiver(receiver).AddMailReceiver(mailReceiver).AddMail(mail).AddCc(cc)
	send, err := sender.Send()
	if err != nil {
		fmt.Printf("has error:%s", err.Error())
		return
	} else {
		fmt.Println(send)
	}
}
