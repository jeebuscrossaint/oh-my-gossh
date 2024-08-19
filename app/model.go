package app

import (
	frontend "oh-my-gossh/internal"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
)

type Model struct {
	PageIndex    int
	Pages        []string
	Projects     []string
	ProjectOpen  bool
	OpenProject  int
	ProjectView  string
	ClickCounter int
	Viewport     viewport.Model
	List         list.Model
	Content      string
	Keys         frontend.KeyMap
	Help         help.Model
	Ready        bool
}
