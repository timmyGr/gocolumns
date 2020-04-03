package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

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
	argsWithProg := os.Args

	file, err := os.Open(argsWithProg[1])

	if err != nil {
		log.Panic("CSV not found.")
	}

	reader := csv.NewReader(file)
	header, _ := reader.Read()

	structName := formatText(strings.Replace(argsWithProg[1], ".csv", "", -1))
	structString := fmt.Sprintf("type %s struct { \n", structName)

	maxChars := 0

	for i, h := range header {
		header[i] = formatText(h)

		if len(header[i]) > maxChars {
			maxChars = len(header[i])
		}
	}

	for _, h := range header {
		if len(h) == maxChars {
			structString += fmt.Sprintf("\t%s string `csv:\"%s\"` \n", h, h)
		} else {
			padding := ""

			for i := 0; i < maxChars-len(h)-1; i++ {
				padding += " "
			}

			structString += fmt.Sprintf("\t%s %s string `csv:\"%s\"` \n", h, padding, h)
		}
	}

	structString += "}"

	fmt.Println(structString)
}
