package app

import (
	"oh-my-gossh/internal"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
)

func Tui() (tea.Model, []tea.ProgramOption) {

	err := InitConfig()
	internal.Error(err, "TOML is broken, check your config file", true)

	tabs := []string{"main.go", "projects.cc", "about.rs", "contact.sh"}

	projects := []string{}
	itemizedProjects := []list.Item{}

	initModel := Model{
		PageIndex:   0,
		Pages:       tabs,
		Projects:    projects,
		ProjectOpen: false,
		List:        list.New(itemizedProjects, list.NewDefaultDelegate(), 0, 0),
		Keys:        internal.DefaultKeyMap,
		Help:        help.New(),
	}

	initModel.List.InfiniteScrolling = true
	initModel.List.DisableQuitKeybindings()
	initModel.List.SetFilteringEnabled(false)
	initModel.List.SetShowHelp(false)
	initModel.List.SetShowTitle(false)

	return initModel, []tea.ProgramOption{tea.WithAltScreen(), tea.WithMouseCellMotion()}

}

func TuiSSH(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	initModel, programOptions := Tui()
	return initModel, programOptions
}
