package mailgun

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

var yourDomain string = "sandbox707dabe9e9bd429a8c7066f997b2b7a0.mailgun.org"
var privateAPIKey string = "2b33642558ee30e3fdf34f9aa66422d1-2ac825a1-9ee2e972"

type MailgunAPI struct {
}

func SendAPI() {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	sender := "dwisabrinakadek00@gmail.com"
	subject := "Ticket from Kemahin"
	body := "Hello from Mailgun Go!"
	recipient := "dwisabrina10@gmail.com"

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
