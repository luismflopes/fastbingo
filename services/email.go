package services

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
)

type EmailService struct {
	auth smtp.Auth
}

func NewEmailService() *EmailService {
	return &EmailService{
		auth: smtp.PlainAuth("", "luismiguelflo@gmail.com", "MariaRegina2011!", "smtp.gmail.com"),
	}
}

// SendNewAccountEmail
func (me *EmailService) SendEmailNewAccount(to, name, confirmationToken string) {
	templateData := struct {
		Name string
		URL  string
	}{
		Name: name,
		URL:  fmt.Sprintf("http://localhost:8080/user/email/confirm/%s", confirmationToken),
	}

	body, err := me.parseTemplate("templates/html/new.account.html", templateData)
	if err != nil {
		log.Println(err)
	}

	_, err = me.sendEmail([]string{to}, "Welcome!", body)
	if err != nil {
		log.Println(err)
	}
}

// SendNewAccountEmail
func (me *EmailService) SendEmailResetPassword(to, name, resetPasswordToken string) {
	templateData := struct {
		Name string
		URL  string
	}{
		Name: name,
		URL:  fmt.Sprintf("http://localhost:8080/user/password/forgot?token=%s", resetPasswordToken),
	}

	body, err := me.parseTemplate("templates/html/reset.password.html", templateData)
	if err != nil {
		log.Println(err)
	}

	_, err = me.sendEmail([]string{to}, "Reset Password Requested", body)
	if err != nil {
		log.Println(err)
	}
}

func (me *EmailService) sendEmail(to []string, subject, body string) (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject = fmt.Sprintf("Subject: %s!\n", subject)
	msg := []byte(subject + mime + "\n" + body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, me.auth, "luismiguelflo@gmail.com", to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (me *EmailService) parseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
