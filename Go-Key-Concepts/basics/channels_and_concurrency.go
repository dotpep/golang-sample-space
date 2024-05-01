package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Channels

// Concurrency & Parallelism
// vs
// Synchronous & Sequential

// What is Concurrency
// concurency is the ability
// to perform multiple tasks at the same time.
// Typically, our code is executed one line at a time,
// one after the other.
// This is called `Sequential` execution or `Syncronous` execution.

// Syncronous execution
// --task1-->--task2-->

// Concurrent Execution
// --task1-->
// --task2-->

// if we have multiple cores like 4 in our CPU
// execution will be in the same time
// if we don't have single core
// execute code almost same time with switching between tasks.

// Golang Concurrency
// `go doSomething()`
// this 'go' keyword and given function will execute in concurrent way
// we not able to capture any return values from this function
// and is spawn new `goroutine`

// Concurrency Example 1
func sendEmails(message string, done chan<- bool) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
		done <- true
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func testConcurrency1(message string) {
	done := make(chan bool)
	sendEmails(message, done)
	<-done // Wait for the goroutine to finish
	time.Sleep(time.Microsecond * 500)
	fmt.Println("==================")
}

// Channels are a typed, thread-safe queue.
// Channels allow different goroutines to communicate with each other.

// Create a Channel
// like maps and slices, channels must be created before use.
// They also use same `make` keyword
// ch := make(chan int)

// Send data to a Channel
// ch <- 69
// `<-` operator is called channel operator.
// data flows in direction of arrow
// operator will block until another goroutine is ready to receive value.

// Receive data from a Channel
// v := <-ch

// Channels Example 2
type email struct {
	body string
	date time.Time
}

func filterOldEmails(emails []email) {
	isOldChan := make(chan bool)

	go func() {
		for _, e := range emails {
			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				isOldChan <- true
				continue
			}
			isOldChan <- false
		}
	}()

	isOld := <-isOldChan
	fmt.Println("email 1 is old", isOld)
	isOld = <-isOldChan
	fmt.Println("email 2 is old", isOld)
	isOld = <-isOldChan
	fmt.Println("email 3 is old", isOld)
}

func testChannel2(emails []email) {
	filterOldEmails(emails)
	fmt.Println("==================")
}

// Channels Example 3

// Emty structs are often used as `tokens` in Go programs.
// In this context, a token is a `unary` value.
// `unary` is only one possible value not two like `binary` true and false.
// We don't care what is passed through channel.
// We care when and if it is passed.
// Define unary token channel:
// ch := make(chan struct{})

func waitForDbs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()
	return ch
}

func testChannel3(numDBs int) {
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	time.Sleep(time.Millisecond * 10)
	fmt.Println("All databases are online!")
	fmt.Println("==================")
}

// Channels Example 4

// Buffered Channels
// channels can optionally be buffered.
// You can provide a buffer length as second argument to `make()`
// to create a buffered channel:
// ch := make(chan int, 100)

// sending on a buffered channel only blocks when buffer is full.
// receiving blocks only when buffer is empty.

func addEmailToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

func sendEmails2(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email:", email)
	}
}

func testChannel4(emails ...string) {
	fmt.Printf("Adding %v emails to queue...\n", len(emails))
	ch := addEmailToQueue(emails)
	fmt.Println("Sending emails...")
	sendEmails2(len(emails), ch)
	fmt.Println("==================")
}

// Channels Example 5

// Closing Channels
// channels can be explicitly closed by a sender.

// ch := make(chan int)
// do something stuff with channels
// close(ch)

// Checking if a channel is closed
// similiar to the `ok` value when accessing data in a `map`,
// receivers can chack `ok` value when receiving from a channel to test if a channel was closed.
// val, ok := <-ch
// `ok` is `false` if channel is empty and closed.

// Don't send on a closed Channel
// close channel from sending side

