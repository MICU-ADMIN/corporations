package resend

import (
	"log"

	"github.com/resendlabs/resend-go"
)

func SendEmail(to,from,subject,html) (bool, error) {
    apiKey := "re_Knpe3rhM_ExDXKXY55GTZjafbQom2Mjh2"

    client := resend.NewClient(apiKey)

    params := &resend.SendEmailRequest{
        From:    "onboarding@resend.dev",
        To:      []string{"admin@mosque.icu"},
        Subject: "Hello World",
        Html:    "<p>Congrats on sending your <strong>first email</strong>!</p>",
    }

    sent, err := client.Emails.Send(params)
    if err != nil {
        return false, err
    }
    log.Println(sent)
    return true, err
}
