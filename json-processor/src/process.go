package main

import (
	"sort"
	"strings"
)

type Record struct {
	Label 		string 	`json:"label"`
	Sentence 	string 	`json:"sentence"`
	ContractId 	string 	`json:"contract_id"`
	Interval 	int 	`json:"interval"`
}


// Process its raw json to records
func Process(raw *Raw) (r []Record) {
	completions, data := raw.Completions, raw.Data
	contract_id := strings.TrimSuffix(data.Filename, ".html")
	r = make([]Record, 0)

	for _, complete := range completions {
		for _, res := range complete.Result {
			record := Record{
				Label: strings.Join(res.Value.HTMLLabels, ","),
				Sentence: res.Value.Text,
				ContractId: contract_id,
				Interval: res.Value.EndOffset - res.Value.StartOffset,
			}

			r = append(r, record)
		}
	}

	return
}

// Sort its records by start offset and end offset
func Sort(records []Record) (r []Record) {
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].Interval < records[j].Interval
	})

	r = records
	return
}
