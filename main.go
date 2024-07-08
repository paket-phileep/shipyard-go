package main

import (
	"shipyard/cmd/app"

	logger "github.com/charmbracelet/log"
)

func main() {
	logger.Warn("building battleship!")
	app.ExtractImages()
}
