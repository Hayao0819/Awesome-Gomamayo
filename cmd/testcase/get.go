package testcase

import (
	"encoding/json"
	"os"
)

func GetTestCase(path string) (*Testcase, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tc Testcase
	if err = json.Unmarshal(data, &tc); err != nil {
		return nil, err
	}

	return &tc, nil
}
