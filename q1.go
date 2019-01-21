package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	inputcsv, _ := os.Open("resale-flat-prices.csv")
	csvr := csv.NewReader(inputcsv)

	q1csv, _ := os.OpenFile("q1.csv", os.O_RDWR|os.O_CREATE, 0666)
	q1wr := csv.NewWriter(q1csv)
	q1 := [][]string{}

	fmt.Printf("loop starting... ")
	for {
		row, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if row[2] == "3 ROOM" && row[7] == "ADJOINED FLAT" {
			q1 = append(q1, row)
		}
	}
	fmt.Println("loop completed")
	q1wr.WriteAll(q1)
}