func countReports(numSentCh chan int) int {
	total := 0

	for {
		numSent, ok := <-numSentCh
		if !ok {
			break
		}
		total += numSent
	}

	return total
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf("Sent batch of %v reports\n", numReports)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

func testChannel5(numBatches int) {
	numSentCh := make(chan int)
	go sendReports(numBatches, numSentCh)

	fmt.Println("Start counting...")
	numReports := countReports(numSentCh)
	fmt.Printf("%v reports sent!\n", numReports)
	fmt.Println("==================")
}

// Channels Example 6

// Range keyword for Channels
// similiar to slice, and maps, channels can be ranged over.

// for item := range ch { }
// item is next value received from channel
// This will receive values over channel
// (blocking at each iteration if nothing ne wis there)
// and will exit only when channel is closed.

func concurrentFib(num int) {
	chInts := make(chan int)
	go func() {
		fibonacci(num, chInts)
	}()

	for chNum := range chInts {
		fmt.Println(chNum)
	}
}

func fibonacci(num int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < num; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}

func testChannel6(num int) {
	fmt.Printf("Printing %v numbers...\n", num)
	concurrentFib(num)
	fmt.Println("==================")
}

// Channels Example 7

// Select Statement
// sometimes we have a single goroutine listening to multiple channels
// and want to process data in order it comes through each channel.

// `select` statement is used to listen to multiple channels at the same time.
// It is similar to a `switch` statement but for channels.

// select {
// 	case num, ok := <- chInts:
// 		fmt.Println(i)
// 	case text, ok := <- chStrings:
// 		fmt.Println(s)
// }

// First channel with a value ready to be received will fire and its body will execute.
// If multiple channels are ready at the same time one is chosen randomly.
// `ok` variable in the example above refers to whether or not channel has been closed by sender yet.
func logMessages(chSms, chEmails chan string) {
	for {
		select {
		case sms, ok := <-chSms:
			if !ok {
				return
			}
			logSms(sms)
		case email, ok := <-chEmails:
			if !ok {
				return
			}
			logEmail(email)
		}
	}
}

func logSms(sms string) {
	fmt.Println("SMS: ", sms)
}

func logEmail(email string) {
	fmt.Println("Email: ", email)
}

func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
	chSms = make(chan string)
	chEmails = make(chan string)

	go func() {
		for i := 0; i < len(sms) && i < len(emails); i++ {
			done := make(chan struct{})
			s := sms[i]
			e := emails[i]
			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
			t2 := time.Millisecond * time.Duration(rand.Intn(1000))

			go func() {
				time.Sleep(t1)
				chSms <- s
				done <- struct{}{}
			}()
			go func() {
				time.Sleep(t2)
				chEmails <- e
				done <- struct{}{}
			}()

			<-done
			<-done
			time.Sleep(time.Millisecond * 10)
		}
		close(chSms)
		close(chEmails)
	}()

	return chSms, chEmails
}

func testChannel7(sms, emails []string) {
	fmt.Println("Starting...")

	chSms, chEmails := sendToLogger(sms, emails)

	logMessages(chSms, chEmails)
	fmt.Println("==================")
}

// Channels Example 8

// Select Default Case
// `default` case in a `select` statement executes immediately if no other channel has a value ready.
// `default` case stops `select` statement from blocking

// select {
// case v := <-ch:
// 	// use v
// default:
// 	// receiving from ch would block
// 	// so do something else
// }

// Tickers
// `time.Tick()` is a standard library function that returns a channel that sends a value on a given interval.
// `time.After()` sends a value once after the duration has passed.
// `time.Sleep()` blocks current goroutine for specified amout of time.

// Read-Only Channels
// a channel can be marked as read-only by casting it from a `chan` to a `<- chan` type:

// func main() {
// 	ch := make(chan int)
// 	readCh(ch)
// }
//
// func readCh(ch <-chan int) {
// 	// ch can only be read from
// 	// in this function
// }

// Write-Only Channels
// same goes for write-only channels, but arrow's position moves:

// func writeCh(ch chan<- int) {
// 	// ch can only be written to
// 	// in this function
// }

func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot()
		case <-saveAfter:
			saveSnapshot()
			return
		default:
			waitForData()
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func takeSnapshot() {
	fmt.Println("Taking a backup snapshot...")
}

func saveSnapshot() {
	fmt.Println("All backups saved!")
}

func waitForData() {
	fmt.Println("Nothing to do, waiting...")
}

func testChannel8() {
	snapshotTicker := time.Tick(time.Millisecond * 800)
	saveAfter := time.After(time.Millisecond * 2800)
	saveBackups(snapshotTicker, saveAfter)
	fmt.Println("==================")
}

func main() {
	// Concurrency Example 1
	testConcurrency1("Hello there Stacy!")
	testConcurrency1("Hi there John!")
	testConcurrency1("Hey there Jane!")

	fmt.Println("---")

	// Channels Example 2
	testChannel2([]email{
		{
			body: "Are you going to make it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "I need a break",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What were you thinking?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	testChannel2([]email{
		{
			body: "Yo are you okay?",
			date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Have you heard of that website Boot.dev?",
			date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "It's awesome honestly.",
			date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})

	fmt.Println("---")

	// Channels Example 3
	testChannel3(3)
	testChannel3(4)
	testChannel3(5)

	fmt.Println("---")

	// Channels Example 4
	testChannel4("Hello John, tell Kathy I said hi", "Whazzup bruther")
	testChannel4("I find that hard to believe.", "When? I don't know if I can", "What time are you thinking?")
	testChannel4("She says hi!", "Yeah its tomorrow. So we're good.", "Cool see you then!", "Bye!")

	fmt.Println("---")

	// Channels Example 5
	testChannel5(3)
	testChannel5(4)
	testChannel5(5)
	testChannel5(6)

	fmt.Println("---")

	// Channels Example 6
	testChannel6(10)
	testChannel6(50)
	testChannel6(20)
	testChannel6(30)

	fmt.Println("---")

	// Channels Example 7
	testChannel7(
		[]string{
			"hi friend",
			"What's going on?",
			"Welcome to the business",
			"I'll pay you to be my friend",
		},
		[]string{
			"Will you make your appointment?",
			"Let's be friends",
			"What are you doing?",
			"I can't believe you've done this.",
		},
	)
	testChannel7(
		[]string{
			"this song slaps hard",
			"yooo hoooo",
			"i'm a big fan",
		},
		[]string{
			"What do you think of this song?",
			"I hate this band",
			"Can you believe this song?",
		},
	)

	fmt.Println("---")

	// Channels Example 8
	testChannel8()
}
