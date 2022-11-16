package main

import (
	"context"
	"golang.org/x/time/rate"
	"sort"
)

// RateLimiter defines an interface that allows MultiLimiter to recursively define other MultiLimiter instances
type RateLimiter interface {
	Wait(ctx context.Context) error
	Limit() rate.Limit
}

func MultiLimiter(limiters ...RateLimiter) *multiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	// We implement an optimization and sort by the Limit() of each RateLimiter
	sort.Slice(limiters, byLimit)
	return &multiLimiter{limiters: limiters}
}

type multiLimiter struct {
	limiters []RateLimiter
}

func (m *multiLimiter) Wait(ctx context.Context) error {
	for _, l := range m.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (m *multiLimiter) Limit() rate.Limit {
	//Since we sort the child RateLimiter instances when multiLimiter is instantiated,
	//we can simply return the most restrictive limit, which will be the first element in the slice
	return m.limiters[0].Limit()
}
