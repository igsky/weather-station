package main

import (
	"fmt"
	"log"
)

func logFatalErr(err error) {
	if err != nil {
		log.Fatal(err)
		fmt.pr
	}
}

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
