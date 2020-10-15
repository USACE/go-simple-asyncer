package asyncer

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SQSAsyncer implements the Asyncer Interface for SNS Messages
type SQSAsyncer struct {
	Local     bool
	Endpoint  string
	QueueURL  string
	QueueName string
}

// Name returns name of Asyncer
func (a SQSAsyncer) Name() string {
	return "AWSSQS"
}

func (a SQSAsyncer) getSQS() *sqs.SQS {
	if a.Local {
		return sqs.New(
			session.New(),
			&aws.Config{
				Endpoint: aws.String(a.Endpoint),
				Region:   aws.String("local"),
			},
		)
	}
	return sqs.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})))
}

func (a SQSAsyncer) getQueueURL(svc *sqs.SQS) (*string, error) {
	if a.Local {
		return &a.QueueURL, nil
	}
	// Build Queue URL using AWS Credential Information
	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &a.QueueName,
	})
	if err != nil {
		return nil, err
	}
	return result.QueueUrl, nil
}

// CallAsync implements the Asyncer interface
func (a SQSAsyncer) CallAsync(payload []byte) error {

	svc := a.getSQS()
	queueURL, err := a.getQueueURL(svc)
	if err != nil {
		return err
	}
	params := &sqs.SendMessageInput{
		MessageBody: aws.String(string(payload)), // Required
		QueueUrl:    queueURL,                    // Required
	}
	if _, err := svc.SendMessage(params); err != nil {
		return err
	}
	return nil
}
