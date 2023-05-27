package infrastructure

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"os"
)

type SendGridEmailClient struct {
	apiKey string
}

func NewSendGridEmailClient(apiKey string) *SendGridEmailClient {
	client := &SendGridEmailClient{apiKey: apiKey}

	return client
}

func (emailClient *SendGridEmailClient) Send(recipients []string, htmlContent string) {
	if len(recipients) == 0 {
		return
	}

	from := mail.NewEmail("Genesis School Applicant", "genesis.applicant@example.com")
	firstTo := mail.NewEmail("Rate Recipient", recipients[0])
	subject := "Current BTC to UAH rate"
	message := mail.NewSingleEmail(from, subject, firstTo, "", htmlContent)

	for i := 1; i < len(recipients); i++ {
		personalization := mail.NewPersonalization()
		personalization.AddTos(mail.NewEmail("Rate Recipient", recipients[i]))
		message.AddPersonalizations(personalization)
	}

	client := sendgrid.NewSendClient(os.Getenv(emailClient.apiKey))

	response, err := client.Send(message)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
