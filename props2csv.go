package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	props "github.com/magiconair/properties"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: %s keyColumnName valueColumnName < some.properties\n", os.Args[0])
		os.Exit(1)
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}

	data := props.MustLoadString(string(input))

	w := csv.NewWriter(os.Stdout)

	w.Write(os.Args[1:])

	for _, k := range data.Keys() {
		v, _ := data.Get(k)
		w.Write([]string{k, v})
	}

	w.Flush()
}
