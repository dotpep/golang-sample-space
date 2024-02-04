# Custom module in Go - step by step Tutors

1. Learn [go here in official site](https://go.dev/learn/)
2. Install Golang [click here](https://go.dev/dl/)
3. Verify Golang installed (in command prompt):

```bash
go version
```

- Confirm that the command prints the installed version of Go. Must return:

```bash must-return
go version go1.21.6 windows/amd64
```

- if not, add Golang to path environment variables if you see Golang folder

> [official tutors](https://go.dev/doc/tutorial/create-module)

## Table Of Content

- Create a Go module
- Call your code from another module
- Return and handle an error
- Return a random greeting
- Return greetings for multiple people
- Add a test
- Compile and install the application

## Create a Go module

```bash
$ mkdir greetings
$ cd greetings
$ go mod init example.com/greetings
go: creating new go.mod: module example.com/greetings
$ code greetings.go
```

```go
package greetings

import "fmt"

// Hello returns a greeting for the given name.
// Hello (function name), name (string parameter), string (return type)
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    // var message == same to message := (is dynamic variable and initialized variable)
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

```go pseudocode
function Name(param_name input type) output return type {
    variable_name :dynamicly declare (like var) type= fmt.Sprintf("Hi, %v. Welcome!", replace this %v==to_that name like $"{}" or f"{}" string formatter)
    output return message
}
```

## Call your code from another module

1. `mkdir greetings`
2. `cd greetings`
3. `go mod init example.com/greetings`
4. `code greetings.go` greetings dir

```go
package greetings

import "fmt"

func Hello(name string, surname string) string {
    message := fmt.Sprintf("Hello, %v %v. Welcome!", name, surname)
    return message
}
```

1. `cd ..`
2. `mkdir hello`
3. `go mod init example.com/hello`
4. `code hello.go` hello dir

```go
package main

import (
    "fmt"
    "example.com/greetings"
)

func main() {
    message := greetings.Hello("Gladys", "Alex")
    fmt.Println(message)
}
```

1. Declare example.com as local and on your file system module `edit` and install it `tidy` - [Edit the example.com/hello module to use your local example.com/greetings module.](https://go.dev/doc/tutorial/call-module-code):
2. `go mod edit -replace example.com/greetings=../greetings` hello dir
3. `go mod tidy`
4. To work with multiple modules [gopls workspace docs here](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md):
5. `cd ..` base dir
6. `go work init`
7. `go work use .\greetings\ .\hello\`
8. `cd .\hello`
9. `go run .`

## Return and handle an error

- greetings.go file

```go
package greetings

import (
    "errors"
    "fmt"
)

func Hello(name string, surname string) (string, error) {
    // If no name was given, return an error with message
    if name == "" || surname == "" {
        return "", errors.New("empty name or surname")
    }
    // If a name was received, return a value that  embeds the name
    // in a greating message.
    message := fmt.Sprintf("Hello, %v %v. Welcome!", name, surname)
    return message, nil // nil (meaning no error) successful return
}
```

- hello.go file

```go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request aa greeting message.
    message, err := greetings.Hello("Gladys", "")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}
```

- returns:

```bash
$ go run .
greetings: empty name
exit status 1
```

## Return a random greeting

- greetings.go file

```go
package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string, surname string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" || surname == "" {
        return "", errors.New("empty name or surname")
    }

    // Create a message using random format
    message := fmt.Sprintf(randomFormat(), name, surname)
    return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message format (slice is like an array)
    formats := []string{
        "Hi, %v %v. Welcome!",
        "Great to see you, %v %v!",
        "Hail, %v %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}
```

- hello.go file

```go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request aa greeting message.
    message, err := greetings.Hello("Gladys", "Alex")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}
```

- returns:

```bash
$ go run .
Great to see you, Gladys!

$ go run .
Hi, Gladys. Welcome!

$ go run .
Hail, Gladys! Well met!
```

## Return greetings for multiple people

- greetings.go file

```go
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
        return "", errors.New("empty name or surname")
    }

    // Create a message using random format
    message := fmt.Sprintf(randomFormat(), name)
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
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}
```

- hello.go file

```go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

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
    fmt.Println(message2)
}
```

- returns:

```bash
$ go run .
map[Darrin:Hail, Darrin! Well met! Gladys:Hi, Gladys. Welcome! Samantha:Hail, Samantha! Well met!]
```

## Add a test

1. Ending a file's name with `_test.go` tells the go test command that this file contains test functions.
2. `go test` write it in greetigs dir and command prompt to run go test
3. The go test command executes test functions (whose names begin with `Test`) in test files (whose names end with `_test.go`). You can add the `-v flag` to get verbose output that lists all of the tests and their results.

- greetings_test.go file in greetings dir

```go
package greetings

import (
    "regexp"
    "testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b` + name + `\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
```

- returns:

```bash
$ go test
PASS
ok      example.com/greetings   0.364s

$ go test -v
=== RUN   TestHelloName
--- PASS: TestHelloName (0.00s)
=== RUN   TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)
PASS
ok      example.com/greetings   0.372s
```

- check fail test:
- greetings.go file in greetings dir
- replace `message := fmt.Sprintf(randomFormat(), name)` to `message := fmt.Sprintf(randomFormat())`

```go
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // Create a message using random format
    //message := fmt.Sprintf(randomFormat(), name)
    message := fmt.Sprintf(randomFormat())
    return message, nil
}
```

- returns:

```bash
$ go test -v
=== RUN   TestHelloName
    greetings_test.go:15: Hello("Gladys") = "Hi, %!v(MISSING). Welcome!", <nil>, want match for `\bGladys\b`, nil
--- FAIL: TestHelloName (0.00s)
=== RUN   TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)
FAIL
exit status 1
FAIL    example.com/greetings   0.139s
```

## Compile and install the application

- The [go build command](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies) compiles the packages, along with their dependencies, but it doesn't install the results.
- The [go install command](https://go.dev/ref/mod#go-install) compiles and installs the packages.

- Compile code into an executable:
- `go build` hello dir
- Run exe file:
- `.\hello.exe`

- Install executable so you can run without specifying its path (add it to Go install path):

```bash
$ go list -f '{{.Target}}'
C:\Users\pepel\go\bin\hello.exe
```

- `C:\Users\pepel\go\bin\hello.exe` set it to environment variable on your system:

```bash
$ set PATH=%PATH%;C:\path\to\your\install\directory
```

- or into GOBIN variable:

```bash
$ go env -w GOBIN=C:\path\to\your\bin
```

- Once you've added `go list -f '{{.Target}}'` returns dir to env path, run the `go install` command to compile and install the package.

```bash
$ go install
```

- result (in home dir):

```bash
$ ~  hello
1. Hello, Gladys! Well met!
2. Hello, Samantha! Well met!
3. Hello, Darrin! Well met!
4. Hi, Alex. Welcome!
5. Hello, Matthew! Well met!
```
