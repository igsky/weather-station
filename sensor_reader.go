package main

// SensorReader is an interface to read data from sensors
type SensorReader interface {
	Read() (SensorReading, error)
}
