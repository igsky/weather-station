package main

import (
	"gobot.io/x/gobot/drivers/i2c"
)

type SensorReaderBME280 struct {
	driver *i2c.BME280Driver
}

func (s *SensorReaderBME280) Read() (SensorReading, error) {
	var err error
	r := SensorReading{}

	r.Temperature, err = s.driver.Temperature()
	if err != nil {
		return r, err
	}
	r.Humidity, err = s.driver.Humidity()
	if err != nil {
		return r, err
	}
	r.Pressure, err = s.driver.Pressure()

	return r, err
}

func NewSensorReaderBME280(d *i2c.BME280Driver) *SensorReaderBME280 {
	return &SensorReaderBME280{driver: d}
}
