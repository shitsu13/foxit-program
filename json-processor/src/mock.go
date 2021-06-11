package main

import (
	"encoding/json"
	"io/ioutil"
)

type (
	Raw struct {
		Completions []Completions `json:"completions"`
		Data Data `json:"data"`
	}

	Completions struct {
		Result []Result `json:"result"`
	}

	Result struct {
		Value Value `json:"value"`
	}

	Value struct {
		HTMLLabels 	[]string 	`json:"htmllabels"`
		Text 		string 		`json:"text"`
		StartOffset int 		`json:"startOffset"`
		EndOffset 	int 		`json:"endOffset"`
	}

	Data struct {
		Filename string `json:"filename"`
	}
)

// Load return raw json to use
func Load() (r *Raw, err error) {
	file, err := ioutil.ReadFile("src\\raw.json")
	if err != nil {
        return
    }

	if err = json.Unmarshal([]byte(file), &r); err != nil {
		return
	}

	return
}