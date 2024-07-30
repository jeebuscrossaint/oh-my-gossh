package main

import (
	//"context"
	//"errors"
	"fmt"
	//"net"
	"bufio"
	"io/ioutil"
	"math/rand"
	"os"

	"strings"
	"time"

	//"github.com/charmbracelet/log"
	//"github.com/charmbracelet/ssh"
	//"github.com/charmbracelet/wish"
	//"github.com/charmbracelet/wish/activeterm"
	//"github.com/charmbracelet/wish/bubbletea"
	//"github.com/charmbracelet/wish/logging"

	//"github.com/charmbracelet/bubbles/help"
	//"github.com/charmbracelet/bubbles/key"
	//"github.com/charmbracelet/bubbles/list"
	//"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// constants include host port and ascii for now
const (
	ASCIIHeader string = `
    :::     ::::    ::::      :::     :::::::::  ::::    :::     ::: ::::::::::: :::    :::      :::::::::     ::: ::::::::::: :::::::::: :::        
  :+: :+:   +:+:+: :+:+:+   :+: :+:   :+:    :+: :+:+:   :+:   :+: :+:   :+:     :+:    :+:      :+:    :+:  :+: :+:   :+:     :+:        :+:        
 +:+   +:+  +:+ +:+:+ +:+  +:+   +:+  +:+    +:+ :+:+:+  +:+  +:+   +:+  +:+     +:+    +:+      +:+    +:+ +:+   +:+  +:+     +:+        +:+        
+#++:++#++: +#+  +:+  +#+ +#++:++#++: +#++:++#:  +#+ +:+ +#+ +#++:++#++: +#+     +#++:++#++      +#++:++#+ +#++:++#++: +#+     +#++:++#   +#+        
+#+     +#+ +#+       +#+ +#+     +#+ +#+    +#+ +#+  +#+#+# +#+     +#+ +#+     +#+    +#+      +#+       +#+     +#+ +#+     +#+        +#+        
#+#     #+# #+#       #+# #+#     #+# #+#    #+# #+#   #+#+# #+#     #+# #+#     #+#    #+#      #+#       #+#     #+# #+#     #+#        #+#        
###     ### ###       ### ###     ### ###    ### ###    #### ###     ### ###     ###    ###      ###       ###     ### ###     ########## ########## 
`

	host = "localhost"
	port = "19"
)

// check is a helper function to check for errors
func check(e error, check string) {
	if e != nil {
		fmt.Printf("Error running program - In %v: %v", check, e)
	}
}

// begin area for quote generation
func generateQuote() {
	filePath := "assets/quotes.txt"
	lines, err := countLines(filePath)
	check(err, "counting lines")

	lineNumber := getRandomLine(lines)
	quote, err := getLineContent(filePath, lineNumber)
	check(err, "getting line content")

	// Print the quote in green color
	fmt.Printf("\033[38;2;0;255;0m%s\033[0m\n", quote)
}

func countLines(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	return lineCount, scanner.Err()
}

func getRandomLine(lines int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(lines) + 1
}

func getLineContent(filePath string, lineNumber int) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentLine := 0
	for scanner.Scan() {
		currentLine++
		if currentLine == lineNumber {
			return scanner.Text(), nil
		}
	}

	return "", fmt.Errorf("line %d not found", lineNumber)
}

// end area for quote generation

// begin text animation and introduction manipulation

func typeOutText(text string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		for i := 0; i <= len(line); i++ {
			fmt.Print("\033[H\033[2J") // Clear the screen
			fmt.Printf("\033[38;2;0;255;0m%s\033[0m\n", line[:i])
			time.Sleep(50 * time.Millisecond)
		}
		time.Sleep(500 * time.Millisecond) // Pause after each line
	}
}

func waitForKeyPress() {
	buf := make([]byte, 1)
	os.Stdin.Read(buf)
}

// end text animation and introduction manipulation

type model struct {
	quote       string
	buttons     []string
	activeBtn   int
	markdown    string
	fileName    string
	currentView string
	quitting    bool
}

func initialModel() model {
	initialView := "Home"
	markdown, fileName := renderMarkdown(initialView)
	return model{
		buttons:     []string{"About", "Home", "Stuff"},
		activeBtn:   1, // Set to "Home" by default
		fileName:    fileName,
		currentView: initialView,
		markdown:    markdown,
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
			m.quitting = true
			return m, tea.Quit
		case "left":
			if m.activeBtn > 0 {
				m.activeBtn--
				markdown, _ := renderMarkdown(m.buttons[m.activeBtn])
				m.markdown = markdown
			}
		case "right":
			if m.activeBtn < len(m.buttons)-1 {
				m.activeBtn++
				markdown, _ := renderMarkdown(m.buttons[m.activeBtn])
				m.markdown = markdown
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "Thanks for reading!\nGoodbye..."
	}
	var b strings.Builder

	// Print the ASCII header
	// Print the quote
	b.WriteString(fmt.Sprintf("\033[38;2;0;255;0m%s\033[0m\n", m.quote))

	// Print the buttons
	for i, btn := range m.buttons {
		if i == m.activeBtn {
			b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(fmt.Sprintf("[%s]", btn)))
		} else {
			b.WriteString(fmt.Sprintf(" %s ", btn))
		}
	}

	// Draw a separating line with the file name
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80 // Default width if unable to get terminal size
	}
	lineWidth := width - len(m.fileName) - 6 // Subtract 6 for 3 spaces on each side
	if lineWidth < 0 {
		lineWidth = 0 // Ensure lineWidth is not negative
	}
	b.WriteString(fmt.Sprintf("\n%s %s%s\n", m.fileName, strings.Repeat("_", lineWidth), strings.Repeat(" ", 3)))

	// Print the markdown content
	b.WriteString(m.markdown)

	return b.String()
}

func renderMarkdown(button string) (string, string) {
	var fileName string
	switch button {
	case "About":
		fileName = "about.md"
	case "Home":
		fileName = "home.md"
	case "Stuff":
		fileName = "stuff.md"
	}

	content, err := ioutil.ReadFile("assets/" + fileName)
	if err != nil {
		return fmt.Sprintf("Error reading file: %v", err), fileName
	}

	renderedContent, err := glamour.Render(string(content), "dark")
	if err != nil {
		return fmt.Sprintf("Error rendering markdown: %v", err), fileName
	}
	return renderedContent, fileName
}

func main() {
	welcomeMessage := "Welcome to my TUI Portfolio!\nPlease hire me!\nPress any key to continue..."
	typeOutText(welcomeMessage)
	waitForKeyPress()
	fmt.Print("\033[H\033[2J")
	fmt.Printf("\033[38;2;0;255;0m%s\033[0m\n", ASCIIHeader)
	generateQuote()

	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	finalModel := m.(model)
	if finalModel.quitting {
		typeOutText(finalModel.View())
	}
}
