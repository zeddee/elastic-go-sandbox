package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/zeddee/elastic-go-sandbox/exploringdata"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// basicoperations.ExampleOne()
	// basicoperations.ExampleTwo()
	// basicoperations.ExampleThree()
	// basicoperations.ExampleFour()
	// a.getLoop()

	exploringdata.Example()
}
