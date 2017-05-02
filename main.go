package main

import (
	"flag"
	"time"

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

func main() {
	var period = flag.Int("s", 60, "Read period in seconds")
	flag.Parse()

	var r = raspi.NewAdaptor()
	var sensor = i2c.NewBME280Driver(
		r,
		i2c.WithBus(1),
		i2c.WithAddress(0x76),
	)

	InitDB("main.db")
	defer db.Close()

	CreateTable()

	writerClose, writeReading := CreateWriter()
	defer writerClose()

	work := func() {
		gobot.Every(time.Duration(*period)*time.Second, func() {
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
