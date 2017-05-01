package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(name string) {
	db, err := sql.Open("sqlite3", name)
	logFatalErr(err)
	defer db.Close()

	err = db.Ping()
	logFatalErr(err)
}

func CreateTable() {
	_, err := db.Exec(
		`CREATE TABLE timeline (
			id INTEGER NOT NULL PRIMARY KEY,
			temperature REAL,
			humidity REAL,
			pressure REAL,
			timestamp DATETIME)`)
	logErr(err)
}

func CreateWriteStmt() (*sql.Stmt, func(r sensorReading)) {
	stmt, err := db.Prepare(`INSERT INTO timeline 
							(temperature, humidity, pressure, timestamp)
							VALUES (?, ?, ?, datetime('now'));`)
	logFatalErr(err)

	writeReading := func(reading sensorReading) {
		_, err := stmt.Exec(
			reading.Temperature,
			reading.Humidity,
			reading.Pressure)
		logFatalErr(err)
	}

	return stmt, writeReading
}
