package retry

import (
	"context"
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