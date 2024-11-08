package app

import (
	"github.com/charmbracelet/lipgloss"
)

var TermHeight int

var (
	NavStyle          = lipgloss.NewStyle().Margin(1, 0).Padding(0, 2)
	ListStyle         = lipgloss.NewStyle().Padding(1, 2)
	BubbleLetterStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(GlobalConfig.Color.Letter))

	ActivePageStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color(GlobalConfig.Color.Active))
	InactivePageStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(GlobalConfig.Color.Inactive))

	BorderTitleStyle = func() lipgloss.Style {
		b := lipgloss.HiddenBorder()
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()
	BorderInfoStyle = func() lipgloss.Style {
		b := lipgloss.HiddenBorder()
		return BorderTitleStyle.BorderStyle(b)
	}()
)
