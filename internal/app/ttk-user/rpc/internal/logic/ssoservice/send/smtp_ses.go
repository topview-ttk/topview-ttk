package send

import (
	"fmt"
	"net/smtp"
)

const (
	smtpServer = "your_smtp_server"
	smtpPort   = "587"
	username   = "your_smtp_user_name"
	password   = "your_smtp_password"
	fromEmail  = "your_smtp_user_name"
	sub        = "TTK验证码"
)

func SMTPEmail(toEmail, verifyCode string) error {
	auth := smtp.PlainAuth("", username, password, smtpServer)
	body := fmt.Sprintf("您的验证码为\"%s\",有效期为5分钟", verifyCode)
	msg := []byte("To: " + toEmail + "\r\n" +
		"From: " + fromEmail + "\r\n" +
		"Subject: " + sub + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n" +
		"\r\n" +
		body)

	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, fromEmail, []string{toEmail}, msg)
	if err != nil {
		return err
	}
	return nil
}
