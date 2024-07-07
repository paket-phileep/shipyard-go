package fs

import (
	"fmt"
	"log"
	"os"
	"shipyard/notif"

	"gopkg.in/yaml.v2"
)

type Image struct {
	URI         string `yaml:"uri"`
	Destination string `yaml:"destination"`
	Proxy       bool   `yaml:"proxy"`
}

// Config struct is used to match the contents of the "config" section in the bundle.yaml file
type Config struct {
	Name    string `yaml:"name"`
	Package string `yaml:"package"`
}

// Bundle struct is used to match the contents of a "bundle.yaml" file,
// which includes the configuration and the list of repositories
type Bundle struct {
	Config Config  `yaml:"config"`
	Images []Image `yaml:"images"`
}

// InitRepos is the main function to read the bundle.yml file,
// parse it, and process each repository listed in the repositories section
func ReadBundles() Bundle {
	// Read the bundle.yml file
	notif.ReadingBundleFile()
	bundleFilePath := "/Users/sullemanhossam/Desktop/shipyard/bundle.yml"
	data, err := os.ReadFile(bundleFilePath)
	if err != nil {
		log.Printf("Failed to read %s: %v. Exiting...\n", bundleFilePath, err)
		fmt.Printf("Failed to read %s. Exiting...\n", bundleFilePath)
		os.Exit(1)
	}
	notif.CompletedReadingBundleFile(data)

	notif.UnmarshalingBundles()
	var bundle Bundle
	if err := yaml.Unmarshal(data, &bundle); err != nil {
		log.Printf("Failed to parse %s: %v. Exiting...\n", bundleFilePath, err)
		fmt.Printf("Failed to parse %s. Exiting...\n", bundleFilePath)
		os.Exit(1)
	}
	notif.CompletedUnmarshalingBundles(bundle)

	return bundle

}
