package retry

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)


func TestDoContext(t *testing.T) {
	ctx , cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	defer cancel()

	fn := func(ctx context.Context) error {
		fmt.Println("Hello from function")
		return nil
	}

	err := DoWithContext(ctx,fn)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWithError(t *testing.T) {

	ctx , cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	defer cancel()

	calls := 0
	fn := func(ctx context.Context) error {
		calls++
		fmt.Println("Hello from function")
		return fmt.Errorf("Something failed")
	}

	
	err := DoWithContext(ctx,fn) 

	if err == nil {
		t.Fatal("Should have an error")
	}
	if calls != 3 {                                                                               
          t.Fatalf("expected 3 attempts, got %d", calls)
    }  
}

func TestWithMaxAttempts(t *testing.T) {
	ctx , cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	calls := 0
	fn := func(ctx context.Context) error {
		calls++
		fmt.Println("Hello from function")
		return fmt.Errorf("Something failed")
	}

	options := WithMaxAttempts(5)

	err := DoWithContext(ctx,fn,options) 

	if err == nil {
		t.Fatal("Should have an error")
	}
	if calls != 5 {                                                                               
          t.Fatalf("expected 5 attempts, got %d", calls)
    }  
	
}

func TestWithFixedStrategy(t *testing.T) {
	ctx , cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	optionInitial := WithInitialDelay(time.Millisecond * 100)
	options := WithBackoff(Fixed)

	calls := 0
	fn := func(ctx context.Context) error {
		calls++
		fmt.Println("Hello from test with fixed strategy backoff ")
		return fmt.Errorf("Something failed")	
	}

	err := DoWithContext(ctx,fn,optionInitial,options)

	if err == nil {
		t.Fatalf("Expected error")
	}

	if calls != 3 {
		t.Fatalf("Expected 3 calls")
	}
}


func TestWithTestStrategyDeadline(t *testing.T) {
	ctx , cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	options := WithBackoff(TestBackoff)

	calls := 0 
	

	fn := func(ctx context.Context) error {
		calls++
		fmt.Println("fired")
		return fmt.Errorf("Something went wrong")
	}

	err := DoWithContext(ctx, fn, options)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal("The error should be context deadline ")
	}


}


func TestWithMaxDelay(t *testing.T) {
	const maxDelay = 10 * time.Millisecond

	cfg := defaultConfig()
	WithMaxDelay(maxDelay)(cfg)
	WithBackoff(TestBackoff)(cfg)

	delay := delayForAttempt(cfg, 0)
	if delay != maxDelay {
		t.Fatalf("expected delay to be capped at %v, got %v", maxDelay, delay)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	operationErr := errors.New("simulated operation failure")
	calls := 0

	fn := func(context.Context) error {
		calls++
		return operationErr
	}

	err := DoWithContext(
		ctx,
		fn,
		WithMaxDelay(maxDelay),
		WithBackoff(TestBackoff),
	)

	if !errors.Is(err, operationErr) {
		t.Fatalf("expected operation error, got %v", err)
	}
	if calls != 3 {
		t.Fatalf("expected 3 attempts, got %d", calls)
	}
}
