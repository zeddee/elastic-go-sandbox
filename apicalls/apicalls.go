package apicalls

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const sep = "=======================\n"

// Put wraps a put call to the elasticsearch server REST API
func Put(path string, data string) {
	res, err := putThis(path, data)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf(sep)
	fmt.Printf("PUT result: %s\n", res)
}

// Post wraps a post call to the elasticsearch server REST API
func Post(path string, data string) {
	res, err := postThis(path, data)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf(sep)
	fmt.Printf("POST result: %s\n", res)
}

// Get wraps a get call to the elasticsearch server REST API
func Get(path string) {
	val, err := getThis(path)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf(sep)
	fmt.Printf("%s\n", val)
}

// GetWithJSONQuery wraps a GET call + request body to the elasticsearch server REST API
func GetWithJSONQuery(path string, data string) {
	res, err := getThisWithReqBody(path, data)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf(sep)
	fmt.Printf("GET+ result: %s\n", res)
}

// Delete wraps a delete call to the elasticsearch server REST API
func Delete(path string) {
	res, err := deleteThis(path)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf(sep)
	fmt.Printf("DELETE result: %s\n", res)
}

func getThis(path string) (resBody []byte, error error) {
	return newRequest(http.MethodGet, path, "")
}

func getThisWithReqBody(path, data string) (resBody []byte, error error) {
	return newRequest(http.MethodGet, path, data)
}

func putThis(path string, data string) (resBody []byte, error error) {
	return newRequest(http.MethodPut, path, data)
}

func postThis(path string, data string) (resBody []byte, error error) {
	return newRequest(http.MethodPost, path, data)
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

// CLI wraps a for loop that starts getting input from stdin to parse as CLI commands
func CLI() {
	for {
		fmt.Printf(sep)
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
