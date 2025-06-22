package utils

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"os"
)

func SendInvoiceMail(toEmail, subject, body string) {
	from := mail.NewEmail("GoInvoice", "no-reply@goinvoice.com")
	to := mail.NewEmail("User", toEmail)
	message := mail.NewSingleEmail(from, subject, to, body, body)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	if err != nil {
		log.Println("Sendgrid error:", err)
	}
}
