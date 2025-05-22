package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	password      string
	selectedIndex int
	done          bool
}

type doneMsg string

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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
				passlength := 10
				var password string
				for i := 0; i < passlength; i++ {
					random := rand.Float64()*(126-33) + 33
					password += string(rune(int(math.Floor(random))))
				}
				m.password = password
			
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	
	title := titleStyle.Render("Password Generator")

	options := []string{"Press [Enter] to Generate Password", "Press [Q] to Quit"}
	renderedOptions := ""
	for i, option := range options {
		if i == m.selectedIndex {
			renderedOptions += highlightStyle.Render(option) + "\n"
		} else {
			renderedOptions += lipgloss.NewStyle().Faint(true).Render(option) + "\n"
		}
	}

	// Always show the password box, even if the password is not generated yet
	content := passwordBoxStyle.Render(fmt.Sprintf("Generated Password:\n %s", m.password))

	return lipgloss.Place(screenWidth, 0, lipgloss.Center, lipgloss.Top,
		title) + "\n\n" + content + "\n\n" + textBoxStyle.Render(renderedOptions)
}

func main() {
	m := model{}

	prog := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := prog.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}

// Styling
var (
	titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true).Padding(2, 2).Align(lipgloss.Center)

	passwordBoxStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("15")).
		Bold(true).
		Padding(2, 2).
		Align(lipgloss.Center)

	highlightStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("0")).Background(lipgloss.Color("15")).Bold(true).Padding(0, 1)

	textBoxStyle = lipgloss.NewStyle()

	screenWidth = 50
)
