package asyncer

// Asyncer interface
type Asyncer interface {
	Name() string
	CallAsync(payload []byte) error
}
