package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
	//"github.com/charmbracelet/bubbles/paginator"
	//"github.com/charmbracelet/bubbles/viewport"
	//tea "github.com/charmbracelet/bubbletea"
	//"github.com/charmbracelet/glamour"
	//"golang.org/x/term"
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

func check(e error, check string) {
	if e != nil {
		fmt.Printf("Error running program - In %v: %v", check, e)
	}
}

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

func main() {
	fmt.Println(ASCIIHeader)
	generateQuote()
}
