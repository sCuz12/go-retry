package retry

import "time"

type config struct {
	maxAttempts int           // how many times to try total
	fixedDelay  time.Duration //how long to wait between attempts
}

type Option func(*config)

func defaultConfig() *config {
	return &config{
		maxAttempts: 3,
		fixedDelay:  10 * time.Millisecond,
	}
}
func WithMaxAttempts(n int) Option {
	return func(c *config) { c.maxAttempts = n }
}

func WithFixedDelay(d time.Duration) Option {
	return func(c *config) { c.fixedDelay = d }
}
