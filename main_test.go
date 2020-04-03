package main

import (
	"strings"
	"testing"
)

func TestFormatText(t *testing.T) {
	s := formatText(" this-is_ A-col umn_name ")

	if s != "ThisIsAColumnName" {
		t.Errorf("formatText(\"This-is_A-column_name\")= %s; want ThisIsAColumnName", s)
	}
}

func TestProcessCSV(t *testing.T) {
	s := processCSV("test_path/test_this-stuff.csv")

	expectedString := "type TestThisStuff struct {\n" +
		"\tLongColWithUnderscores string `csv:\"LongColWithUnderscores\"`\n" +
		"\tShortUnderscore        string `csv:\"ShortUnderscore\"`\n" +
		"\tShortDashes            string `csv:\"ShortDashes\"`\n" +
		"\tPlain                  string `csv:\"Plain\"`\n" +
		"}"

	if strings.Compare(s, expectedString) != 0 {
		t.Errorf("processCSV returned an incorrect struct")
	}
}
