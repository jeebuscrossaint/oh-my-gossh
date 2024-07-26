package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

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
	waiting   bool
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
		waiting:   true,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.waiting {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			m.waiting = false
		case tea.MouseMsg:
			if msg.Type == tea.MouseLeft {
				m.waiting = false
			}
		}
		return m, nil
	}

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
	if m.waiting {
		return "Welcome to my Portfolio! \nEnjoy your stay! \nPlease hire me! \n\nPress any key to continue..."
	}
	return fmt.Sprintf("%s\n\n%s\n\n%s", m.header, m.viewport.View(), m.paginator.View())
}

func typeOutText(text string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		for i := 0; i <= len(line); i++ {
			fmt.Print("\033[H\033[2J") // Clear the screen
			fmt.Println(line[:i])
			time.Sleep(50 * time.Millisecond)
		}
		time.Sleep(500 * time.Millisecond) // Pause after each line
	}
}

func main() {
	// Read the ASCII header from ascii.conf
	asciiHeader, err := ioutil.ReadFile("assets/ascii.conf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading ASCII header file: %v\n", err)
		os.Exit(1)
	}

	// Apply ANSI escape code for white color
	whiteAsciiHeader := "\033[37m" + string(asciiHeader) + "\033[0m"

	// Read the main content from portfolio.md
	content, err := ioutil.ReadFile("portfolio.md")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading content file: %v\n", err)
		os.Exit(1)
	}

	// Render the markdown content
	renderedContent, err := glamour.RenderBytes(content, "dark")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering markdown: %v\n", err)
		os.Exit(1)
	}

	// Display the typing animation welcome message
	welcomeMessage := "Welcome to my Portfolio! \nEnjoy your stay! \nPlease hire me! \nPress any key to continue..."
	typeOutText(welcomeMessage)

	// Initialize the TUI program with the content and header
	p := tea.NewProgram(initialModel(string(renderedContent), whiteAsciiHeader))
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
