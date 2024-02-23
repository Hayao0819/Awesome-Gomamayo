package testcase

import (
	"encoding/json"
	"os"
	"strconv"
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

func (tc *Testcase) Write(j, c, t string) error {
	var err error

	// write json
	err = tc.WriteJson(j)
	if err != nil {
		return err
	}

	// write csv
	err = tc.WriteCsv(c)
	if err != nil {
		return err
	}

	// write txt
	err = tc.WriteTxt(t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Testcase) WriteJson(path string) error {
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

func (t *Testcase) WriteCsv(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// write header
	_, err = file.WriteString("word,yomi,terms,dim\n")
	if err != nil {
		return err
	}

	for _, g := range t.Gomamayo {
		_, err = file.WriteString(g.Word + "," + g.Yomi + "," + strconv.Itoa(g.Terms) + "," + strconv.Itoa(g.Dim) + "\n")
		if err != nil {
			return err
		}
	}

	for _, n := range t.NotGomamayo {
		_, err = file.WriteString(n.Word + "," + n.Yomi + "," + strconv.Itoa(n.Terms) + "," + strconv.Itoa(n.Dim) + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Testcase) WriteTxt(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, u := range t.Gomamayo {
		_, err = file.WriteString(u.Word + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
