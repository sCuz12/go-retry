package retry

import (
	"context"
	"time"
)

 func Do(fn func() error, opts ...Option) error {
     return DoWithContext(context.Background(), func(_ context.Context) error {
         return fn()
     }, opts...)
 }
 

func DoWithContext(ctx context.Context,fn func(context.Context) error,opts ...Option) error {
	cfg := defaultConfig()
	for _,opt := range opts {
		opt(cfg)
	}
	var err error

	for i:=0 ; i<cfg.maxAttempts;i++ {
		//respect cancellation global
		if ctx.Err() != nil {
			return ctx.Err()
		}
		err = fn(ctx)

		if err == nil {
			return nil
		}

		if i == cfg.maxAttempts - 1 {
			break
		}
		select {
		case <- ctx.Done():
			return ctx.Err()
		case <-time.After(cfg.fixedDelay):

		}
	}
	return err
}
