package powernotify

type MessageType string

const (
	TextType     MessageType = "text"
	HtmlType     MessageType = "html"
	LinkType     MessageType = "link"     // 钉钉消息机器人
	MarkdownType MessageType = "markdown" // 钉钉消息机器人
)

type Message struct {
	Title   string
	Message string
	Type    MessageType
}

type DingTalkMessage struct {
	Message
	PicUrl     string
	MessageUrl string
	AtAll      bool
	AtMobiles  []string
}
