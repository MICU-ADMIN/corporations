```go

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

```

```mermaid

Here is the Mermaid Markdown overview for the Go file you provided:
```mermaid
graph LR
    A[Go Package] --> B[Import Statements]
    B --> C[Logging]
    C --> D[SendEmail Function]
    D --> E[API Key]
    E --> F[Client Instantiation]
    F --> G[SendEmail Request]
    G --> H[Emails Endpoint]
    H --> I[Send Email Method]
    I --> J[Return Values]
    J --> K[Error Handling]
    K --> L[Logging Statements]
```
This overview shows the relationships between the different components of the Go file, using arrows to represent the dependencies and interactions between them.

* The Go package is the top-level module, which contains the import statements and the SendEmail function.
* The import statements bring in the necessary packages, including the `log` package for logging and the `resend-go` package for sending emails.
* The logging function is used to log messages to the console.
* The SendEmail function takes four parameters: `to`, `

```
