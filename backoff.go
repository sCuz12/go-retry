package retry

import "time"

type BackoffStrategy func(attempt int ,initialDelay time.Duration) time.Duration

var Fixed BackoffStrategy = func(attempt int, d time.Duration) time.Duration {
	return d
}

var Linear BackoffStrategy = func(attempt int, d time.Duration) time.Duration {
	return d * time.Duration(attempt+1)
}

var Exponential BackoffStrategy = func(attempt int, d time.Duration) time.Duration {                                                                                                                              
      return d * (1 << attempt)  // d * 2^attempt
}  

var TestBackoff BackoffStrategy = func(attempt int, initialDelay time.Duration) time.Duration {
	return 100 * time.Second
}
