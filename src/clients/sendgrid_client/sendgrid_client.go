package sendgrid_client

import (
	"github.com/sendgrid/sendgrid-go"
	"os"
)

var (
	sendGridApiKey = "SENDGRID_API_KEY"
	client         *sendgrid.Client
)

func init() {
	if client == nil {
		client = sendgrid.NewSendClient(os.Getenv(sendGridApiKey))
	}
}

func Get() *sendgrid.Client {
	return client
}
