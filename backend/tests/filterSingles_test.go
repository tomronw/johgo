package tests

import (
	"fmt"
	"regexp"
	"testing"
)

func TestFilterNames(t *testing.T) {
	names := []string{
		"test 1/1",
		"test 22/22",
		"test 333/333",
		"normal name",
		"other normal name",
	}

	expectedNames := []string{
		"normal name",
		"other normal name",
	}

	regexPattern := `[0-9]{1,3}/[0-9]{1,3}`
	regex, err := regexp.Compile(regexPattern)

	if err != nil {
		t.Fatalf("Error compiling regex: %v", err)
	}

	var filteredNames []string
	for _, name := range names {
		if !regex.MatchString(name) {
			filteredNames = append(filteredNames, name)
			fmt.Println(name)
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
