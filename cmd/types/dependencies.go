package types

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
)

type Dependencies struct {
	Packages []string `json:"packages"`
}

type Result struct {
	Emoji       string
	Duration    time.Duration
	PackageName string
}

type Model struct {
	Spinner  spinner.Model
	Results  []Result
	Quitting bool
}
