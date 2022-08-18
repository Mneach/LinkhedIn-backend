package service

import (
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func useDB() *gorm.DB { return db }

func SendEmail(ToMail string) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "MneachDev@gmail.com")
	msg.SetHeader("To", ToMail)
	msg.SetHeader("Subject", "Confirm your account on LinkhedIn")
	msg.SetBody("text/html", "Activate your account via this link : <br> http://localhost:8080/activation/LKdjflksewnlDFLKSJ")

	n := gomail.NewDialer("smtp.gmail.com", 587, "MneachDev@gmail.com", "qpdwvwkkyxovxtue")

	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}

}
