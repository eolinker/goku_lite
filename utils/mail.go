package utils

import (
	"crypto/tls"
	"fmt"

	"github.com/eolinker/goku-api-gateway/server/entity"

	"gopkg.in/gomail.v2"
)

//SendMails 发送邮件
func SendMails(smtpInfo entity.SMTPInfo, mailTo []string, subject string, body string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s<%s>", smtpInfo.Sender, smtpInfo.Account))
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtpInfo.Address, smtpInfo.Port, smtpInfo.Account, smtpInfo.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	var err error
	err = d.DialAndSend(m)
	return err
}

////SendToMail 发送邮件
//func SendToMail(user, password, host, to, subject, body, mailtype, smtpProtocol string) error {
//	hp := strings.Split(host, ":")
//	auth := smtp.PlainAuth("", user, password, hp[0])
//	sendTO := strings.Split(to, ",")
//	if len(sendTO) < 2 {
//		if sendTO[0] == "" {
//			return nil
//		}
//	}
//	log.Debug(user, password, auth)
//	var contentType string
//	if mailtype == "html" {
//		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
//	} else {
//		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
//	}
//	var err error
//	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
//	log.Debug("SendToMail", msg)
//	if smtpProtocol != "0" {
//		err = SendMailUsingTLS(
//			host,
//			auth,
//			user,
//			sendTO,
//			msg,
//		)
//		if err != nil {
//
//			err = SendMailUsingTLS(
//				host,
//				nil,
//				user,
//				sendTO,
//				msg,
//			)
//
//		}
//	} else {
//		err = smtp.SendMail(host, auth, user, sendTO, msg)
//		if err != nil {
//			err = smtp.SendMail(host, nil, user, sendTO, msg)
//		}
//	}
//	return err
//}
//
////Dial return a smtp client
//func Dial(addr string) (*smtp.Client, error) {
//	conn, err := tls.Dial("tcp", addr, nil)
//	if err != nil {
//		return nil, err
//	}
//	//分解主机端口字符串
//	host, _, _ := net.SplitHostPort(addr)
//	return smtp.NewClient(conn, host)
//}
//
////SendMailUsingTLS 参考net/smtp的func SendMail()
////使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
////len(to)>1时,to[1]开始提示是密送
//func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
//	to []string, msg []byte) (err error) {
//
//	//create smtp client
//	c, err := Dial(addr)
//	if err != nil {
//		log.Error("Create smpt client error:", err)
//		return err
//	}
//	defer c.Close()
//
//	if auth != nil {
//		if ok, _ := c.Extension("AUTH"); ok {
//			if err = c.Auth(auth); err != nil {
//				log.Error("Error during AUTH", err)
//				return err
//			}
//		}
//	}
//
//	if err = c.Mail(from); err != nil {
//		return err
//	}
//
//	for _, addr := range to {
//		if err = c.Rcpt(addr); err != nil {
//			return err
//		}
//	}
//
//	w, err := c.Data()
//	if err != nil {
//		return err
//	}
//
//	_, err = w.Write(msg)
//	if err != nil {
//		return err
//	}
//
//	err = w.Close()
//	if err != nil {
//		return err
//	}
//	return c.Quit()
//}
