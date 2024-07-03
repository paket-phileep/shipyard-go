package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"shipyard/types"
)

func ReadFile(filePath string) (types.Dependencies, error) {
	var deps types.Dependencies

	file, err := os.Open(filePath)
	if err != nil {
		return deps, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return deps, err
	}

	err = json.Unmarshal(bytes, &deps)
	return deps, err
}
