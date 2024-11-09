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

    tabs := GlobalConfig.Title.Pages
    if len(tabs) == 0 {
        // Fallback to defaults if not specified
        tabs = []string{"main.go", "projects.cc", "about.rs", "contact.sh"}
    }

    // Create slices to hold projects
    projects := []string{}
    itemizedProjects := []list.Item{}

    // Populate projects from GlobalConfig
    for _, project := range GlobalConfig.Projects {
        // Add project file path to projects slice
        projects = append(projects, project.File)
        
        // Add project as list item
        itemizedProjects = append(itemizedProjects, internal.Item{
            TitleText: project.Name,
            Desc:      project.About,
        })
    }

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
