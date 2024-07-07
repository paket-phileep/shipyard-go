package fs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// updatePackageJSON updates the "name" field in package.json if it exists,
// or creates package.json with the specified package name if it doesn't exist
func UpdatePackageJSON(targetDir, packageName string) error {
	packageJSONPath := filepath.Join(targetDir, "package.json")
	data, err := os.ReadFile(packageJSONPath)
	if err != nil {
		// Create package.json if it doesn't exist
		if os.IsNotExist(err) {
			fmt.Println("Creating package.json...")
			content := fmt.Sprintf(`{
  "name": "%s"
}`, packageName)
			if writeErr := os.WriteFile(packageJSONPath, []byte(content), 0644); writeErr != nil {
				log.Printf("Failed to create package.json at %s: %v\n", packageJSONPath, writeErr)
				return writeErr
			}
			return nil
		}
		log.Printf("Failed to read package.json at %s: %v\n", packageJSONPath, err)
		return err
	}

	// Update package.json if it exists
	fmt.Println("Updating package.json...")
	var packageJSON map[string]interface{}
	if err := yaml.Unmarshal(data, &packageJSON); err != nil {
		log.Printf("Failed to parse package.json at %s: %v\n", packageJSONPath, err)
		return err
	}
	packageJSON["name"] = packageName
	newData, err := yaml.Marshal(&packageJSON)
	if err != nil {
		log.Printf("Failed to marshal updated package.json at %s: %v\n", packageJSONPath, err)
		return err
	}
	if err := os.WriteFile(packageJSONPath, newData, 0644); err != nil {
		log.Printf("Failed to write updated package.json at %s: %v\n", packageJSONPath, err)
		return err
	}
	return nil
}

// updateGoMod updates the "module" name in go.mod if it exists
func UpdateGoMod(targetDir, packageName string) error {
	goModPath := filepath.Join(targetDir, "go.mod")
	data, err := os.ReadFile(goModPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Do nothing if go.mod doesn't exist
		}
		log.Printf("Failed to read go.mod at %s: %v\n", goModPath, err)
		return err
	}

	// Update go.mod if it exists
	fmt.Println("Updating go.mod...")
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "module ") {
			lines[i] = "module " + packageName
			break
		}
	}
	newData := strings.Join(lines, "\n")
	if err := os.WriteFile(goModPath, []byte(newData), 0644); err != nil {
		log.Printf("Failed to write updated go.mod at %s: %v\n", goModPath, err)
		return err
	}
	return nil
}
