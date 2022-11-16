package main

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"os"
	"sync"
	"time"
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
	apiLimit,
	diskLimit,
	networkLimit RateLimiter
}

func Open() *APIConnection {
	return &APIConnection{
		// 1. We define our limit per second with no burstiness
		// 2. We define our limit per minute with a burstiness of 10 to give the users their initial pool
		apiLimit: MultiLimiter(
			rate.NewLimiter(rate.Every(time.Second), 2),
			rate.NewLimiter(rate.Every(time.Minute), 10),
		),
		diskLimit:    MultiLimiter(rate.NewLimiter(rate.Limit(1), 1)), // One read per second
		networkLimit: MultiLimiter(rate.NewLimiter(rate.Every(time.Second), 3)),
	}
}

// ReadFile is an endpoint for reading files
func (a *APIConnection) ReadFile(ctx context.Context) error {
	// We wait on the rate limiter to have enough access tokens for us to complete our request
	if err := MultiLimiter(a.apiLimit, a.diskLimit).Wait(ctx); err != nil {
		return err
	}
	//Pretend to do some work here
	return nil
}

// ResolveAddress is an endpoint for resolving a domain name to an IP address
func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := MultiLimiter(a.apiLimit, a.networkLimit).Wait(ctx); err != nil {
		return err
	}
	//Pretend to do some work here
	return nil
}
