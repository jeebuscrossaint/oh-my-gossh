package app

import (
	"fmt"
	"oh-my-gossh/internal"
	"strings"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

// ViewportHeader renders the header for the viewport with a title and line
func (m Model) ViewportHeader(pageTitle string) string {
	title := BorderTitleStyle.Render(pageTitle)
	line := strings.Repeat("─", internal.Max(0, m.Viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

// ViewportFooter renders the footer showing scroll percentage
func (m Model) ViewportFooter() string {
	info := BorderInfoStyle.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := strings.Repeat("─", internal.Max(0, m.Viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

// CyclePage handles page navigation in both directions
func (m Model) CyclePage(direction string) Model {
	if direction == "right" {
		if m.PageIndex == len(m.Pages)-1 {
			m.PageIndex = 0
		} else {
			m.PageIndex++
		}
	} else if direction == "left" {
		if m.PageIndex == 0 {
			m.PageIndex = len(m.Pages) - 1
		} else {
			m.PageIndex--
		}
	}
	return m
}

// SaturateContent renders the content for the current page
func SaturateContent(m Model, viewportWidth int) string {
	// Get style path from config or use default
	stylePath := "assets/MDStyle.json"
	if GlobalConfig.Style.MDPath != "" {
	    stylePath = os.ExpandEnv(GlobalConfig.Style.MDPath)
	}
    
	// Create renderer with custom style from JSON file
	renderer, err := glamour.NewTermRenderer(
	    glamour.WithStylePath(stylePath),
	    glamour.WithWordWrap(viewportWidth-20),
	)
	internal.Check(err, "Glamour renderer creation", false)
    
	var content string
	switch m.PageIndex {
	case 0: // main.go
	    content = internal.GetMarkdown("main")
	case 1: // projects.cc 
	    content = internal.OpenProject(m.OpenProject, m.Projects, viewportWidth)
	case 2: // about.rs
	    content = internal.GetMarkdown("about")
	case 3: // contact.sh
	    content = internal.GetMarkdown("contact")
	}
    
	rendered, err := renderer.Render(content)
	internal.Check(err, "Content render", false)
	return rendered
    }

// CalculateNavItemPosition returns the position of navigation items
func (m Model) CalculateNavItemPosition(title string) (int, int) {
	startingPoint := m.Viewport.Width/2 - 57
	width, height := internal.CalculateNavItemSize(title)
	return startingPoint + width, height
}
