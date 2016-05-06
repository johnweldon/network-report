package main

import (
	"fmt"
	"log"

	"github.com/johnweldon/network-report/lib"
)

func main() {
	checker := report.NewMockChecker()
	r, err := checker.Test("8.8.4.4")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", r)
}
