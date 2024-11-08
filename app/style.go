package app

import (
    "github.com/charmbracelet/lipgloss"
)

var TermHeight int

var (
    // Navigation and layout styles
    NavStyle = lipgloss.NewStyle().
        Margin(1, 0).
        Padding(0, 2)

    ListStyle = lipgloss.NewStyle().
        Padding(1, 2)

    // Use color from config for header text
    BubbleLetterStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color(GlobalConfig.Color.Letter))

    // Use colors from config for page indicators
    ActivePageStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color(GlobalConfig.Color.Active))

    InactivePageStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color(GlobalConfig.Color.Inactive))

    // Border styles using active color for consistency
    BorderTitleStyle = func() lipgloss.Style {
        b := lipgloss.HiddenBorder()
        return lipgloss.NewStyle().
            BorderStyle(b).
            Padding(0, 1).
            Foreground(lipgloss.Color(GlobalConfig.Color.Active))
    }()

    BorderInfoStyle = func() lipgloss.Style {
        b := lipgloss.HiddenBorder()
        return BorderTitleStyle.Copy().
            BorderStyle(b).
            Foreground(lipgloss.Color(GlobalConfig.Color.Active))
    }()
)