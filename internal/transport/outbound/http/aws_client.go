package http_out

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	"one-stop/internal/config"
)

type AWSClient struct {
	env  *config.Environment
	sess *ses.SES
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

func (ac AWSClient) SendEmail(recipient string, category string) error {

	//Build html template
	HtmlBody := ""

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
