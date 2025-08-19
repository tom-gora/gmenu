// Package menu
package menu

import (
	"encoding/json"
	"fmt"
	"gmenu/cli"
	"os"
)

func ReadJSONFile(pFilePath *string) (menuEntries []cli.MenuEntry, err error) {
	jsonData, err := os.ReadFile(*pFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config file '%v': %v\n", *pFilePath, err)
		os.Exit(1)
	}
	var unmarshalledEntries []cli.MenuEntry
	if err := json.Unmarshal(jsonData, &unmarshalledEntries); err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling JSON from '%v': %v\n", pFilePath, err)
		os.Exit(1)
	}

	return unmarshalledEntries, nil
}
