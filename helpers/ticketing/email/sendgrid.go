package email

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Data struct {
	Title      string
	Name       string
	NIM        string
	TicketCode string
	ImgLink    string
	QRLink     string
	TimeEvent  time.Time
	DateEvent  string
	MonthEvent string
}

func SendEmail(id_event, id_mhs int) {
	var (
		eventName  = "event name"
		name       = "mhs name"
		email      = "mhs@email"
		ticketCode = "adjakjdkaj"
	)
	newTemplate := Data{}
	newTemplate.QRLink = fmt.Sprintf("https://api.qrserver.com/v1/create-qr-code/?data=%s&size=220x220&margin=0", ticketCode)

	from := mail.NewEmail("Example User", "test@example.com")
	subject := fmt.Sprintf("Tiket %s", eventName)
	to := mail.NewEmail(name, email)

	plainTextContent := "and easy to do anywhere, even with Go"
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
}

// fmt.Println(time.Now().UTC().Format("Jan")) // Aug

//     t := time.Now()
//     str := fmt.Sprintf("%d %s %02d", t.Year(), t.Month().String()[:3], t.Day())
//     fmt.Println(str) // 2016 Aug 03
