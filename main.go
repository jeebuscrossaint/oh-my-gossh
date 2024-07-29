package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	//"github.com/charmbracelet/bubbles/paginator"
	//"github.com/charmbracelet/bubbles/viewport"
	//tea "github.com/charmbracelet/bubbletea"
	//"github.com/charmbracelet/glamour"
	"golang.org/x/term"
)

/*
type KeyMap struct {
	Navigate key.Binding
	Up       key.Binding
	Down     key.Binding
	Left     key.Binding
	Right    key.Binding
	LCycle   key.Binding
	RCycle   key.Binding
	Enter    key.Binding
	Back     key.Binding
	Help     key.Binding
	Quit     key.Binding
}

var DefaultKeyMap = KeyMap{
	Navigate: key.NewBinding(
		key.WithKeys("j", "k", "up", "down"),
		key.WithHelp("↑↓", "navigate"),
	),
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "down"),
	),
	Left: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("←/h", "prev page"),
	),
	Right: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("→/l", "next page"),
	),
	LCycle: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("^tab", "prev section"),
	),
	RCycle: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "section"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("enter", "select"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc", "backspace"),
		key.WithHelp("esc", "go back"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

type model struct {
	pageIndex    int
	pages        []string
	projects     []string
	projectOpen  bool
	openProject  int
	projectView  string
	clickCounter int
	viewport     viewport.Model
	list         list.Model
	content      string
	keys         KeyMap
	help         help.Model
	ready        bool
}
*/

const ASCIIHeader string = `
    :::     ::::    ::::      :::     :::::::::  ::::    :::     ::: ::::::::::: :::    :::      :::::::::     ::: ::::::::::: :::::::::: :::        
  :+: :+:   +:+:+: :+:+:+   :+: :+:   :+:    :+: :+:+:   :+:   :+: :+:   :+:     :+:    :+:      :+:    :+:  :+: :+:   :+:     :+:        :+:        
 +:+   +:+  +:+ +:+:+ +:+  +:+   +:+  +:+    +:+ :+:+:+  +:+  +:+   +:+  +:+     +:+    +:+      +:+    +:+ +:+   +:+  +:+     +:+        +:+        
+#++:++#++: +#+  +:+  +#+ +#++:++#++: +#++:++#:  +#+ +:+ +#+ +#++:++#++: +#+     +#++:++#++      +#++:++#+ +#++:++#++: +#+     +#++:++#   +#+        
+#+     +#+ +#+       +#+ +#+     +#+ +#+    +#+ +#+  +#+#+# +#+     +#+ +#+     +#+    +#+      +#+       +#+     +#+ +#+     +#+        +#+        
#+#     #+# #+#       #+# #+#     #+# #+#    #+# #+#   #+#+# #+#     #+# #+#     #+#    #+#      #+#       #+#     #+# #+#     #+#        #+#        
###     ### ###       ### ###     ### ###    ### ###    #### ###     ### ###     ###    ###      ###       ###     ### ###     ########## ########## 
`

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

	fmt.Println(quote)
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
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		if lineCount == lineNumber {
			return scanner.Text(), nil
		}
	}

	return "", fmt.Errorf("line not found")
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

// end text animation and introduction manipulation

func main() {
	welcomeMessage := "Welcome to my TUI Portfolio!\nPlease hire me!\nPress any key to continue..."
	typeOutText(welcomeMessage)
	waitForKeyPress()
	fmt.Print("\033[H\033[2J")
	fmt.Println(ASCIIHeader)
	generateQuote()
}
