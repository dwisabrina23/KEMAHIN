package sendgrid

import (
	"errors"
	"fmt"
	"kemahin/businesses/sendgrids"
	"log"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendAPI struct {
	httpClient http.Client
}

func NewSendAPI() sendgrids.Repository {
	return &SendAPI{
		httpClient: http.Client{},
	}
}

func (repo *SendAPI) Send(emailData *sendgrids.Domain) (sendgrids.Domain, error) {

	var eventName, UserName, email, ticketCode string
	newTemplate := sendgrids.Domain{}
	newTemplate.QRLink = fmt.Sprintf("https://api.qrserver.com/v1/create-qr-code/?data=%s&size=220x220&margin=0", ticketCode)

	from := mail.NewEmail(UserName, email)
	subject := fmt.Sprintf("Tiket %s", eventName)
	to := mail.NewEmail(UserName, email)

	plainTextContent := fmt.Sprintf("terima kasih sudah membeli tiket %s. Pembayaranmu sudah kami terima, berikut kode tiketnya!", eventName)
	htmlContent := fmt.Sprintf("<strong>Kode tiket: %s </strong>", ticketCode)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return sendgrids.Domain{}, errors.New("success to send ticket")
}
