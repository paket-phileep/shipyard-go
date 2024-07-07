package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// cloneOrForkRepo attempts to clone a repository from a given URL into a target directory
// If cloning fails and the repository is not owned by the specified owner, it attempts to fork the repository
func cloneOrForkRepo(repoURL, targetDir, owner string) error {
	fmt.Printf("Cloning repository %s into %s...\n", repoURL, targetDir)
	if err := RunCommand("git", "clone", repoURL, targetDir); err != nil {
		log.Printf("Failed to clone repository %s: %v\n", repoURL, err)
		fmt.Printf("Failed to clone repository %s. Attempting to fork...\n", repoURL)
		// TODO: Implement forking logic here
		return err
	}
	return nil
}

// processRepo performs the following tasks for a given repository URL and target directory:
// 1. Clones or forks the repository
// 2. Updates or creates package.json with the specified package name
// 3. Updates go.mod with the specified package name
// 4. Recursively updates go.mod in subdirectories
func ProcessRepo(repoURL, targetDir, packageName string) error {
	// Clone or fork the repository
	if err := cloneOrForkRepo(repoURL, targetDir, packageName); err != nil {
		log.Printf("Failed to process repository %s: %v\n", repoURL, err)
		return err
	}

	// Update package.json if it exists or create if it doesn't
	if err := UpdatePackageJSON(targetDir, packageName); err != nil {
		log.Printf("Failed to update package.json in %s: %v\n", targetDir, err)
		return err
	}

	// Update go.mod if it exists
	if err := UpdateGoMod(targetDir, packageName); err != nil {
		log.Printf("Failed to update go.mod in %s: %v\n", targetDir, err)
		return err
	}

	// Recursively update go.mod in subdirectories
	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error walking through directory %s: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			subGoModPath := filepath.Join(path, "go.mod")
			if _, err := os.Stat(subGoModPath); err == nil {
				if err := UpdateGoMod(path, packageName+"/"+filepath.Base(path)); err != nil {
					log.Printf("Failed to update go.mod in subdirectory %s: %v\n", path, err)
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("Failed to walk through directory %s: %v\n", targetDir, err)
	}
	return err
}
