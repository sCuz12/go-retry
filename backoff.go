package retry

import "time"

type BackoffStrategy func(attempt int, initialDelay time.Duration) time.Duration



func Fixed(_ int, delay time.Duration) time.Duration {
      return delay
}
func Linear (attempt int, d time.Duration) time.Duration {
	return d * time.Duration(attempt+1)
}

func Exponential(attemt int , d time.Duration) time.Duration {
	return d * (1 <<attemt) // d * 2^attempt
}
