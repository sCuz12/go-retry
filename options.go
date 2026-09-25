package retry

import "time"

type config struct {
	maxAttempts  int           // how many times to try total
	fixedDelay   time.Duration //how long to wait between attempts
	backoff      BackoffStrategy
	initialDelay time.Duration // base delay fed into strategy
	maxDelay     time.Duration // cap - no delay ever exceed this duration
}

// we this we give the ability to consumers to customize behaviour without exposing conifg
type Option func(*config)

func defaultConfig() *config {
	return &config{
		maxAttempts:  3,
		fixedDelay:   10 * time.Millisecond,
		initialDelay: 3 * time.Millisecond,
		maxDelay:     30 * time.Second,
		backoff:      Exponential, // ← default strategy
	}
}
func WithMaxAttempts(n int) Option {
	return func(c *config) { c.maxAttempts = n }
}

func WithBackoff(backoff BackoffStrategy) Option {
	return func(c *config) { c.backoff = backoff }
}

func WithInitialDelay(d time.Duration) Option {
	return func(c *config) { c.initialDelay = d }
}

func WithMaxDelay(d time.Duration) Option {
	return func(c *config) { c.maxDelay = d }
}
