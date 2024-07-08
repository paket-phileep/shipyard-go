package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	rootFile := "docker-compose.yml"
	composeFiles := []string{}

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == "docker-compose.yml" && path != "./"+rootFile {
			composeFiles = append(composeFiles, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the path:", err)
		return
	}

	err = writeRootComposeFile(rootFile, composeFiles)
	if err != nil {
		fmt.Println("Error writing root compose file:", err)
	}
}

func writeRootComposeFile(rootFile string, composeFiles []string) error {
	var sb strings.Builder

	sb.WriteString("version: 'SHIPYARD V.0001'\n\n")

	for _, file := range composeFiles {
		if fileExists(file) {
			sb.WriteString(fmt.Sprintf("include:\n  - %s\n\n", file))
		} else {
			fmt.Printf("Warning: File %s does not exist and will not be included.\n", file)
		}
	}

	// Create the root file if it doesn't exist
	if !fileExists(rootFile) {
		err := ioutil.WriteFile(rootFile, []byte(sb.String()), 0644)
		if err != nil {
			return fmt.Errorf("error creating root compose file: %v", err)
		}
		fmt.Printf("Root compose file %s created.\n", rootFile)
	} else {
		err := ioutil.WriteFile(rootFile, []byte(sb.String()), 0644)
		if err != nil {
			return fmt.Errorf("error updating root compose file: %v", err)
		}
		fmt.Printf("Root compose file %s updated.\n", rootFile)
	}

	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
