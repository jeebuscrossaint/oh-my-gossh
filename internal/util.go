package internal

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"

	//"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/glamour"
)

// Check handles error checking with optional fatal exit
func Check(e error, context string, fatal bool) {
	if e != nil {
		fmt.Printf("Error running program - In %v: %v", context, e)
		if fatal {
			os.Exit(1)
		}
	}
}

// Item represents a selectable menu item
type Item struct {
	TitleText string
	Desc      string
}

func (i Item) Title() string       { return i.TitleText }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.TitleText }

// GetMarkdown reads and returns markdown file contents
func GetMarkdown(filename string) string {
	// For core pages (main, about, contact)
	basePath := os.ExpandEnv("$HOME/.config/ohmygossh")
	var fullPath string
    
	// Check if it's a project file
	if strings.HasPrefix(filename, "projects/") {
	    fullPath = filepath.Join(basePath, filename)
	} else {
	    // Core pages
	    fullPath = filepath.Join(basePath, filename + ".md")
	}
    
	fileData, err := os.ReadFile(fullPath)
	Check(err, "Markdown File IO", false)
	return string(fileData)
    }

// OpenProject renders a markdown project file with glamour
func OpenProject(selectedProject int, projects []string, viewportWidth int) string {
	if selectedProject < 0 || selectedProject >= len(projects) {
		return "Invalid project selection"
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithWordWrap(viewportWidth - 20),
	)
	Check(err, "Glamour renderer creation", false)

	project := projects[selectedProject]
	content := GetMarkdown(project)

	rendered, err := renderer.Render(content)
	Check(err, "Project render", false)

	return rendered
}

// CalculateNavItemSize returns the width and height of navigation items
func CalculateNavItemSize(title string) (width int, height int) {
	switch title {
	case "hosts":
		return 10, 2
	case "groups":
		return 11, 2
	case "settings":
		return 13, 2
	default:
		return 8, 2
	}
}

// Max returns the larger of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// CountLines returns the number of lines in a string
func CountLines(s string) int {
	return len(strings.Split(s, "\n"))
}
