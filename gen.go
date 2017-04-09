//Generates database and sample datapoints for later use
package main

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func logFatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func startup(count int) {
	db, err := sql.Open("sqlite3", "main.db")
	logFatalErr(err)
	defer db.Close()

	err = db.Ping()
	logFatalErr(err)

	_, err = db.Exec(`CREATE TABLE timeline (id INTEGER NOT NULL PRIMARY KEY,
						   					value INTEGER NOT NULL,
						   					timestamp DATETIME)`)
	logErr(err)

	insertStmt, err := db.Prepare(`INSERT INTO timeline 
								   (value, timestamp) VALUES (?, ?);`)
	logFatalErr(err)

	span := time.Duration(5) * time.Second
	now := time.Now()

	for i := 0; i < count; i++ {
		insertStmt.Exec(rand.Intn(100), now)
		now = now.Add(span)
	}

}
