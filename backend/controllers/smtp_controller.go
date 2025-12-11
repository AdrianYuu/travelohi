package controllers

import (
	"os"
	"strconv"

	"github.com/AdrianYuu/TraveloHI_TPA_Web_AY/models"
	"gopkg.in/gomail.v2"
)

var HOST = os.Getenv("SMTP_HOST")
var PORT, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
var EMAIL = os.Getenv("SMTP_EMAIL")
var PASSWORD = os.Getenv("SMTP_PASSWORD")

func SendWelcomeEmail(emailReceiver string) {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_EMAIL"))
	m.SetHeader("To", emailReceiver)
	m.SetHeader("Subject", "Welcome to TraveloHI")
	m.SetBody("text/html", "Thank you for registering to TraveloHI. Hope you have a good day!")

	d := gomail.NewDialer(HOST, PORT, EMAIL, PASSWORD)
	d.DialAndSend(m)
}

func SendOTPEmail(emailReceiver string, OTPCode string){
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_EMAIL"))
	m.SetHeader("To", emailReceiver)
	m.SetHeader("Subject", "TraveloHI OTP Code")
	m.SetBody("text/html", "This is your OTP Code " + OTPCode + ". Please use it wisely!")

	d := gomail.NewDialer(HOST, PORT, EMAIL, PASSWORD)
	d.DialAndSend(m)
}

func SendPaymentEmail(emailReceiver string, message string){
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_EMAIL"))
	m.SetHeader("To", emailReceiver)
	m.SetHeader("Subject", "TraveloHI Payment Information")
	m.SetBody("text/html", message)

	d := gomail.NewDialer(HOST, PORT, EMAIL, PASSWORD)
	d.DialAndSend(m)
}

func SendBroadcastEmail(broadcast models.Broadcast){
	emails := GetAllSubscribedEmails()
	m := gomail.NewMessage()
	d := gomail.NewDialer(HOST, PORT, EMAIL, PASSWORD)

	for _, email := range emails {
		m.SetHeader("From", os.Getenv("SMTP_EMAIL"))
		m.SetHeader("To", email)
		m.SetHeader("Subject", broadcast.Title)
		m.SetBody("text/html", broadcast.Message)
		d.DialAndSend(m)
	}
}