package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type model struct {
	password      string
	selectedIndex int
	title         string
	width         int
	height        int
}

func (m model) Init() tea.Cmd {
	// Get terminal size
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80
		height = 24
	}
	m.width = width
	m.height = height
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.selectedIndex > 0 {
				m.selectedIndex--
			}
		case "down":
			if m.selectedIndex < 1 {
				m.selectedIndex++
			}
		case "enter":
			if m.selectedIndex == 0 {
			passlength := 10
			password := ""
			for i := 0; i < passlength; i++ {
				random := rand.Float64()*(126-33) + 33
				password += string(rune(int(math.Floor(random))))
			}
			m.password = password
		} else if m.selectedIndex == 1{
			return m,tea.Quit
		}
		case "q":
			return m,tea.Quit
		
		}
	}
	return m, nil
}


func (m model) View() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Bold(true).
		Padding(1, 0).
		Align(lipgloss.Center).
		Render(`
███╗   ██╗███████╗ ██████╗     ██████╗  ██████╗ ██╗  ██╗
████╗  ██║██╔════╝██╔═══██╗    ██╔══██╗██╔═══██╗╚██╗██╔╝
██╔██╗ ██║█████╗  ██║   ██║    ██████╔╝██║   ██║ ╚███╔╝ 
██║╚██╗██║██╔══╝  ██║   ██║    ██╔══██╗██║   ██║ ██╔██╗ 
██║ ╚████║███████╗╚██████╔╝    ██████╔╝╚██████╔╝██╔╝ ██╗
╚═╝  ╚═══╝╚══════╝ ╚═════╝     ╚═════╝  ╚═════╝ ╚═╝  ╚═╝`)

	
	passwordBox := passwordBoxStyle.Render(fmt.Sprintf("Generated Password:\n\n%s", m.password))

	
	options := []string{"Press [Enter] to Generate Password", "Press [Q] to Quit","press UP/Down to select"}
	renderedOptions := ""
	for i, option := range options {
		if i == m.selectedIndex {
			renderedOptions += highlightStyle.Render(option) + "\n"
		} else {
			renderedOptions += lipgloss.NewStyle().Faint(true).Render(option) + "\n"
		}
	}
	optionBox := textBoxStyle.Render(renderedOptions)

	
	fullView := lipgloss.JoinVertical(lipgloss.Center, title, passwordBox, optionBox)


	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, fullView)
}

func main() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}



func execCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	go func() {
		defer stdin.Close()
		stdin.Write([]byte(arg[len(arg)-1]))
	}()
	return cmd.Run()
}

func isWindows() bool {
	return os.PathSeparator == '\\'
}

func isMac() bool {
	return os.Getenv("OSTYPE") == "darwin"
}

func isLinux() bool {
	return os.Getenv("OSTYPE") == "linux"
}

// Styling
var (
	passwordBoxStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("15")).
		Bold(true).
		Padding(1, 2).
		Align(lipgloss.Center).
		Margin(1, 0)

	highlightStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")).
		Background(lipgloss.Color("15")).
		Bold(true).
		Padding(0, 1)

	textBoxStyle = lipgloss.NewStyle().
		Margin(1, 0).
		Align(lipgloss.Center)
)
