package main

import (
	"fmt"
	"os"
	"encoding/json"
)

func main() {
	// load raw data
	raw, err := Load()
	if err != nil {
		fmt.Println("load and unmarshal error: ", err)
		os.Exit(1)
	}

	// process & sort
	v, err := json.MarshalIndent(Sort(Process(raw)), "", "  ")
	if err != nil {
		fmt.Println("record marshal error: ", err)
		os.Exit(1)
	}

	// print by indent
	fmt.Println(string(v))
}
