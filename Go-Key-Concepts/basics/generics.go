package main

import (
	"errors"
	"fmt"
	"time"
)

// Generics

// if we need that function to work with another slice data type
// create two function that identical but with changes in slice data type
func splitIntSlice(s []int) ([]int, []int) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func splitStringSlice(s []string) ([]string, []string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// or use `interface{}` as data type that means anything
// but we must check and process data type casting
func splitAnyInterfaceSlice(s []interface{}) ([]interface{}, []interface{}) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// better DRY way is using Generics type of `any`
func splitAnyGenericSlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// Generics Example 1
type fromToEmail struct {
	senderEmail    string
	recipientEmail string
}

type emailG struct {
	message string
	fromToEmail
}

type payment struct {
	amount int
	fromToEmail
}

func getLast[T any](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	lastIndex := len(slice) - 1
	return slice[lastIndex]
}

func testGeneric1[T any](s []T, desc string) {
	last := getLast(s)
	fmt.Printf("Getting last %v from slice of length: %v\n", desc, len(s))
	for i, v := range s {
		fmt.Printf("Item #%v: %v\n", i+1, v)
	}
	fmt.Printf("Last item in list: %v\n", last)
	fmt.Println("==================")
}

// Generic Constraints

// sometimes you need logic in your generic function to know something about the types it operates on.
// in first example we used interface built-in constraint `any`
// that didn't need to know anything about types in slice.
// we can create our custom constraints like 'any' built-in

// Creating Custom Constraint

type stringer interface {
	String() string
}

func concat[T stringer](strs []T) string {
	var result string

	for _, val := range strs {
		// this is where the `.String()` method
		// is used. That's why we need a more specific
		// constraint insead of `any` built-in constraint
		result += val.String()
	}

	return result
}

// Generics Example 2
type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

const (
	monthlySubsCost float64 = 25.00
	yearlySubsCost  float64 = 250.00
	costPerEmail    float64 = 0.03
)

func chargeForLineItem[T lineItem](
	newItem T,
	oldItems []T,
	balance float64,
) ([]T, float64, error) {
	newBalance := balance - newItem.GetCost()
	if newBalance < 0.0 {
		return nil, 0.0, errors.New("insufficient funds")
	}

	oldItems = append(oldItems, newItem)

	return oldItems, newBalance, nil
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return monthlySubsCost
	}
	if s.interval == "yearly" {
		return yearlySubsCost
	}
	return 0.0
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	return float64(otup.numEmailsAllowed) * costPerEmail
}

func testGeneric2[T lineItem](newItem T, oldItems []T, balance float64) {
	fmt.Printf("Charging customer for a '%s', current balance is %v...\n", newItem.GetName(), balance)
	newItems, newBalance, err := chargeForLineItem(newItem, oldItems, balance)

	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		return
	}

	fmt.Printf("New balance is: %v. Total number of line items is now %v\n", newBalance, len(newItems))
	fmt.Println("==================")
}

// Interface Type Lists

// we can list a bunch of types to get a new interface/constraint

// Ordered is type constraint that mathces any ordered type.
// An ordered type is one that supports the <, <=, > and >= comparison operators.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Parametric Constraints

// The store interface represents a store that sells products.
// It takes a type parameter P that represents the type of product.
type store[P product] interface {
	Sell(P)
}

type product interface {
	Price() float64
	Name() string
}

type book struct {
	title  string
	author string
	price  float64
}

func (b book) Price() float64 {
	return b.price
}

func (b book) Name() string {
	return fmt.Sprintf("%s by %s", b.title, b.author)
}

type toy struct {
	name  string
	price float64
}

func (t toy) Price() float64 {
	return t.price
}

func (t toy) Name() string {
	return t.name
}

// The bookStore struct represents a store that sells book
type bookStore struct {
	bookSold []book
}

// Sell adds a book to the bookStore's inventory.
func (bs *bookStore) Sell(b book) {
	bs.bookSold = append(bs.bookSold, b)
}

// The toyStore struct represents a store that sells toys.
type toyStore struct {
	toysSold []toy
}

// Sell adds a toy to the toyStore's inventory.
func (ts *toyStore) Sell(t toy) {
	ts.toysSold = append(ts.toysSold, t)
}

// sellProducts takes a store and a slice of products and sells
// each product one by one
func sellProducts[P product](s store[P], products []P) {
	for _, p := range products {
		s.Sell(p)
	}
}

// Generics Example 3
type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type userG struct {
	UserEmail string
}

type org struct {
	Admin userG
	Name  string
}

type userBiller struct {
	Plan string
}

type orgBiller struct {
	Plan string
}

