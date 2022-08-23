package service

import (
	"bytes"
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/MneachDev/LinkhedIn-backend/graph/model"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func SendEmailActivation(ToMail string, linkActivation string) {

	link := linkActivation

	log.Print(linkActivation)

	msg := gomail.NewMessage()
	msg.SetHeader("From", "MneachDev@gmail.com")
	msg.SetHeader("To", ToMail)
	msg.SetHeader("Subject", "Confirm your account on LinkhedIn")
	msg.SetBody("text/html", `Activate your account via this link : <br> http://127.0.0.1:5173/activationAccount/`+link)

	n := gomail.NewDialer("smtp.gmail.com", 587, "MneachDev@gmail.com", "qpdwvwkkyxovxtue")

	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}

func SendEmailResetPassword(ToMail string, linkActivation string) {

	link := linkActivation

	log.Print(linkActivation)

	msg := gomail.NewMessage()
	msg.SetHeader("From", "MneachDev@gmail.com")
	msg.SetHeader("To", ToMail)
	msg.SetHeader("Subject", "Reset Password LinkhedIn")
	msg.SetBody("text/html", `You recently took steps to reset the password for your LinkhedIn Account. 
	Click on the link below to reset your password.
	: <br> http://127.0.0.1:5173/resetPassword/`+link)

	n := gomail.NewDialer("smtp.gmail.com", 587, "MneachDev@gmail.com", "qpdwvwkkyxovxtue")

	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}

func GenerateRandomLinkActivation() string {

	var link bytes.Buffer
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 35; i++ {
		randomIndex := rand.Intn(len(charset))
		link.WriteByte(charset[randomIndex])
	}

	return link.String()
}

func CreateActiveLink(db *gorm.DB, ctx context.Context, userID string, linkActivation string) (*model.ActivateAccount, error) {
	modelActivationLink := &model.ActivateAccount{
		ID:     linkActivation,
		UserID: userID,
	}

	return modelActivationLink, db.Create(modelActivationLink).Error
}
