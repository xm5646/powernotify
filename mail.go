package powernotify

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
)

type MailSender struct {
	Receivers []MailReceiver
	Cc        []MailReceiver
	Mails     []MailMessage
	MailConfig
}

type MailReceiver struct {
	Name    string
	Address string
}

type MailMessage struct {
	Message
	AttachesPath []string
}

func NewMailReceiver(name, address string) MailReceiver {
	return MailReceiver{
		Name:    name,
		Address: address,
	}
}

func NewMailMessage(title, message string, messageType MessageType, attaches []string) MailMessage {
	return MailMessage{
		Message: Message{
			Title:   title,
			Message: message,
			Type:    messageType,
		},
		AttachesPath: attaches,
	}
}

func NewMailConfig(host string, port int, username, password string, tls bool) MailConfig {
	return MailConfig{
		Host:      host,
		Port:      port,
		Username:  username,
		Password:  password,
		TLSSecure: tls,
	}
}

func NewMailSender(mailConfig MailConfig, to []MailReceiver, cc []MailReceiver, mails []MailMessage) *MailSender {
	return &MailSender{
		Receivers:  to,
		Cc:         cc,
		Mails:      mails,
		MailConfig: mailConfig,
	}
}

// Sending multiple emails, return number of success send, and if has error.
func (m *MailSender) Send() (int, error) {
	if len(m.Mails) > 0 && m.Mails != nil {
		count := 0
		// create dialer
		dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)
		if !m.TLSSecure {
			dialer.TLSConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
		var errs []error
		for _, mail := range m.Mails {
			// build message
			message := gomail.NewMessage()
			message.SetHeader("From", m.Username)
			message.SetHeader("To", m.getReceiversList())
			if m.Cc != nil {
				for _, mailReceiver := range m.Cc {
					message.SetAddressHeader("Cc", mailReceiver.Address, mailReceiver.Name)
				}
			}
			message.SetHeader("Subject", mail.Title)
			switch mail.Type {
			case TextType:
				message.SetBody("text/plain", mail.Message.Message)
			case HtmlType:
				message.SetBody("text/html", mail.Message.Message)
			default:
				message.SetBody("text/plain", mail.Message.Message)
			}
			if mail.AttachesPath != nil {
				for _, attach := range mail.AttachesPath {
					message.Attach(attach)
				}
			}
			if err := dialer.DialAndSend(message); err != nil {
				errs = append(errs, err)
				continue
			}
			count += 1
		}
		if count < len(m.Mails) {
			return count, fmt.Errorf("发送邮件出现异常, errors: %+v", errs)
		} else {
			return count, nil
		}
	} else {
		return 0, nil
	}

}

func (m *MailSender) getReceiversList() string {
	result := ""
	if m.Receivers == nil || len(m.Receivers) < 1 {
		return ""
	}
	for _, receiver := range m.Receivers {
		if result == "" {
			result = receiver.Address
		} else {
			result = result + " " + receiver.Address
		}
	}
	return result
}

type MailConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	TLSSecure bool   `json:"tls_secure"`
}
