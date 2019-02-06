// Package exploringdata follows this guide:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/getting-started-explore-data.html
package exploringdata

import (
	"log"

	a "github.com/zeddee/elastic-go-sandbox/apicalls"
	"github.com/zeddee/elastic-go-sandbox/common"
)

// LoadData loads the prescribed dataset
func LoadData() {
	payload, err := common.LoadElasticJSON("data/accounts.json")
	if err != nil {
		log.Println(err)
	}

	a.Post("bank/_doc/_bulk?pretty&refresh", payload)
	a.Get("_cat/indices?v")
}

// Cleanup deletes the bank index, cleaning up the elastic instance for this tutorial section
func Cleanup() {
	a.Delete("bank")        // cleanup by deleting customer index
	a.Get("_cat/indices?v") // get list of indexes/indices
}
