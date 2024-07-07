package main

import (
	"shipyard/app"

	logger "github.com/charmbracelet/log"
)

func main() {
	logger.Warn("building battleship!")
	app.InitRepos()
}
