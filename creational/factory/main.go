package main

import (
	"errors"
	"fmt"
)

type Notification interface {
	Send(msg string)
}

type EmailNotification struct {
}

func (EmailNotification) Send(msg string) {
	fmt.Println("Send email..." + msg)
}

type SmsNotification struct {
}

func (SmsNotification) Send(msg string) {
	fmt.Println("Sending sms ..." + msg)
}

func createNotification(notifType string) (Notification, error) {
	switch notifType {
	case "email":
		return EmailNotification{}, nil
	case "sms":
		return SmsNotification{}, nil
	default:
		return nil, errors.New("Provide a correct type!")

	}
}

func main() {
	emailnotif, _ := createNotification("email")
	emailnotif.Send("email")
	smsnotif, _ := createNotification("sms")
	smsnotif.Send("sms")
}
