package main

type SensorReading struct {
	Temperature float32
	Humidity    float32
	Pressure    float32
}

// SensorReader is an interface to read data from sensors
type SensorReader interface {
	Read() (SensorReading, error)
}
