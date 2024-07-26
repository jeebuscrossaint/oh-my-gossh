package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

type model struct {
	viewport  viewport.Model
	paginator paginator.Model
	content   string
	header    string
	err       error
}

func initialModel(content, header string) model {
	vp := viewport.New(80, 20)
	vp.SetContent(content)

	p := paginator.New()
	p.SetTotalPages(len(content)/(80*20) + 1)

	return model{
		viewport:  vp,
		paginator: p,
		content:   content,
		header:    header,
		err:       nil,
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
		case "left", "h":
			m.paginator.PrevPage()
		case "right", "l":
			m.paginator.NextPage()
		}
	}

	m.viewport, _ = m.viewport.Update(msg)
	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("%s\n\n%s\n\n%s", centerText(m.header, 80), m.viewport.View(), m.paginator.View())
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

func extractHeader(content string) (string, string) {
	lines := strings.Split(content, "\n")
	var headerLines []string
	var contentLines []string
	inHeaderSection := false

	for _, line := range lines {
		if strings.HasPrefix(line, "# ASCII") {
			inHeaderSection = true
			continue
		}
		if inHeaderSection {
			if strings.HasPrefix(line, "#") {
				inHeaderSection = false
			} else {
				headerLines = append(headerLines, line)
				continue
			}
		}
		contentLines = append(contentLines, line)
	}

	return strings.Join(headerLines, "\n"), strings.Join(contentLines, "\n")
}

func main() {
	filePath := "portfolio.md"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	header, modifiedContent := extractHeader(string(content))

	renderedContent, err := glamour.RenderBytes([]byte(modifiedContent), "dark")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering markdown: %v\n", err)
		os.Exit(1)
	}

	p := tea.NewProgram(initialModel(string(renderedContent), header))
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
