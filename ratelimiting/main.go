package main

import "context"

func main() {

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