func (ub userBiller) Charge(u userG) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (ub userBiller) Name() string {
	return fmt.Sprintf("%s user biller", ub.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (u userG) GetBillingEmail() string {
	return u.UserEmail
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

func testBillerGeneric3[C customer](b biller[C], c C) {
	fmt.Printf("Using '%s' to create a bill for '%s'\n", b.Name(), c.GetBillingEmail())
	bill := b.Charge(c)
	fmt.Printf("Bill created for %v dollars\n", bill.Amount)
	fmt.Println(" --- ")
}

// Naming Generic types

// 'T' is common convention name rather that own custom name
// `T any`

func splitMyAnyTypeSlice[MyAnyType any](s []MyAnyType) ([]MyAnyType, []MyAnyType) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func main() {
	//var x struct{}
	//var x interface{}
	var x any
	x = 'a'
	fmt.Println(x)
	fmt.Printf("x data type: %T\n", x)

	var yByte byte
	fmt.Printf("yByte data type: %T\n", yByte)

	var yRune rune
	fmt.Printf("yRune data type: %T\n", yRune)

	fmt.Println("---")

	firstInts, secondInts := splitAnyGenericSlice([]int{0, 1, 2, 3})
	fmt.Println(firstInts, secondInts)

	firstStrings, secondStrings := splitAnyGenericSlice[string]([]string{"text", "str", "string", "message"})
	fmt.Println(firstStrings, secondStrings)

	fmt.Println("---")

	// Generics Example 1
	janeToMargo := fromToEmail{
		"jane@example.com",
		"margo@example.com",
	}
	janeToSally := fromToEmail{
		"jane@example.com",
		"sally@example.com",
	}

	testGeneric1([]emailG{}, "emailG")
	testGeneric1([]emailG{
		{
			"Hi Margo",
			janeToMargo,
		},
		{
			"Hey Margo I really wanna chat",
			janeToMargo,
		},
		{
			"ANSWER ME",
			janeToMargo,
		},
	}, "emailG")
	testGeneric1([]payment{
		{
			5,
			janeToSally,
		},
		{
			25,
			janeToMargo,
		},
		{
			1,
			janeToSally,
		},
		{
			16,
			janeToMargo,
		},
	}, "payment")

	fmt.Println("---")

	// Generics Example 2
	testGeneric2(subscription{
		userEmail: "john@example.com",
		startDate: time.Now().UTC(),
		interval:  "yearly",
	},
		[]subscription{},
		1000.00,
	)
	testGeneric2(subscription{
		userEmail: "jane@example.com",
		startDate: time.Now().UTC(),
		interval:  "monthly",
	},
		[]subscription{
			{
				userEmail: "jane@example.com",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7),
				interval:  "monthly",
			},
			{
				userEmail: "jane@example.com",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7 * 52 * 2),
				interval:  "yearly",
			},
		},
		686.20,
	)
	testGeneric2(oneTimeUsagePlan{
		userEmail:        "dillon@example.com",
		numEmailsAllowed: 5000,
	},
		[]oneTimeUsagePlan{},
		756.20,
	)
	testGeneric2(oneTimeUsagePlan{
		userEmail:        "dalton@example.com",
		numEmailsAllowed: 100000,
	},
		[]oneTimeUsagePlan{
			{
				userEmail:        "dalton@example.com",
				numEmailsAllowed: 34200,
			},
		},
		32.20,
	)

	fmt.Println("---")

	// Parametric Constraints
	bs := bookStore{
		bookSold: []book{},
	}

	// by passing 'book' as a type parameter,
	// we can use the sellProducts function to sell books in a bookStore
	sellProducts[book](&bs, []book{
		{
			title:  "The Hobbit",
			author: "J.R.R. Tolkien",
			price:  10.0,
		},
		{
			title:  "The Lord of the Rings",
			author: "J.R.R. Tolkien",
			price:  20.0,
		},
	})
	fmt.Println(bs.bookSold)

	// same thing for toys
	ts := toyStore{
		toysSold: []toy{},
	}
	sellProducts[toy](&ts, []toy{
		{
			name:  "Lego",
			price: 10.0,
		},
		{
			name:  "Barbie",
			price: 20.0,
		},
	})
	fmt.Println(ts.toysSold)

	fmt.Println("---")

	// Generics Example 3
	testBillerGeneric3[userG](
		userBiller{Plan: "basic"},
		userG{UserEmail: "joe@example.com"},
	)
	testBillerGeneric3[userG](
		userBiller{Plan: "basic"},
		userG{UserEmail: "samuel.boggs@example.com"},
	)
	testBillerGeneric3[userG](
		userBiller{Plan: "pro"},
		userG{UserEmail: "jade.row@example.com"},
	)
	testBillerGeneric3[org](
		orgBiller{Plan: "basic"},
		org{Admin: userG{UserEmail: "challis.rane@example.com"}},
	)
	testBillerGeneric3[org](
		orgBiller{Plan: "pro"},
		org{Admin: userG{UserEmail: "challis.rane@example.com"}},
	)
}
