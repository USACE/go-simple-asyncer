package asyncer

import "log"

// MockAsyncer implements the Asyncer Interface for a mock
type MockAsyncer struct {
	Target string
}

// Name returns asyncer name
func (a MockAsyncer) Name() string {
	return "MOCK"
}

// CallAsync implements Asyncer interface for Mock
func (a MockAsyncer) CallAsync(payload []byte) error {
	log.Printf(
		"ASYNC ENV: %s; FUNCTION: %s;PAYLOAD: %s",
		a.Name(),
		a.Target,
		payload,
	)
	return nil
}
