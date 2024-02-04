package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// you can declare Docstring just comment text on top of function like this:

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Create a message using random format
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprintf(randomFormat())
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	// initialize a map with syntax: make(map[key-type]value-type).
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello finction to get a message for each name.
	// You don't need the index, so you use the Go blank identifier (an underscore) to ignore it.
	// unpacking and looping slice is like python dictionary that have ( key: value ) to loop it it need both for key, value in dicts:
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		// map of received names (as a key) (unpacked with for loop in range names)
		// and with a generated message (as a value) (calling Hello function in for loop while string[] slice names length is 0).
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message format (slice is like an array)
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hello, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
