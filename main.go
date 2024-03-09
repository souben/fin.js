package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/souben/fin.js/processing"
)

func main() {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")
	dec := json.NewDecoder(os.Stdin)
	var input map[string]interface{}
	if err := dec.Decode(&input); err != nil {
		fmt.Fprintf(os.Stderr, "error : %v\n", err)
	}
	result := processing.TransformAndProcess(input)
	enc.Encode(result)
}
