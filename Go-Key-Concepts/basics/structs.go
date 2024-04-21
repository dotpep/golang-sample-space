package main

import (
	"fmt"
)

// structs is just custom type that contains other types
type car struct {
	Make   string
	Model  string
	Height int
	Width  int
}

// Nested
type carNested struct {
	Make   string
	Model  string
	Height int
	Width  int
	// Nested Struct
	FrontWheel Wheel
	BackWheel  Wheel
}

type Wheel struct {
	Radius   int
	Material string
}

// Anonymous
type carAnonNested struct {
	Make   string
	Model  string
	Height int
	Width  int
	// Anonymous Struct in Nested struct
	Wheel struct {
		Radius   int
		Material string
	}
}

// Embedded
type carEmbedded struct {
	make  string
	model string
}

// Embedded Struct type is like inheriting all fields/types of other provided struct
type truck struct {
	// "car" is embedded, so defination of a
	// "truck" now also additionally contains all
	// of the fields of car struct
	// if we provide name "car carEmbedded"
	// then it cannot provide fields in top-level
	// and we have to truck.car.fields instead truck.fields
	carEmbedded
	bedSize int
}

// Example 1
type messageToSend struct {
	message     string
	phoneNumber int
}

func printMessage(msg messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", msg.message, msg.phoneNumber)
	fmt.Println("==================")
}

// Nested Example 2
type messageSendToUser struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func printUserMessage(msg messageSendToUser) {
	fmt.Printf(
		"Sending message: '%s', \nfrom: %s, (%v) \nto: %s, (%v)\n",
		msg.message,
		msg.sender.name, msg.sender.number,
		msg.recipient.name, msg.recipient.number,
	)

	if canUserSendMessage(msg) {
		fmt.Println("...sent!")
	} else {
		fmt.Println("...cat't send message!")
	}

	fmt.Println("==================")
}

// Validator
func canUserSendMessage(msgToSend messageSendToUser) bool {
	sender := msgToSend.sender
	recipient := msgToSend.recipient

	if sender.name == "" && sender.number == 0 {
		return false
	}
	if recipient.name == "" && recipient.number == 0 {
		return false
	}

	return true
}

// Embedded Example 3
type sender struct {
	user
	rateLimit int
}

func printUserSenderRateLimit(sndr sender) {
	fmt.Println("Sender name: ", sndr.name)
	fmt.Println("Sender number: ", sndr.number)
	fmt.Println("Sender rate limit: ", sndr.rateLimit)
	fmt.Println("==================")
}

// Struct Methods
type rectangle struct {
	width  int
	height int
}

// area has a receiver of (rect rectangle)
// "area" is name of method on "rectangle" struct
// just func but with special parameter before name of func
// that special parameter named "rect" type "rectangle" comes into function
func (rect rectangle) area() int {
	return rect.width * rect.height
}

// Struct Methods Example 4
type authenticationInfo struct {
	username string
	password string
}

func (authInfo authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: Basic %s:%s",
		authInfo.username, authInfo.password,
	)
}

func authenticateQueryMessage(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("==================")
}

// Test Cases separated by Println("---")
func main() {
	printMessage(messageToSend{
		phoneNumber: 81231230010,
		message:     "How are you doing today, hope that all is fine.",
	})
	printMessage(messageToSend{
		phoneNumber: 11231230010,
		message:     "Today I am learning new programming language",
	})

	fmt.Println("---")

	// Instantiate new instance of empty car Struct called myCar
	myCar := carNested{}
	myCar.FrontWheel.Radius = 5
	fmt.Printf("my car info: %+v\n", myCar)
	fmt.Printf("front wheel radius: %v\n", myCar.FrontWheel.Radius)

	fmt.Println("---")

	userAlex := user{name: "Alex", number: 81231230010}
	userJohn := user{name: "John", number: 11231230010}
	userAshley := user{name: "Ashley"}
	defaultUser := user{}

	printUserMessage(messageSendToUser{
		message:   "You have appointment today",
		sender:    userAlex,
		recipient: userJohn,
	})
	printUserMessage(messageSendToUser{
		message:   "Thank for reminding me",
		sender:    userJohn,
		recipient: userAlex,
	})
	printUserMessage(messageSendToUser{
		message: "I'll see you there at 8 pm tonight",
		sender:  userAlex,
		//recipient: userJohn,
	})
	printUserMessage(messageSendToUser{
		message:   "Hello Ash, are you busy tonight",
		sender:    defaultUser,
		recipient: userAshley,
	})

	fmt.Println("---")

	// Anonymous Structs
	anonCar := struct {
		Make  string
		Model string
	}{
		Make:  "Tesla",
		Model: "Model 3",
	}
	fmt.Println("anonymous car struct instance: ", anonCar)

	fmt.Println("---")

	// Embedded Structs
	lanesTruck := truck{
		bedSize: 10,
		carEmbedded: carEmbedded{
			make:  "Toyota",
			model: "Camry",
		}}
	fmt.Println("truck bedSize: ", lanesTruck.bedSize)

	// embedded fields promoted to top-level
	// instead of lanesTruck.car.make
	fmt.Println("truck make:", lanesTruck.make)
	fmt.Println("truck model:", lanesTruck.model)

	fmt.Println("---")

	// Embedded Structs Example
	printUserSenderRateLimit(sender{
		rateLimit: 1000,
		user: user{
			name:   "Daniel",
			number: 14561230044,
		},
	})
	printUserSenderRateLimit(sender{
		rateLimit: 200,
		user:      userAlex,
	})

	fmt.Println("---")

	// Struct Methods
	rect1 := rectangle{
		width:  5,
		height: 10,
	}
	fmt.Printf("rectangle %+v \n", rect1)
	fmt.Println("area of rectangle: ", rect1.area())

	fmt.Println("---")

	// Struct Methods Example
	authAlex := authenticationInfo{username: "Alex123", password: "artur123"}
	authJohn := authenticationInfo{username: "JohnDoe", password: "doesecret"}
	authAshley := authenticationInfo{username: "Ash", password: "qwerty2024"}

	authenticateQueryMessage(authAlex)
	authenticateQueryMessage(authJohn)
	authenticateQueryMessage(authAshley)
}
