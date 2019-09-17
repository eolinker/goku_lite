package utils

import (
	"crypto/tls"
	log "github.com/eolinker/goku/goku-log"
	"net"
	"net/smtp"
	"strings"
)

var period map[string]string = map[string]string{
	"0": "1",
	"1": "5",
	"2": "15",
	"3": "30",
	"4": "60",
}

func SendMail(sender, subject, senderPassword, smtpAddress, smtpPort, smtpProtocol, receiverMail, content string) {
	host := net.JoinHostPort(smtpAddress, smtpPort)
	err := SendToMail(sender, senderPassword, host, receiverMail, subject, content, "html", smtpProtocol)
	if err != nil {
		log.Warn("SendMail:",err)
	}
}

func SendToMail(user, password, host, to, subject, body, mailtype, smtpProtocol string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	send_to := strings.Split(to, ",")
	if len(send_to) < 2 {
		if send_to[0] == "" {
			return nil
		}
	}
	log.Debug(user, password, auth)
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	var err error
	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	log.Debug("SendToMail",msg)
	if smtpProtocol != "0" {
		err = SendMailUsingTLS(
			host,
			auth,
			user,
			send_to,
			msg,
		)
		if err != nil {

			err = SendMailUsingTLS(
				host,
				nil,
				user,
				send_to,
				msg,
			)

		}
	} else {
		err = smtp.SendMail(host, auth, user, send_to, msg)
		if err != nil {
			err = smtp.SendMail(host, nil, user, send_to, msg)
		}
	}
	return err
}

//return a smtp client
func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
//len(to)>1时,to[1]开始提示是密送
func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {

	//create smtp client
	c, err := Dial(addr)
	if err != nil {
		log.Error("Create smpt client error:", err)
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Error("Error during AUTH", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
