package http_out

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	"one-stop/internal/config"
)

type AWSClient struct {
	env  *config.Environment
	sess *ses.SES
}

type EmailData struct {
	FirstName string
	Category  string
	UserName  string
	UserEmail string
}

func NewAwsClient(env *config.Environment) (*AWSClient, error) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2")},
	)
	if err != nil {
		return nil, err
	}

	svc := ses.New(sess)

	return &AWSClient{
		env:  env,
		sess: svc,
	}, nil

}

func (ac AWSClient) SendEmail(recipient, supplierName, name, email, category string) error {

	//Build html template
	HtmlBody := getHTMLTemplate(supplierName, name, category, email)

	Sender := ac.env.SesSenderAddress

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(HtmlBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(fmt.Sprintf("New user is enquiring about %s", category)),
			},
		},
		Source: aws.String(Sender),
	}

	_, err := ac.sess.SendEmail(input)
	if err != nil {
		return fmt.Errorf("problem sending email to %s: %w", recipient, err)
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
	htmlData, err := os.ReadFile("supplier-template.html")
	htmlTemplate := template.Must(template.New("email.html").Parse(string(htmlData)))

	err = htmlTemplate.ExecuteTemplate(&templateBuffer, "email.html", data)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	return templateBuffer.String()
}
