package app

import (
	"fmt"
	"log"
	"os"
	"shipyard/cmd/fs"
	"shipyard/cmd/notif"

	logger "github.com/charmbracelet/log"
)

func ExtractImages() {
	// Initialize logging
	logger.Info("Initializing repositories")
	logFile, err := os.OpenFile("init_repos.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		os.Exit(1)
	}

	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("Starting the process of initializing repositories...")
	var bundle fs.Bundle
	bundle = fs.ReadBundles()
	// Iterate through the images and process each image
	for _, image := range bundle.Images {
		notif.CompletedUnmarshalingImages(image)
		if image.URI == "" || image.Destination == "" {
			log.Printf("Invalid image entry in bundle.yml. Skipping entry: %v\n", image)
			fmt.Println("Invalid image entry in bundle.yml. Skipping...")
			continue
		}

		//
	}
}
