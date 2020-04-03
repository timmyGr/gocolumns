package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// type Client struct {
// 	Id      string `csv:"client_id"`
// 	Name    string `csv:"client_name"`
// 	Age     string `csv:"client_age"`
// 	NotUsed string `csv:"-"`
// }

func main() {
	file, _ := os.Open("test.csv")

	reader := csv.NewReader(file)
	header, _ := reader.Read()

	structString := "type YourCSV struct { \n"
	_ = structString

	colSlice := make([]string, 0)

	for _, h := range header {
		t := strings.ToTitle(strings.TrimSpace(h))
		colSlice = append(colSlice, fmt.Sprintf("\t%s string `csv:\"%s\"` \n", t, t))
	}

	// pad struct values to properly line up when printing

	structString += "}"

	fmt.Println(colSlice)
	fmt.Println(structString)
}
