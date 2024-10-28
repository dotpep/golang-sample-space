package mysqldb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
)

// if you come from OOP PLs
// you can thing that things as

// Class
// that actually implements our Contract Interface
// in main.go file called IDBContract interface
// that needs Close(), InsertUser(), and SelectSingleUser()
type Mysql struct {
	db *sql.DB
}

// Constructor of Mysql class
// create new instance of Mysql class
func New(dbUser, dbPassword, dbHost, dbPort, dbName string) (*Mysql, error) {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Println("mysqldb connection failure: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("mysqldb ping failure: %v", err)
		return nil, err
	}

	return &Mysql{db: db}, nil
}

// Method of Mysql class
func (this Mysql) Close() {
	err := this.db.Close()
	if err != nil {
		log.Println("mysqldb close failure: %v", err)
	}
}

// Method of Mysql class
func (this Mysql) InsertUser(userName string) error {
	this.db.Exec("INSERT...")

	return nil
}

// Method of Mysql class
func (this Mysql) SelectSingleUser(userName string) (usr string, err error) {
	this.db.Exec("SELECT...")

	return "user", nil
}
