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

func stripCharsAndTitle(s string, h string) string {
	for i := strings.Index(s, h); i > 0; i = strings.Index(s, h) {
		s = s[:i] + "" + strings.Title(s[i+1:])
	}

	return s
}

func formatText(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Title(s)
	s = stripCharsAndTitle(s, "_")
	s = stripCharsAndTitle(s, "-")

	return s
}

func main() {
	file, _ := os.Open("test.csv")

	reader := csv.NewReader(file)
	header, _ := reader.Read()

	structString := "type YourCSV struct { \n"
	_ = structString

	colSlice := make([]string, 0)
	maxChars := 0

	for _, h := range header {
		t := strings.ToTitle(strings.TrimSpace(h))
		colSlice = append(colSlice, t)
		// colSlice = append(colSlice, fmt.Sprintf("\t%s string `csv:\"%s\"` \n", t, t))

		if len(t) > maxChars {
			maxChars = len(t)
		}
	}

	for _, h := range colSlice {
		if len(h) == maxChars {
			structString += fmt.Sprintf("\t%s string `csv:\"%s\"` \n", h, h)
		} else {
			padding := ""

			for i := 0; i < maxChars-len(h); i++ {
				padding += " "
			}

			structString += fmt.Sprintf("\t%s %sstring `csv:\"%s\"` \n", h, padding, h)
		}
	}

	structString += "}"

	fmt.Println(structString)
}
