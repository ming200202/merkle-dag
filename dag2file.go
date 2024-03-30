package merkledag

import (
	"encoding/json"
	"strings"
)

// Hash to file
func Hash2File(store KVStore, hash []byte, path string, hp HashPool) []byte {

	exists, err := store.Has(hash)
	if err != nil {
		// Handle error
		return nil
	}
	if !exists {
		// Hash does not exist in the KVStore
		return nil
	}

	data, err := store.Get(hash)
	if err != nil {
		// Handle error
		return nil
	}
	var tree map[string]interface{}
	if err := json.Unmarshal(data, &tree); err != nil {
		// Handle error
		return nil
	}

	var currentNode interface{} = tree
	components := strings.Split(path, "/")
	for _, component := range components {
		if dir, ok := currentNode.(map[string]interface{}); ok {

			node, found := dir[component]
			if !found {
				return nil // Component not found in directory
			}
			currentNode = node
		} else {
			return nil // Current node is not a directory
		}
	}

	if file, ok := currentNode.([]byte); ok {
		return file
	}

	return nil
}
