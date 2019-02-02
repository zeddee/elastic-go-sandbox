package common

import (
	"fmt"
	"io/ioutil"
	"os"
)

// LoadElasticJSON wraps loading a file + unmarshals a json response
func LoadElasticJSON(filepath string) (jsonStr string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("could not open file: %s", err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("could not read from file: %s", err)
	}

	return string(bytes), nil
}
