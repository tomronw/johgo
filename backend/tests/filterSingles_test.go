package tests

import (
	"fmt"
	"johgo-search-engine/internal/core"
	"regexp"
	"testing"
)

func TestFilterNames(t *testing.T) {
	names := []string{
		"test 1/1",
		"test 22/22",
		"Pokemon Glastier (Holo) Lost Abyss s11 034/100",
		"normal name",
		"other normal name",
	}

	expectedNames := []string{
		"normal name",
		"other normal name",
	}

	var filteredNames []string
	for _, name := range names {
		match, err := regexp.MatchString("[a-zA-Z]{0,2}[0-9]{1,3}/[a-zA-Z]{0,2}[0-9]{1,3}|[a-zA-Z]{0,2}[0-9]{1,3}\\s/\\s[a-zA-Z]{0,2}[0-9]{1,3}", name)
		if err != nil {
			core.ErrorLogger.Printf("Error matching regex: %s", err.Error())
		}
		if !match {
			fmt.Println(name)
			filteredNames = append(filteredNames, name)
		}
	}

	if len(filteredNames) != len(expectedNames) {
		t.Fatalf("Expected length: %v, but got: %v", len(expectedNames), len(filteredNames))
	}

	for i, name := range filteredNames {
		if name != expectedNames[i] {
			t.Fatalf("Expected name: %v, but got: %v", expectedNames[i], name)
		}
	}
}
