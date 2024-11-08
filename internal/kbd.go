package internal

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Navigate        key.Binding
	Up              key.Binding
	Down            key.Binding
	Left            key.Binding
	Right           key.Binding
	LCycle          key.Binding
	RCycle          key.Binding
	Enter           key.Binding
	Back            key.Binding
	Help            key.Binding
	Quit            key.Binding
	FullHelpBinding key.Binding
}

// ShortHelp returns keybindings to show in the minimal help view
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},    // First column
		{k.Left, k.Right}, // Second column
		{k.Enter, k.Back}, // Third column
		{k.Help, k.Quit},  // Fourth column
	}
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
	FullHelpBinding: key.NewBinding(
		key.WithKeys("shift+?", "f1"),
		key.WithHelp("f1", "full help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
