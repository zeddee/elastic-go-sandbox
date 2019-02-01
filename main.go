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
	getLoop()
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

func getThis(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", os.Getenv("ELASTIC_BASEURL")+"/"+url, nil)
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
