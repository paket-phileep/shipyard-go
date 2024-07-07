package out

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"shipyard/controller"
	"shipyard/types"
	"shipyard/utils"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-isatty"
)

type model struct {
	spinner  spinner.Model
	results  []types.Result
	packages []string
	index    int
	quitting bool
}

var (
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
	mainStyle = lipgloss.NewStyle().MarginLeft(1)
)

func InstallDependencies() {
	var (
		daemonMode bool
		showHelp   bool
		opts       []tea.ProgramOption
	)

	flag.BoolVar(&daemonMode, "d", false, "run as a daemon")
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.Parse()

	if showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if daemonMode || !isatty.IsTerminal(os.Stdout.Fd()) {
		// If we're in daemon mode don't render the TUI
		opts = []tea.ProgramOption{tea.WithoutRenderer()}
	} else {
		// If we're in TUI mode, discard log output
		log.SetOutput(io.Discard)
	}

	dependencies, err := controller.ReadDependencies("./dependencies.json")
	if err != nil {
		fmt.Println("Error reading dependencies:", err)
		os.Exit(1)
	}

	fmt.Println("Dependencies to install:")
	for _, pkg := range dependencies.Packages {
		fmt.Println(pkg)
	}

	p := tea.NewProgram(newModel(dependencies.Packages), opts...)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting Bubble Tea program:", err)
		os.Exit(1)
	}
}

func newModel(packages []string) model {
	const showLastResults = 5

	sp := spinner.New()
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("206"))

	return model{
		spinner:  sp,
		results:  make([]types.Result, showLastResults),
		packages: packages,
	}
}

func (m model) Init() tea.Cmd {
	log.Println("Starting work...")
	return tea.Batch(
		m.spinner.Tick,
		m.runPretendProcess,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.quitting = true
		return m, tea.Quit
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case processFinishedMsg:
		d := time.Duration(msg)
		res := types.Result{Emoji: utils.RandomEmoji(), PackageName: m.packages[m.index], Duration: d}
		log.Printf("%s Package %s installed in %s", res.Emoji, res.PackageName, res.Duration)
		m.results = append(m.results[1:], res)
		m.index++
		if m.index < len(m.packages) {
			return m, m.runPretendProcess
		}
		m.quitting = true
		return m, tea.Quit
	default:
		return m, nil
	}
}

func (m model) View() string {
	s := "\n" +
		m.spinner.View() + " Installing packages...\n\n"

	for _, res := range m.results {
		if res.Duration == 0 {
			s += "........................\n"
		} else {
			s += fmt.Sprintf("%s Package %s installed in %s\n", res.Emoji, res.PackageName, res.Duration)
		}
	}

	s += helpStyle("\nPress any key to exit\n")

	if m.quitting {
		s += "\n"
	}

	return mainStyle.Render(s)
}

// processFinishedMsg is sent when a pretend process completes.
type processFinishedMsg time.Duration

// runPretendProcess simulates a long-running process.
func (m model) runPretendProcess() tea.Msg {
	pause := time.Duration(rand.Int63n(899)+100) * time.Millisecond // nolint:gosec
	time.Sleep(pause)
	return processFinishedMsg(pause)
}
