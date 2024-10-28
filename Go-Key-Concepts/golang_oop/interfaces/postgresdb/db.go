package postgresdb

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

type Postgres struct {
	db *sql.DB
}

func New(dbUrl string) (*Postgres, error) {
	var err error

	//connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=enabled", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open(dbDriver, dbUrl)
	if err != nil {
		log.Println("postgres connection failure: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("postgres ping failure: %v", err)
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (this Postgres) Close() {
	err := this.db.Close()
	if err != nil {
		log.Println("postgres close failure: %v", err)
	}
}

func (this Postgres) InsertUser(userName string) error {
	this.db.Exec("INSERT...")

	return nil
}

func (this Postgres) SelectSingleUser(userName string) (usr string, err error) {
	this.db.Exec("SELECT...")

	return "user", nil
}
