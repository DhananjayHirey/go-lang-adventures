package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens chan struct{}
}

func NewRateLimiter(capacity int, refillInterval time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, capacity),
	}

	for i := 0; i < capacity; i++ {
		rl.tokens <- struct{}{}
	}

	ticker := time.NewTicker(refillInterval)

	go func() {
		for range ticker.C {
			select {
			case rl.tokens <- struct{}{}:
				fmt.Print("Token added\n")
			default:
				fmt.Print("Bucket is full\n")

			}
		}
	}()
	return rl
}

func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false

	}
}

func main() {
	limiter := NewRateLimiter(3, 4*time.Second)
	for i := 0; i < 30; i++ {
		fmt.Println("Request allowed:", limiter.Allow())
		time.Sleep(time.Second)
	}
}
