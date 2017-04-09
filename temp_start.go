package main

import (
	"flag"
)

var genCount = flag.Int("gen", 0, "Specifies how many datapoints to generate")

func main() {
	flag.Parse()

	if *genCount > 0 {
		startup(*genCount)
	}

}
