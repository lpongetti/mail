package mail

import "net/smtp"

type ISmtpServer interface {
	SendMail(to []string, msg []byte) error
}

type smtpServer struct {
	host   string
	port   string
	sender string
	auth   smtp.Auth
}

func NewSMTPServer(host, port, username, password, sender string) ISmtpServer {
	return &smtpServer{
		host:   host,
		port:   port,
		sender: sender,
		auth:   smtp.PlainAuth("", username, password, host),
	}
}

func (ss *smtpServer) address() string {
	return ss.host + ":" + ss.port
}

func (ss *smtpServer) SendMail(to []string, msg []byte) error {
	return smtp.SendMail(ss.address(), ss.auth, ss.sender, to, msg)
}
