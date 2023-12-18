package http_out

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"

	"github.com/mailersend/mailersend-go"

	"one-stop/internal/config"
)

type EmailData struct {
	FirstName string
	Category  string
	UserName  string
	UserEmail string
}

type MailerSendClient struct {
	env *config.Environment
	msc *mailersend.Mailersend
}

func NewMailerSendClient(env *config.Environment) *MailerSendClient {

	ms := mailersend.NewMailersend(env.MailerSendApiKey)
	return &MailerSendClient{
		env: env,
		msc: ms,
	}
}

func (mc MailerSendClient) SendEmail(ctx context.Context, recipient, supplierName, name, email, category string) error {

	subject := "One Stop - A new user is interested in your product"
	html := getHTMLTemplate(supplierName, name, category, email)

	from := mailersend.From{
		Name:  "One Stop",
		Email: "info@onestopbuildingshop.co.uk",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  supplierName,
			Email: recipient,
		},
	}
	message := mc.msc.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)

	_, err := mc.msc.Email.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}

func getHTMLTemplate(firstName, userName, category, userEmail string) string {
	var templateBuffer bytes.Buffer

	data := EmailData{
		FirstName: firstName,
		Category:  category,
		UserName:  userName,
		UserEmail: userEmail,
	}
	htmlData := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
</head>
<body>
<p> Hello {{.FirstName}} </p>
<p> The below user has enquired about {{.Category}} products </p>
<p> Here are their details </p>
<p> {{.UserName}}, {{.UserEmail}}</p>
<p> Best Regards from, </p>
<p> OneStop Building Supples </p>

</body>
</html>`

	//, err := os.ReadFile("supplier-template.html")
	htmlTemplate := template.Must(template.New("email.html").Parse(htmlData))

	err := htmlTemplate.ExecuteTemplate(&templateBuffer, "email.html", data)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	return templateBuffer.String()
}
