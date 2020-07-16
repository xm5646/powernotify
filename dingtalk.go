package powernotify

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type DingTalkReceiver struct {
	AccessToken string
	Secret      string
}

type DingTalkSender struct {
	Receivers []DingTalkReceiver
	DingTalkMessage
}

func NewDingTalkReceiver(token, secret string) DingTalkReceiver {
	return DingTalkReceiver{
		AccessToken: token,
		Secret:      secret,
	}
}

// build text and markdown message
func NewDingTalkNormalMessage(title, message string, messageType MessageType, atAll bool, atMobiles []string) DingTalkMessage {
	return DingTalkMessage{
		Message: Message{
			Title:   title,
			Message: message,
			Type:    messageType,
		},
		AtAll:     atAll,
		AtMobiles: atMobiles,
	}
}

// build link message
func NewDingTalkLinkMessage(title, message string, messageType MessageType, picUrl, link string, atAll bool, atMobiles []string) DingTalkMessage {
	return DingTalkMessage{
		Message: Message{
			Title:   title,
			Message: message,
			Type:    messageType,
		},
		PicUrl:     picUrl,
		MessageUrl: link,
		AtAll:      atAll,
		AtMobiles:  atMobiles,
	}
}

func NewDingTalkSender(receivers []DingTalkReceiver, message DingTalkMessage) DingTalkSender {
	return DingTalkSender{
		Receivers:       receivers,
		DingTalkMessage: message,
	}
}

func (d *DingTalkSender) Send() (int, error) {
	if len(d.Receivers) > 0 {
		count := 0
		for _, receiver := range d.Receivers {
			timestamp, sign := buildSign(receiver)
			var message string
			switch d.Message.Type {
			case TextType:
				message = fmt.Sprintf(`{"msgtype": "text","text": {"content": "%s"},"at":{"isAtAll": %t, "atMobiles": %s}}`, d.Message.Message, d.AtAll, fmt.Sprintf("[%s]", strings.Join(d.AtMobiles, ",")))
			case MarkdownType:
				message = fmt.Sprintf(`{"msgtype": "markdown","markdown": {"title": "%s", "text": "%s"},"at":{"isAtAll": %t, "atMobiles": %s}}`, d.Message.Title, d.Message.Message, d.AtAll, fmt.Sprintf("[%s]", strings.Join(d.AtMobiles, ",")))
			case LinkType:
				message = fmt.Sprintf(`{"msgtype": "link","link": {"title": "%s", "text": "%s", "picUrl": "%s", "messageUrl": "%s"},"at":{"isAtAll": %t, "atMobiles": %s}}`, d.Message.Title, d.Message.Message, d.DingTalkMessage.PicUrl, d.DingTalkMessage.MessageUrl, d.AtAll, fmt.Sprintf("[%s]", strings.Join(d.AtMobiles, ",")))
			default:
				message = fmt.Sprintf(`{"msgtype": "text","text": {"content": "%s"},"at":{"isAtAll": %t, "atMobiles": %s}}`, d.Message.Message, d.AtAll, fmt.Sprintf("[%s]", strings.Join(d.AtMobiles, ",")))
			}
			client := &http.Client{}
			value := url.Values{}
			value.Set("access_token", receiver.AccessToken)
			if receiver.Secret != "" {
				value.Set("timestamp", fmt.Sprintf("%d", timestamp))
				value.Set("sign", sign)
			}
			request, _ := http.NewRequest("POST", "https://oapi.dingtalk.com/robot/send", bytes.NewBuffer([]byte(message)))
			fmt.Println(value.Encode())
			request.URL.RawQuery = value.Encode()
			request.Header.Set("Content-type", "application/json;charset=utf-8")
			response, err := client.Do(request)
			if err != nil {
				return 0, fmt.Errorf("访问钉钉URL(%+v) 出错了: %s", receiver, err)
			}
			if response.StatusCode != 200 {
				body, _ := ioutil.ReadAll(response.Body)
				return 0, fmt.Errorf("访问钉钉URL(%+v) 出错了: %s", receiver, string(body))
			}
			all, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return 0, fmt.Errorf("访问钉钉URL(%+v) 出错了: %s", receiver, err)
			}
			var dtErr DingTalkErr
			err = json.Unmarshal(all, &dtErr)
			if err != nil {
				return 0, fmt.Errorf("无法解析钉钉接口返回数据: %s", string(all))
			}
			if dtErr.Errcode != 0 {
				return 0, fmt.Errorf("发送钉钉通知失败, errcode:%d, errmsg:%s", dtErr.Errcode, dtErr.Errmsg)
			}
			fmt.Println(string(all))
			count += 1
		}
		return count, nil
	} else {
		return 0, nil
	}
}

func buildSign(auth DingTalkReceiver) (timestamp int64, sign string) {
	timestamp = time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, auth.Secret)
	h := hmac.New(sha256.New, []byte(auth.Secret))
	h.Write([]byte(stringToSign))
	sha := h.Sum(nil)
	return timestamp, base64.StdEncoding.EncodeToString(sha)
}

type DingTalkErr struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
