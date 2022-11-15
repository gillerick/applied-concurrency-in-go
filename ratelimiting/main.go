package main

import (
	"context"
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

type APIConnection struct{}

func Open() *APIConnection {
	return &APIConnection{}
}

// ReadFile is an endpoint for reading files
func (a *APIConnection) ReadFile(ctx context.Context) error {
	//ToDO: Add some work here
	return nil
}

// ResolveAddress is an endpoint for resolving a domain name to an IP address
func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	//TODO: Add some work here
	return nil
}
