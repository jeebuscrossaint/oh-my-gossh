package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	//"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	//"github.com/charmbracelet/glamour"
)

type model struct {
	viewport viewport.Model
	content  string
	err      error
}

func initialModel(content string) model {
	vp := viewport.New(80, 20)
	vp.SetContent(content)

	return model{
		viewport: vp,
		content:  content,
		err:      nil,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			m.viewport.LineUp(1)
		case "down", "j":
			m.viewport.LineDown(1)
		}
	}

	m.viewport, _ = m.viewport.Update(msg)
	return m, nil
}

func (m model) View() string {
	return centerText(m.content, 80)
}

func centerText(text string, width int) string {
	lines := strings.Split(text, "\n")
	var centeredLines []string
	for _, line := range lines {
		padding := (width - len(line)) / 2
		if padding < 0 {
			padding = 0
		}
		centeredLines = append(centeredLines, strings.Repeat(" ", padding)+line)
	}
	return strings.Join(centeredLines, "\n")
}

func main() {
	filePath := "ascii.conf"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	p := tea.NewProgram(initialModel(string(content)))
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
