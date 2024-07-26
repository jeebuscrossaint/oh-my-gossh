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
	"golang.org/x/term"
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
	// Read the ASCII header from ascii.conf
	asciiHeader, err := ioutil.ReadFile("assets/ascii.conf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading ASCII header file: %v\n", err)
		os.Exit(1)
	}

	// Apply ANSI escape code for green color and center the ASCII header
	greenAsciiHeader := "\033[32m" + centerText(string(asciiHeader), 80) + "\033[0m"

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

	// Wait for a single keypress
	waitForKeyPress()

	// Clear the screen
	fmt.Print("\033[H\033[2J")

	// Initialize the TUI program with the content and header
	p := tea.NewProgram(initialModel(string(renderedContent), greenAsciiHeader))
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func waitForKeyPress() {
	// Put terminal into raw mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Failed to set raw mode:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	// Read a single byte
	buffer := make([]byte, 1)
	os.Stdin.Read(buffer)
}
