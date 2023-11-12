package lib

import (
	"log"

	"github.com/resendlabs/resend-go"
)

func SendEmail(to, from, subject, html string) (bool, error) {
    apiKey := "re_Knpe3rhM_ExDXKXY55GTZjafbQom2Mjh2"

    client := resend.NewClient(apiKey)

    params := &resend.SendEmailRequest{
        From:    from,
        To:      []string{to},
        Subject: subject,
        Html:    html,
    }

    sent, err := client.Emails.Send(params)
    if err != nil {
        return false, err
    }
    log.Println(sent)
    return true, nil  // Return nil for the error since the email was sent successfully
}