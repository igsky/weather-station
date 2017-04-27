package main

import (
	"time"
	"database/sql"
	"flag"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
	"gobot.io/x/gobot/drivers/i2c"
	_ "github.com/mattn/go-sqlite3"
)


type SensorReading struct {
	Temperature, Humidity, Pressure float32
}

func (bme i2c.BME280Driver) readSensors() SensorReading {
	return SensorReading {
		bme.Temperature(),
		bme.Humidity(),
		bme.Pressure()
	}
}

func createTable(db sql.DB) {
	err = db.Ping()
	logFatalErr(err)

	_, err = db.Exec(`CREATE TABLE timeline (id INTEGER NOT NULL PRIMARY KEY,
											 temperature REAL NOT NULL,
						   					 humidity REAL NOT NULL,
											 pressure REAL NOT NULL,
						   					 timestamp DATETIME)`)
	logErr(err)
}

func writeReading(db sql.DB, reading SensorReading) {
	_, err :=  stmt.Exec(reading.Temperature,
						 reading.Humidity,
						 reading.Pressure)
   logFatalErr(err)
}

func main() {
	var genCount = flag.Int("g", 0, "Specifies how many datapoints to generate")

	var r = raspi.NewAdaptor()
	var sensor = i2c.NewBME280Driver(
		r,
		i2c.WithBus(1),
		i2c.WithAddress(0x76)
	)

	db, err := sql.Open("sqlite3", "main.db")
	logFatalErr(err)
	defer db.Close()


	stmt, err := db.Prepare(`INSERT INTO timeline 
							(temperature, humidity, temperature, timestamp)
							VALUES (?, ? , ?, date('now'));`)
	logFatalErr(err)
	defer stmt.Close()

	flag.Parse()

	work := func() {
		gobot.Every(1*time.Minute, func() {
			writeReading(db, sensor.readSensors())
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
