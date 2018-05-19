package main

import (
	"flag"
	"log"
	"time"

	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

var readPeriod *int

func init() {
	readPeriod = flag.Int("s", 10, "Read period in seconds")
	flag.Parse()
}

func main() {
	var r = raspi.NewAdaptor()
	var sensor = i2c.NewBME280Driver(
		r,
		i2c.WithBus(1),
		i2c.WithAddress(0x76),
	)
	sensorReader := NewSensorReaderBME280(sensor)

	go func() {
		for {
			reading, err := sensorReader.Read()
			logErr(err)

			log.Printf(
				"Temperature: %.2f, Humidity: %.2f, Pressure: %.2f",
				reading.Temperature, reading.Humidity, reading.Pressure,
			)

			time.Sleep(time.Duration(*readPeriod) * time.Second)
		}
	}()

	forever := make(chan interface{})
	<-forever
}
