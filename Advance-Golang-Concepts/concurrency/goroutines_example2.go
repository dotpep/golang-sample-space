package main

import (
	"fmt"
	"time"
)

const (
	IterCount = 100
	TimeSleep = 10
)

var opTookTime string

func say(s string) {
	for i := 0; i < IterCount; i++ {
		time.Sleep(TimeSleep * time.Millisecond)
		fmt.Println(s)
	}
}

func usual() {
	startNow := time.Now()

	say("world")
	say("hello")

	opTookTime = fmt.Sprintf("(usual) This operation took: %s", time.Since(startNow))

	//fmt.Println("This operation took: ", time.Since(startNow))
}

func concurrent() {
	startNow := time.Now()

	go say("world")
	say("hello")

	opTookTime += fmt.Sprintf("\n(concurrent) This operation took: %s", time.Since(startNow))

	//fmt.Println("This operation took: ", time.Since(startNow))
}

func main() {
	usual()
	fmt.Println("---")
	concurrent()
	fmt.Println("---")
	//fmt.Println("usual 1: ", opTookTime)
	//fmt.Println("concurrent 2:", opTookTime)
	fmt.Println(opTookTime)

	// TODO: Implement Decorator Pattern for this (startNow := time.Now() - This operation took: time.Since(startNow))
	// TODO: clear this code!
}
