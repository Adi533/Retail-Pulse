package store

import (
	"fmt"
)

// Dummy function to simulate store details lookup
// For simplicity, using a hardcoded map
// I created this function to separate the store details but thought i can make it work with a hardcoded map
// If in production i would set up a database for storing data like postgressql or any databases

func GetStoreDetails(storeID string) (string, string, error) {
	storeData := map[string]struct {
		Name     string
		AreaCode string
	}{
		"S00339218": {"Store A", "12345"},
		"S01408764": {"Store B", "67890"},
	}

	if store, ok := storeData[storeID]; ok {
		return store.Name, store.AreaCode, nil
	}

	return "", "", fmt.Errorf("store ID %s not found", storeID)
}
