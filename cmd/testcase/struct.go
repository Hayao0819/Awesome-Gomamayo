package testcase

import (
	"encoding/json"
	"os"
)

type Testcase struct {
	Gomamayo []struct {
		Word  string `json:"word"`
		Yomi  string `json:"yomi"`
		Terms int    `json:"terms"`
		Dim   int    `json:"dim"`
	} `json:"gomamayo"`
	NotGomamayo []struct {
		Word  string `json:"word"`
		Yomi  string `json:"yomi"`
		Terms int    `json:"terms"`
		Dim   int    `json:"dim"`
	} `json:"not-gomamayo"`
	Uncategorized []string `json:"uncategorized"`
}

func (t Testcase) Write(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "    ")
	if err := enc.Encode(t); err != nil {
		return err
	}
	return nil
}
