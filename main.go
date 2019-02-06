package main

import (
	"log"

	"github.com/joho/godotenv"
	a "github.com/zeddee/elastic-go-sandbox/apicalls"
	"github.com/zeddee/elastic-go-sandbox/common"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// basicoperations.exampleOne()
	// basicoperations.exampleTwo()
	// basicoperations.exampleThree()
	// basicoperations.exampleFour()
	loadElasticData()
	// a.getLoop()
	defer cleanup()
}

// https://www.elastic.co/guide/en/elasticsearch/reference/current/getting-started-explore-data.html
func loadElasticData() {
	payload, err := common.LoadElasticJSON("data/accounts.json")
	if err != nil {
		log.Println(err)
	}

	a.Post("bank/_doc/_bulk?pretty&refresh", payload)
	a.Get("_cat/indices?v")
}

func cleanup() {
	a.Delete("bank")        // cleanup by deleting customer index
	a.Get("_cat/indices?v") // get list of indexes/indices
}
