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

	// basicoperations.exampleOne()
	// basicoperations.exampleTwo()
	// basicoperations.exampleThree()
	// basicoperations.exampleFour()
	// a.getLoop()

	exploringdata.LoadData()

	defer exploringdata.Cleanup()
}
