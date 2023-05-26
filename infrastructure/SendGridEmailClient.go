package infrastructure

import "os"

type SendGridEmailClient struct {
	apiKey string
}

func NewSendGridEmailClient() *SendGridEmailClient {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	client := &SendGridEmailClient{apiKey: apiKey}

	return client
}
