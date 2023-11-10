package resend

import (
	"log"
	"os"
	"time"

	"github.com/resendlabs/resend-go"
)


type EmailResponse struct {
	Object     string    `json:"object"`
	ID         string    `json:"id"`
	To         []string  `json:"to"`
	From       string    `json:"from"`
	CreatedAt  time.Time `json:"created_at"`
	Subject    string    `json:"subject"`
	HTML       string    `json:"html"`
	Text       *string   `json:"text"`
	BCC        []*string `json:"bcc"`
	CC         []*string `json:"cc"`
	ReplyTo    []*string `json:"reply_to"`
	LastEvent  string    `json:"last_event"`
}

func GetEmail() (resend.Email, error) {
	client := resend.NewClient("re_123456789")

	email, err := client.Emails.Get("4ef9a417-02e9-4d39-ad75-9611e0fcc33c")
	if err != nil {
		log.Fatal("failed to get email", err)
	}

	return email, nil
}





func SendEmail() {
	pwd, _ := os.Getwd()
f, _ := os.ReadFile(pwd + "./resources/invoice.pdf")

client := resend.NewClient("re_123456789")
 if client != nil {
    log.Println("Error loading client")
 }

pdfAttachment := &resend.Attachment{
    Content:  string(f),
    Filename: "invoice.pdf",
}

params := &resend.SendEmailRequest{
    From:        "Acme <onboarding@resend.dev>",
    To:          []string{"delivered@resend.dev"},
    Text:        "it works!",
    Subject:     "hello world",
    Headers:     map[string]string{
      "X-Entity-Ref-ID": "123456789",
    },
    Attachments: []resend.Attachment{*pdfAttachment},
}
 if params != nil {
    log.Println("Error loading params")
 }

}
