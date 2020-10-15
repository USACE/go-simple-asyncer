package asyncer

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SQSAsyncer implements the Asyncer Interface for SNS Messages
type SQSAsyncer struct {
}

// Name returns name of Asyncer
func (a SQSAsyncer) Name() string {
	return "AWS SQS"
}

// CallAsync implements Asyncer interface for AWS Lambda
// target should be
func (a SQSAsyncer) CallAsync(target string, payload []byte) error {

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	result, err := svc.GetQueueUrl(
		&sqs.GetQueueUrlInput{
			QueueName: &target,
		},
	)
	if err != nil {
		return err
	}

	params := &sqs.SendMessageInput{
		MessageBody: aws.String(string(payload)), // Required
		QueueUrl:    result.QueueUrl,             // Required
	}

	if _, err := svc.SendMessage(params); err != nil {
		return err
	}

	return nil
}
