// Package menu
package menu

import (
	"encoding/json"
	"fmt"
	"gmenu/cli"
	"os"
	"path/filepath"
	"strings"
)

func isPath(input string) bool {
	if len(input) == 0 {
		return false
	}
	if filepath.Separator == '/' {
		return input[0] == '/' || strings.ContainsRune(input, filepath.Separator)
	}
	if filepath.Separator == '\\' {
		return (len(input) >= 2 && ((input[1] == ':' && (input[0] >= 'a' && input[0] <= 'z' || input[0] >= 'A' && input[0] <= 'Z')) || (input[0] == '\\' && input[1] == '\\'))) ||
			strings.ContainsRune(input, filepath.Separator)
	}
	return strings.ContainsRune(input, filepath.Separator)
}

func getJSONdata(input *string) ([]byte, error) {
	if isPath(*input) {
		return os.ReadFile(*input)
	} else {
		return []byte(*input), nil
	}
}

func GetMenuEntriesFromJSON(inputJSON *string) (menuEntries []cli.MenuEntry, err error) {
	// try and see if the value is path to valid json or json string
	data, err := getJSONdata(inputJSON)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config file '%v': %v\n", filepath.Base(*inputJSON), err)
		os.Exit(1)
	}

	var unmarshalledEntries []cli.MenuEntry
	if err := json.Unmarshal(data, &unmarshalledEntries); err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling JSON from '%v': %v\n", inputJSON, err)
		os.Exit(1)
	}

	return unmarshalledEntries, nil
}
