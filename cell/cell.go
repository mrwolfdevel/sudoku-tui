package sudoku

import (
	"log"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/*
This example assumes an existing understanding of commands and messages. If you
haven't already read our tutorials on the basics of Bubble Tea and working
with commands, we recommend reading those first.

Find them at:
https://github.com/charmbracelet/bubbletea/tree/master/tutorials/commands
https://github.com/charmbracelet/bubbletea/tree/master/tutorials/basics
*/

// sessionState is used to track which model is focused

var (
	// Available spinners
	modelStyle = lipgloss.NewStyle().
			Width(3).
			Height(1).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.HiddenBorder())
	focusedModelStyle = lipgloss.NewStyle().
				Width(3).
				Height(1).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("69"))
)

type Model struct {
	focus  bool 
	value  int 
}

func newModel() Model {
	m := Model{focus: true, value: 0}
	return m
}

func (m Model) Init() tea.Cmd {
	// start the timer and spinner on program start
	return nil 
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            return m, tea.Quit
	}
	}
	return m, nil 
}

func (m Model) View() string {
	var s string
	var value = strconv.Itoa(m.value)
	if m.focus == true {
		s += lipgloss.JoinHorizontal(lipgloss.Top, focusedModelStyle.Render( value))
	} else {
		s += lipgloss.JoinHorizontal(lipgloss.Top, modelStyle.Render(value))
	}
	return s
}


func main() {
	p := tea.NewProgram(newModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
