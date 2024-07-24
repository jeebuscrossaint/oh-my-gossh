package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textInput textinput.Model
	err       error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter text"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			fmt.Println("You entered:", m.textInput.Value())
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.textInput.Width = msg.Width
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"Enter some text:\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	)
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
