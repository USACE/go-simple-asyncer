package asyncer

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"strings"
)

// Config holds configuration information to build an Asyncer
type Config struct {
	Engine string
	Target string
}

// NewAsyncer returns a concrete asyncer
func NewAsyncer(cfg Config) (Asyncer, error) {
	switch cfg.Engine {
	case "AMQP":
		return AMQPAsyncer{}, nil
	case "AWSLAMBDA":
		return LambdaAsyncer{}, nil
	case "AWSSQS":
		if strings.ToUpper(cfg.Target[:6]) == "LOCAL/" {
			url, err := url.Parse(cfg.Target[6:])
			if err != nil {
				return nil, err
			}
			_, queueName := path.Split(url.Path)
			return SQSAsyncer{
				Local:     true,
				Endpoint:  fmt.Sprintf("%s://%s", url.Scheme, url.Host),
				QueueName: queueName,
				QueueURL:  cfg.Target[6:],
			}, nil
		}
		return SQSAsyncer{Local: false, QueueName: cfg.Target}, nil
	case "AWSSNS":
		if cfg.Target == "" {
			return nil, errors.New(
				"Engine 'AWSSNS' requires a target (SNS Topic)",
			)
		}
		return SNSAsyncer{Topic: cfg.Target}, nil
	default:
		return MockAsyncer{}, nil
	}
}
