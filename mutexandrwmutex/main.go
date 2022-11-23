package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		// 1. We request exclusive use of the critical section—in this case the count variable—guarded by a Mutex, lock
		lock.Lock()
		// 2. We indicate that we’re done with the critical section lock is guarding
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	//Increment
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	//Decrement
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}

	//The producer function’s second parameter is of the type sync.Locker.
	//This interface has two methods, Lock and Unlock, which the Mutex and RWMutex types satisfy.
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			//We make the producer sleep for one second to make it less active than the observer goroutines
			time.Sleep(1)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var waitGroup sync.WaitGroup
		waitGroup.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&waitGroup, mutex)

		for i := count; i > 0; i-- {
			go observer(&waitGroup, rwMutex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutext\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}

	wg.Wait()
	fmt.Println("Wait-group complete")

}
