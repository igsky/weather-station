package main

import (
	"database/sql"
	"flag"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

type sensorReading struct {
	Temperature, Humidity, Pressure float32
}

func readSensors(bme *i2c.BME280Driver) sensorReading {
	t, err := bme.Temperature()
	h, err := bme.Humidity()
	p, err := bme.Pressure()
	logFatalErr(err)
	return sensorReading{t, h, p}
}

func createTable(db *sql.DB) {
	err := db.Ping()
	logFatalErr(err)

	_, err = db.Exec(`CREATE TABLE timeline (id INTEGER NOT NULL PRIMARY KEY,
											 temperature REAL NOT NULL,
						   					 humidity REAL NOT NULL,
											 pressure REAL NOT NULL,
						   					 timestamp DATETIME)`)
	logErr(err)
}

func main() {
	var genCount = flag.Int("s", 60, "Read period in seconds")

	var r = raspi.NewAdaptor()
	var sensor = i2c.NewBME280Driver(
		r,
		i2c.WithBus(1),
		i2c.WithAddress(0x76),
	)

	db, err := sql.Open("sqlite3", "main.db")
	logFatalErr(err)
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO timeline 
							(temperature, humidity, temperature, timestamp)
							VALUES (?, ?, ?, date('now'));`)
	logFatalErr(err)
	defer stmt.Close()

	writeReading := func(reading sensorReading) {
		_, err := stmt.Exec(
			reading.Temperature,
			reading.Humidity,
			reading.Pressure)
		logFatalErr(err)
	}

	flag.Parse()

	work := func() {
		gobot.Every(time.Duration(*genCount)*time.Minute, func() {
			writeReading(readSensors(sensor))
		})
	}

	robot := gobot.NewRobot(
		"bot",
		[]gobot.Connection{r},
		[]gobot.Device{sensor},
		work,
	)

	go robot.Start()
	startServer()
}
