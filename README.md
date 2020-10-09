# go-simple-asyncer

## Warning: Work in Progress

Nothing fancy, a simple library with a goal of being able to switch between Async services like AWS Lambda, AWS SNS, Various Message Queues, or "Mock" for testing without changing underlying code.  This is accomplished using a common interface with a `CallAsync(functionName string, payload []byte) error` method. Here's what it looks like in practice:

```
// myAsyncer creates a new asyncer
myAsyncer, err := asyncer.NewAsyncer(
    asyncer.Config{Engine: "AWSSNS", Topic: "my-sns-topic"},
)
// Use myAsyncer
if err := ae.CallAsync("corpsmap-cumulus-packager", payload); err != nil {
    return nil, err
}
```

Note: It's helpful to pair this library with something like https://github.com/kelseyhightower/envconfig to store asyncer configuration strings as environment variables.

### Amazon Web Services SNS Asyncer

### Amazon Web Services SQS Asyncer (In Progress...)
