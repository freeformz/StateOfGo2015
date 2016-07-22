package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func swapRows(r []string) {
	r[len(r)-2], r[len(r)-1] = r[len(r)-1], r[len(r)-2]
}
func clipEmail(r []string) []string {
	return r[0 : len(r)-1]
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("filter.go <input> <output>")
	}
	i := os.Args[1]
	o := os.Args[2]
	input, err := os.Open(i)
	if err != nil {
		panic(err)
	}
	output, err := os.Create(o)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(input)
	f := csv.NewWriter(output)
	var hl int //header length
	for row, err := r.Read(); err == nil; row, err = r.Read() {
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		swapRows(row)
		row = clipEmail(row)
		if hl == 0 {
			fmt.Println(row)
			hl = len(row)
		}
		if len(row) != hl {
			fmt.Printf("Row length (%d) different then header row length (%d)\n:%s\n", len(row), hl, row)
		}

		err := f.Write(row)
		if err != nil {
			panic(err)
		}
	}
	f.Flush()
}
