# PowerNotify
## Introduction

PowerNotify是一个简单的发送通知的包，目前实现了邮件通知和钉钉消息通知，后续将继续完善短信通知

邮件通知通过封装gopkg.in/gomail.v2提供发送邮件功能

钉钉通知需要提供钉钉告警机器人webhook地址中的access_token, 以及加签安全设置中的秘钥


## Features

- 邮件通知
  - 一个client发送多封邮件
  - 支持text 和html格式的邮件内容
  - 支持TLS开关
  - 多人收件,多人抄送
  - 多个附件

- 钉钉消息通知
  - 加签安全设置
  - 支持text、markdown, link 格式的通知内容
  - 支持@所有人
  - 支持@部分群成员

## Download

``` go get gopkg.in/gomail.v2
go get github.com/xm5646/powernotify
```

## Examples

`邮件通知`:

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
        mailConfig := NewMailConfig("smtp.qq.com", 465, "530107801@qq.com", "授权码", true)
        sender := NewMailSender(mailConfig, receivers, nil, mails, nil)
        send, err := sender.Send()
        if err != nil {
            fmt.Printf("has error:%s", err.Error())
            return
        } else {
            fmt.Println(send)
        }
    }

`钉钉通知`:

    func TestDingTalkSender_Send(t *testing.T) {
        token := "xxxx"  // webhook中的access_token
        secret := "xxxx"  // 加签 安全设置中的秘钥
        message := NewDingTalkNormalMessage("测试", "#### 杭州天气 \\n> 9度，西北风1级，空气良89，相对温度73%\\n> ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\\n> ###### 10点20分发布 [天气](https://www.dingtalk.com) \\n",
            MarkdownType, false, nil)
        sender := NewDingTalkSender([]DingTalkReceiver{NewDingTalkReceiver(token, secret)}, message)
        send, err := sender.Send()
        if err != nil {
            fmt.Printf("error:%s", err.Error())
        } else {
            fmt.Println(send)
        }
    }


## License

[MIT](LICENSE)

## Contact

如有建议或者问题，可以联系邮箱lixmsucc@163.com, 希望能帮助到你，也欢迎对本项目提交代码