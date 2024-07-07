package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config struct is used to match the contents of the "config" section in the bundle.yaml file
type Config struct {
	Name    string `yaml:"name"`
	Package string `yaml:"package"`
}

// Bundle struct is used to match the contents of a "bundle.yaml" file,
// which includes the configuration and the list of repositories
type Bundle struct {
	Config       Config     `yaml:"config"`
	Repositories [][]string `yaml:"repositories"`
}

// InitRepos is the main function to read the bundle.yml file,
// parse it, and process each repository listed in the repositories section
func main() {
	// Read the bundle.yml file
	bundleFilePath := "/Users/sullemanhossam/Desktop/shipyard/bundles.yml"
	data, err := os.ReadFile(bundleFilePath)
	if err != nil {
		log.Printf("Failed to read %s: %v. Exiting...\n", bundleFilePath, err)
		fmt.Printf("Failed to read %s. Exiting...\n", bundleFilePath)
		os.Exit(1)
	}

	var bundle Bundle
	if err := yaml.Unmarshal(data, &bundle); err != nil {
		log.Printf("Failed to parse %s: %v. Exiting...\n", bundleFilePath, err)
		fmt.Printf("Failed to parse %s. Exiting...\n", bundleFilePath)
		os.Exit(1)
	}

	// Iterate through the repositories and process each repository
	for _, image := range bundle.Repositories {
		if len(image) != 2 {
			log.Printf("Invalid image entry in bundle.yml. Skipping entry: %v\n", image)
			fmt.Println("Invalid image entry in bundle.yml. Skipping...")
			continue
		}

		repoURL := image[0]
		targetDir := image[1]
		packageName := fmt.Sprintf("%s/%s", bundle.Config.Package, filepath.Base(repoURL))

		if err := processRepo(repoURL, targetDir, packageName); err != nil {
			log.Printf("Failed to process repository %s: %v\n", repoURL, err)
			fmt.Printf("Failed to process repository %s. Error: %v\n", repoURL, err)
		}
	}
}

func InitRepos() {
	// Initialize logging
	logFile, err := os.OpenFile("init_repos.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("Starting InitRepos")

	// Run InitRepos
	main()

	log.Println("Completed InitRepos")
}
