// dingtalk test
package powernotify

import (
	"fmt"
	"testing"
)

func TestDingTalkSender_Send(t *testing.T) {
	token := "xxxx"
	secret := "xxxx"
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
