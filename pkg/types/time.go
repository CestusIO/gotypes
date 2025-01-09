package types

import "time"

//TimeProvider can be used to inject times's for testing
type TimeProvider interface {
	Now() time.Time
}

//DefaultTimeProvider is a provider of times'
var DefaultTimeProvider TimeProvider = defaultTimeProvider{}

type defaultTimeProvider struct {
}

// Now returns the currentTime
func (p defaultTimeProvider) Now() time.Time {
	return time.Now()
}

// ProvideTimeProvider provides a Timeprovider
func ProvideTimeProvider() *defaultTimeProvider {
	return &defaultTimeProvider{}
}
