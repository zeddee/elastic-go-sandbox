package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	getOne("_cat/health?v")  // get cluster health
	getOne("_cat/nodes?v")   // get node list
	getOne("_cat/indices?v") // get list of indexes/indices

	putOne("customer?pretty", "") // Add an index
	getOne("_cat/indices?v")      // get list of indexes/indices

	// Add document to index, with payload
	payload := `
			{
				"name": "Jane Doe"
			}
			`
	putOne("customer/_doc/1?pretty", payload)

	getOne("customer") // get list of documents at customer

	deleteOne("customer") // cleanup by deleting customer index
	//getLoop()
}

func putOne(path string, data string) {
	res, err := putThis(path, data)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("PUT result: %s\n", res)
}

func getOne(path string) {
	val, err := getThis(path)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s\n", val)
}

func deleteOne(path string) {
	res, err := deleteThis(path)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("DELETE result: %s\n", res)
}

func getThis(path string) (resBody []byte, error error) {
	return newRequest(http.MethodGet, path, "")
}

func putThis(path string, data string) (resBody []byte, error error) {
	return newRequest(http.MethodPut, path, data)
}

func deleteThis(path string) (resBody []byte, error error) {
	return newRequest(http.MethodDelete, path, "")
}

func newRequest(method string, path string, data string) ([]byte, error) {
	req, err := http.NewRequest(method, os.Getenv("ELASTIC_BASEURL")+"/"+path, strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("could not make request: %s", err)
	}
	req.Header.Set("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not get valid response: %s", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read body of response: %s", err)
	}

	return body, nil
}

func getLoop() {
	for {
		fmt.Printf("Please enter the path to GET\n")
		input := readStringStdin()
		if input == "" {
			break
		}
		val, err := getThis(input)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Response: %s\n", val)
	}
}

func readStringStdin() string {
	reader := bufio.NewReader(os.Stdin)
	inputVal, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("invalid option: %v\n", err)
		return ""
	}

	output := strings.TrimSuffix(inputVal, "\n") // Important!
	return output
}
