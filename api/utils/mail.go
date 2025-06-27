package utils

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"os"
)

func SendInvoiceMail(toEmail, subject, body string) {
	from := mail.NewEmail("GoInvoice", "xp.aprill@gmail.com")
	to := mail.NewEmail("User", toEmail)
	message := mail.NewSingleEmail(from, subject, to, body, body)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println("Sendgrid error:", err)
	} else {
		log.Printf("Sendgrid Response: Status %d, Body: %s\n", response.StatusCode, response.Body)
	}

}
