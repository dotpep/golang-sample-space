package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// Mutexes
// stands for `Mutual Exclusion`

// mutexes allow is to lock access to data.
// This ensures that we can cotrol which goroutine can access certain data at which time.

// standard Go library provides build-in implementation:
// mutex type is: `sync.Mutex` and its two methods:
// `.Lock()` and `.Unlock()`

// we can protect a block of code by surrounding it with a call to `Lock` and `Unlock`
// as shown `protected()` method below:
// is good practice using `defer` to `Unlock`

// func protected() {
// 	mux.Lock()
// 	defer mux.Unlock()
// 	// rest of function is protected
// 	// any other calls to `mux.Lock()` will block
// }

// Maps are not Thread-Safe
// maps are `not` safe for concurrent use! if you have multiple goroutines
// accessing the same map, and at least one of them is writing to map,
// you must lock your maps with a mutex.

// RW Mutex

// `sync.RWMutex` Read/Write Mutex include two additional methods:
// `RLock()` and `RUnlock()`

// `RWMutex` can help with performance if we have a read-intensive process.
// Many goroutines can safely read from map at the same time
// (multiple `RLock()` calls can happen simultaneously).
// However, only one goroutine can hold `Lock()` and all `RLock()`'s will also be excluded.

// only `one` writers can access a `RWMutex` at once
// `infinite` readers can access a `RWMutex` at once

// Mutexes Example 1

type safeCounter struct {
	counts map[string]int
	//mux    *sync.Mutex
	mux *sync.RWMutex
}

type emailTest struct {
	email string
	count int
}

func (sc safeCounter) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key) // writing dangerous thing
}

func (sc safeCounter) val(key string) int {
	sc.mux.RLock()
	defer sc.mux.RUnlock()
	return sc.counts[key] // reading is also dangerous thing
}

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func testMutex1(sc safeCounter, emailTests []emailTest) {
	emails := make(map[string]struct{})

	var wg sync.WaitGroup
	for _, emailT := range emailTests {
		emails[emailT.email] = struct{}{}
		for i := 0; i < emailT.count; i++ {
			wg.Add(1)
			go func(emailT emailTest) {
				sc.inc(emailT.email)
				wg.Done()
			}(emailT)
		}
	}
	wg.Wait()

	emailsSorted := make([]string, 0, len(emails))
	for email := range emails {
		emailsSorted = append(emailsSorted, email)
	}
	sort.Strings(emailsSorted)

	for _, email := range emailsSorted {
		fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
	}
	fmt.Println("==================")
}

// Mutex Review
// Principle problem that mutexes help us avoid is concurrent read/write problem.
// This problem arises when one thread is writing to a variable,
// while another thread is reading from that same variable at the same time.
// When this happens Go program will panic because reader could be reading bad data
// while it's being mutated in place.

// mutex can `Lock` only `one` threads at once!
// mutex is used to safaly access a data structure concurrently.

func main() {
	// Mutexes Example 1
	sc := safeCounter{
		counts: make(map[string]int),
		mux:    &sync.RWMutex{},
		//mux:    &sync.Mutex{},
	}
	testMutex1(sc, []emailTest{
		{
			email: "john@example.com",
			count: 23,
		},
		{
			email: "john@example.com",
			count: 29,
		},
		{
			email: "jill@example.com",
			count: 31,
		},
		{
			email: "jill@example.com",
			count: 67,
		},
	})
	testMutex1(sc, []emailTest{
		{
			email: "kaden@example.com",
			count: 23,
		},
		{
			email: "george@example.com",
			count: 126,
		},
		{
			email: "kaden@example.com",
			count: 31,
		},
		{
			email: "george@example.com",
			count: 453,
		},
	})
}
