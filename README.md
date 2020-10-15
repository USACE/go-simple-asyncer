# go-simple-asyncer

## Warning: Work in Progress

Nothing fancy, a simple library with a goal of being able to switch between Async services like AWS Lambda, AWS SNS, Various Message Queues, or "Mock" for testing without changing underlying code.  This is accomplished using a common interface with a `CallAsync(functionName string, payload []byte) error` method. Here's what it looks like in practice:

```
// myAsyncer creates a new asyncer
myAsyncer, err := asyncer.NewAsyncer(
    asyncer.Config{Engine: "AWSSNS", Target: "my-sns-topic"},
)
// Use myAsyncer
if err := ae.CallAsync("corpsmap-cumulus-packager", payload); err != nil {
    return nil, err
}
```

Note: It's helpful to pair this library with something like https://github.com/kelseyhightower/envconfig to store asyncer configuration strings as environment variables.

### Amazon Web Services SNS Asyncer

### Amazon Web Services SQS Asyncer (In Progress...)

If only a queue name is provided as `Target`, for example `myqueue1`, it is assumed that this is running on Amazon Web Services against SQS. Environment variables provided at runtime, such as "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_REGION" will be necessary to correctly build appropriate SQS service URLs.

Prefixing `Target` with `local/` and providing an absolute URL (i.e. `Target: local/http://localhost:9324/queues/queue1` allows use of a SQS-api compliant message queue such as ElasticMQ or LocalStack.
