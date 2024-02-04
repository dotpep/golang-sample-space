package main

import (
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin", "Alex", "Matthew"}

	// Request aa greeting message.
	//message1, err := greetings.Hello("Gladys")
	message2, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	//fmt.Println(message1)
	//fmt.Println(message2)

	// unpack and retrieve map in message2
	//for name, message := range message2 {
	//	fmt.Println(name)
	//	fmt.Println(message)
	//}

	//for _, message := range message2 {
	//	fmt.Println(message)
	//}

	count := len(message2)

	for i := 0; i < count; i++ {
		text := message2[names[i]]
		message := fmt.Sprintf("%v. %v", i+1, text)
		fmt.Println(message)
	}

	//iteratorCountPrint(5)

	// Read a byte from the standard input
	waitForKeyPressToExit(err)
}

func waitForKeyPressToExit(err error) {
	fmt.Println("Press any key to exit...")
	b := make([]byte, 1)
	_, err = os.Stdin.Read(b)
	if err != nil {
		fmt.Println(err)
	}
}

func iteratorCountPrint(count int) {
	for i := 0; i < count; i++ {
		fmt.Println(i + 1)
	}
}
