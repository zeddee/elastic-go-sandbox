package main

import (
	"log"

	"github.com/joho/godotenv"
	a "github.com/zeddee/elastic-go-sandbox/apicalls"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	//exampleOne()
	exampleTwo()

	//a.getLoop()
}

// Simple Get, Put, and Delete
func exampleOne() {
	a.Get("_cat/health?v")  // get cluster health
	a.Get("_cat/nodes?v")   // get node list
	a.Get("_cat/indices?v") // get list of indexes/indices

	a.Put("customer?pretty", "") // Add an index
	a.Get("_cat/indices?v")      // get list of indexes/indices

	// Add document to index, with payload
	payload := `
			{
				"name": "Jane Doe"
			}
			`
	a.Put("customer/_doc/1?pretty", payload)

	a.Get("customer") // get list of documents at customer

	a.Delete("customer")    // cleanup by deleting customer index
	a.Get("_cat/indices?v") // get list of indexes/indices
}

// Updating existing indicies/docs, and Putting without explicit ID
func exampleTwo() {
	payload1 := `
				{
				"name": "John Doe"
				}
				`
	a.Put("customer/_doc/1?pretty", payload1)
	a.Get("customer")

	payload2 := `
				{
					"name": "Jane Doe"
				}
	`
	a.Put("customer/_doc/1?pretty", payload2)
	a.Get("customer")

	// Adding doc without explicit ID
	// MUST use POST method for IDless calls
	a.Post("customer/_doc?pretty", payload2)
	a.Get("customer")

	a.Delete("customer")    // cleanup by deleting customer index
	a.Get("_cat/indices?v") // get list of indexes/indices
}
