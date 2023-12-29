package http_out

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"os"

	"github.com/mailersend/mailersend-go"

	"one-stop/internal/config"
)

type EmailData struct {
	FirstName string
	Category  string
	UserName  string
	UserEmail string
	UserMsg   string
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

func (mc MailerSendClient) SendEmail(ctx context.Context, recipient, supplierName, name, email, category, msg string) error {

	subject := "One Stop - A new user is interested in your product"
	html, err := getHTMLTemplate(supplierName, name, category, email, msg)
	if err != nil {
		return fmt.Errorf("error getting template: %w", err)
	}

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

	_, err = mc.msc.Email.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}

func getHTMLTemplate(firstName, userName, category, userEmail, msg string) (string, error) {
	var templateBuffer bytes.Buffer

	data := EmailData{
		FirstName: firstName,
		Category:  category,
		UserName:  userName,
		UserEmail: userEmail,
		UserMsg:   msg,
	}

	htmlData, err := os.ReadFile("supplier-template.html")
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	htmlTemplate := template.Must(template.New("email.html").Parse(string(htmlData)))

	err = htmlTemplate.ExecuteTemplate(&templateBuffer, "email.html", data)
	if err != nil {
		return "", fmt.Errorf("error applying template: %w", err)
	}

	return templateBuffer.String(), nil
}
