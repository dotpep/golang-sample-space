package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dotpep/interfaces/hello"
	_ "github.com/dotpep/interfaces/mysqldb"
	"github.com/dotpep/interfaces/postgresdb"
	"github.com/joho/godotenv"
)

type IDBContract interface {
	Close()
	InsertUser(userName string) error
	SelectSingleUser(userName string) (usr string, err error)
}

type Application struct {
	DB IDBContract
}

func (this Application) Run() {
	userName := "user2"

	err := this.DB.InsertUser(userName)
	if err != nil {
		log.Println(err)
	}

	user, err := this.DB.SelectSingleUser(userName)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("inserted and selected single user: ", user)
}

func NewApplication(db IDBContract) *Application {
	return &Application{DB: db}
}

func main() {
	name := "Alex"
	msg := hello.SayHello(name)
	fmt.Println(msg)

	fmt.Println("---")

	godotenv.Load(".env")

	dbURL := os.Getenv("DB_URL")
	postgresDB, err := postgresdb.New(dbURL)

	if err != nil {
		log.Fatal("failed to initiate dbase connection: %v", err)
	}

	defer postgresDB.Close()

	app := NewApplication(postgresDB)

	app.Run()

	fmt.Println("---")

	//dbUser := os.Getenv("DB_USER")
	//dbPassword := os.Getenv("DB_PASSWORD")
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	//dbName := os.Getenv("DB_NAME")

	//mysqlDB, err := mysqldb.New(dbUser, dbPassword, dbHost, dbPort, dbName)
}
