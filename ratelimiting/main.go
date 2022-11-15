package main

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"os"
	"sync"
)

func main() {
	defer log.Print("Done")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot read file %v", err)
			}
			log.Printf("ReadFile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot resolve address %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}

	wg.Done()

}

type APIConnection struct {
	rateLimiter *rate.Limiter
}

func Open() *APIConnection {
	return &APIConnection{
		// 1. We set the rate limit for all API connections to one event per second
		rateLimiter: rate.NewLimiter(rate.Limit(1), 1),
	}
}

// ReadFile is an endpoint for reading files
func (a *APIConnection) ReadFile(ctx context.Context) error {
	// 2. We wait on the rate limiter to have enough access tokens for us to complete our request
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	//Pretend to do some work here
	return nil
}

// ResolveAddress is an endpoint for resolving a domain name to an IP address
func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	//Pretend to do some work here
	return nil
}
