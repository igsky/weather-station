package main

type SensorReading struct {
	Temperature float32
	Humidity    float32
	Pressure    float32
}

type SensorReader interface {
	Read() (SensorReading, error)
}
