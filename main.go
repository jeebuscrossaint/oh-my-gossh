package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const (
	host = "0.0.0.0"
	port = "19"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	server, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			activeterm.Middleware(), // Bubble Tea apps usually require a PTY.
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("Could not start server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = server.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := server.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not stop server", "error", err)
	}
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pages := []string{"~", "$ whoami", "projects.cc", "find me.js", "funny"}

	/*projects := []string{"AEV-Software", "Constrict", "SSHowcase", "Yavafetch", "mleko-czekoladowe-os", "jeebuscrossaint.github.io", "dotfiles"}
	itemizedProjects := []list.Item{
		item{title: "Alset Solar CyberSedan Software", desc: "Full stack system behind the FAUHS AEV solar car."},
		item{title: "Constrict", desc: "A simple, fast, and easy to use build system for any language."},
		item{title: "SSHowcase", desc: "A terminal user interface for my portfolio."},
		item{title: "Yavafetch", desc: "A simple system information tool written in JavaScript. Works only on windows."},
		item{title: "mleko-czekoladowe-os", desc: "A simple operating system written in Rust and Assembly."},
		item{title: "jeebuscrossaint.github.io", desc: "My personal website. It's a work in progress and where I store my static documentation pages."},
		item{title: "dotfiles", desc: "My personal dotfiles. I use them to configure my system and make it look nice like rice."},
	}*/

	projects, err := loadProjects()
	if err != nil {
		log.Error("Could not load projects", "error", err)
	}

	var itemizedProjects []list.Item
	for _, project := range projects {
		itemizedProjects = append(itemizedProjects, item{title: project.Name, desc: project.Description})
	}

	clientIP := s.RemoteAddr().String()
	if host, _, err := net.SplitHostPort(clientIP); err == nil {
		clientIP = host
	}

	initialModel := model{
		pageIndex:    0,
		pages:        pages,
		projects:     projects,
		projectOpen:  false,
		list:         list.New(itemizedProjects, list.NewDefaultDelegate(), 0, 0),
		keys:         DefaultKeyMap,
		help:         help.New(),
		quote:        getRandomQuote(),
		clientIP:     clientIP,
		cmatrixState: "",
	}

	initialModel.list.InfiniteScrolling = true
	initialModel.list.DisableQuitKeybindings()
	initialModel.list.SetFilteringEnabled(false)
	initialModel.list.SetShowHelp(false)
	initialModel.list.SetShowTitle(false)
	initialModel.list.Title = "Hit ENTER for more details"

	return initialModel, []tea.ProgramOption{tea.WithAltScreen(), tea.WithMouseCellMotion()}
}

func getRandomQuote() string {
	quotes, err := os.ReadFile("assets/quotes.txt")
	if err != nil {
		return "Error reading quotes file."
	}
	lines := strings.Split(string(quotes), "\n")
	// Remove empty lines
	var nonEmptyLines []string
	for _, line := range lines {
		if line != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	if len(nonEmptyLines) == 0 {
		return "No quotes available."
	}
	return nonEmptyLines[rand.Intn(len(nonEmptyLines))]
}

// my name

const ASCII = `
 █████╗ ███╗   ███╗ █████╗ ██████╗ ███╗   ██╗ █████╗ ████████╗██╗  ██╗    ██████╗  █████╗ ████████╗███████╗██╗     
██╔══██╗████╗ ████║██╔══██╗██╔══██╗████╗  ██║██╔══██╗╚══██╔══╝██║  ██║    ██╔══██╗██╔══██╗╚══██╔══╝██╔════╝██║     
███████║██╔████╔██║███████║██████╔╝██╔██╗ ██║███████║   ██║   ███████║    ██████╔╝███████║   ██║   █████╗  ██║     
██╔══██║██║╚██╔╝██║██╔══██║██╔══██╗██║╚██╗██║██╔══██║   ██║   ██╔══██║    ██╔═══╝ ██╔══██║   ██║   ██╔══╝  ██║     
██║  ██║██║ ╚═╝ ██║██║  ██║██║  ██║██║ ╚████║██║  ██║   ██║   ██║  ██║    ██║     ██║  ██║   ██║   ███████╗███████╗
╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝    ╚═╝     ╚═╝  ╚═╝   ╚═╝   ╚══════╝╚══════╝
`

type model struct {
	pageIndex int
	pages     []string
	//projects     []string
	projects     []Project
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
	quote        string
	clientIP     string
	cmatrixState string
}

func check(e error, check string) {
	if e != nil {
		fmt.Printf("Error running program - In %v: %v", check, e)
	}
}

var termHeight int

var (
	navStyle           = lipgloss.NewStyle().Margin(1, 0).Padding(0, 2)
	listStyle          = lipgloss.NewStyle().Padding(1, 2)
	bubbleLettersStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#b9b9b9"))
	// For nav text
	activePageStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#b9b9b9")).Bold(true).PaddingLeft(2).PaddingRight(4)
	inactivePageStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).PaddingLeft(4).PaddingRight(4)

	// Border styles
	borderTitleStyle = func() lipgloss.Style {
		b := lipgloss.HiddenBorder()
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()
	borderInfoStyle = func() lipgloss.Style {
		b := lipgloss.HiddenBorder()
		return borderTitleStyle.BorderStyle(b)
	}()
)

// Bubbletea key mapping (Struct + defaults)
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

// Bubbletea help component full & short displays
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Navigate, k.RCycle, k.Enter, k.Quit, k.Help}
}
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.RCycle, k.Enter},
		{k.Up, k.Down},
		{k.LCycle, k.Back},
		{k.Help, k.Quit},
	}
}

// Projects list setup
type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func openProject(selectedProject int, projects []Project, viewportWidth int) string {
	if selectedProject < 0 || selectedProject >= len(projects) {
		return "Invalid project selection"
	}

	project := projects[selectedProject]
	rawProjectPageTemplate, _ := glamour.NewTermRenderer(
		glamour.WithStylePath("assets/MDStyle.json"),
		glamour.WithWordWrap(viewportWidth-20),
	)

	projectPage, err := rawProjectPageTemplate.Render(getMarkdown("projects/" + project.Name))
	if err != nil {
		return fmt.Sprintf("Error rendering project %s: %v", project.Name, err)
	}

	return projectPage
}

// Function to read and return markdown file data for each page
func getMarkdown(filename string) string {
	fileData, err := os.ReadFile(filepath.Join("assets", "markdown", filename+".md"))
	if err != nil {
		return fmt.Sprintf("Error reading markdown file: %v", err)
	}

	return string(fileData)
}

// Function to get the proper content according to each page
func saturateContent(m model, viewportWidth int) string {
	// Checks which page the user is on and renders it accordingly
	var content string
	var err error

	rawMarkdownPageTemplate, _ := glamour.NewTermRenderer(
		glamour.WithStylePath("assets/MDStyle.json"),
		// glamour.WithAutoStyle(), - For Light/Darkmode styling except I'd rather use my custom style
		glamour.WithWordWrap(viewportWidth-20),
	)

	switch m.pageIndex {
	case 0: // main.rs
		content, err = rawMarkdownPageTemplate.Render(getMarkdown("~"))
		check(err, "Gleam Markdown Render")
	case 1: // Whoami
		content, err = rawMarkdownPageTemplate.Render(getMarkdown("$ whoami"))
		check(err, "Gleam Markdown Render")
	case 3: // irc
		content, err = rawMarkdownPageTemplate.Render(getMarkdown("find me.js"))
		check(err, "Gleam Markdown Render")
	case 4: // funny
		content = fmt.Sprintf("Your IP: %s\n\n%s", m.clientIP, m.cmatrixState)
	}

	return content
}

// Bubbletea function to cycle each page (when tab is clicked, this function handles the update event)
func (m model) cyclePage(direction string) model {
	if m.pageIndex < len(m.pages) && direction == "right" {
		switch m.pageIndex {
		case len(m.pages) - 1:
			m.pageIndex = 0
			return m
		default:
			m.pageIndex++
			return m
		}
	} else if m.pageIndex >= 0 && direction == "left" {
		switch m.pageIndex {
		case 0:
			m.pageIndex = len(m.pages) - 1
			return m
		default:
			m.pageIndex--
			return m
		}
	} else {
		return m
	}
}

// Gets the location and size of each navigation menu button
// (this is hard coded as of now since I have no idea how to programmatically find a components location & size in the terminal)
func (m model) calculateNavItemPosition(title string) (int, int) {
	startingPoint := m.viewport.Width/2 - 57
	switch title {
	case "home":
		return startingPoint + 30, 9
	case "about":
		return startingPoint + 43, 9
	case "projects":
		return startingPoint + 58, 9
	case "contact":
		return startingPoint + 75, 9
	default:
		return 0, 0
	}
}
func calculateNavItemSize(title string) (int, int) {
	switch title {
	case "home":
		return 10, 2
	case "about":
		return 10, 2
	case "projects":
		return 13, 2
	case "contact":
		return 12, 2
	default:
		return 0, 0
	}
}

// Max function for viewport line length
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Page viewport header and footer render
func (m model) viewportHeader(pageTitle string) string {
	title := borderTitleStyle.Render(pageTitle)
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}
func (m model) viewportFooter() string {
	info := borderInfoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

// Empty init for now since there's not much hard logic
func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Amarnath's Portfolio TUI 😀"),
		tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
			return tea.TickMsg(t)
		}),
	)
}

type TickMsg struct {
	Time time.Time
	tag  int
	ID   int
}

// Bubbletea update/msg handling
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Commands to be returned for viewport updating
	var (
		viewportCMD     tea.Cmd
		projectsListCMD tea.Cmd
		cmds            []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.MouseMsg:
		switch tea.MouseAction(msg.Button) {
		case 1: // Mouse left click
			for i, title := range m.pages {
				x, y := m.calculateNavItemPosition(title)
				width, height := calculateNavItemSize(title)

				if msg.X >= x && msg.X <= x+width && msg.Y >= y && msg.Y <= y+height {
					m.pageIndex = i
					m.viewport.SetContent(saturateContent(m, m.viewport.Width))
					return m, nil
				} else if msg.Y >= termHeight-3 {
					m.help.ShowAll = !m.help.ShowAll
					return m, nil
				}
			}
			// This is a very lousy approach for making each item clickable but it's the only way I have time to do as of now...
			// This also causes the mouse support to break on pages past the first if pagination is necessary depending on terminal size
			if m.pageIndex == 2 && !m.projectOpen && msg.Y >= 16 && msg.Y < termHeight-3 {
				projectIndex := 0
				// BUG: for some reason after clicking down the list every once in a while it would enter the project MD even though it had only been clicked once then they all do that from that point on
				for i := 16; projectIndex <= len(m.projects)-1; i += 3 {
					if i <= msg.Y && msg.Y <= i+1 {
						if m.list.Index() == projectIndex {
							m.clickCounter++
						} else {
							m.clickCounter = 0
						}
						m.list.Select(projectIndex)
					} else {
						projectIndex++
					}
					if m.clickCounter >= 2 {
						m.clickCounter = 0
						m.projectOpen = true
						m.openProject = m.list.Index()
					}
				}
			}
		case tea.TickMsg:
			if m.pageIndex == 4 {
				m.cmatrixState = generateCMatrix(m.viewport.Width, m.viewport.Height-5) // -5 for header and IP display
				return m, tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
					return tea.TickMsg(t)
				})
			}
		case 4: // Scroll wheel up
			if m.pageIndex == 2 && !m.projectOpen {
				if m.list.Index() == 0 {
					m.list.Select(len(m.projects))
				} else {
					m.list.Select(m.list.Index() - 1)
				}
			}
		case 5: // Scroll wheel down
			if m.pageIndex == 2 && !m.projectOpen {
				if m.list.Index() == len(m.projects)-1 {
					m.list.Select(0)
				} else {
					m.list.Select(m.list.Index() + 1)
				}
			}
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, DefaultKeyMap.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, DefaultKeyMap.Navigate):
			break
		case key.Matches(msg, DefaultKeyMap.Up):
			break
		case key.Matches(msg, DefaultKeyMap.Down):
			break
		case key.Matches(msg, DefaultKeyMap.RCycle):
			cycled := m.cyclePage("right")
			cycled.viewport.SetContent(saturateContent(cycled, m.viewport.Width))
			return cycled, nil
		case key.Matches(msg, DefaultKeyMap.LCycle):
			cycled := m.cyclePage("left")
			m.viewport.SetContent(saturateContent(cycled, m.viewport.Width))
			return m.cyclePage("left"), nil
		case key.Matches(msg, DefaultKeyMap.Left):
			if m.pageIndex > 0 {
				m.pageIndex--
				m.viewport.SetContent(saturateContent(m, m.viewport.Width))
			}
			return m, nil
		case key.Matches(msg, DefaultKeyMap.Right):
			if m.pageIndex < len(m.pages)-1 {
				m.pageIndex++
				m.viewport.SetContent(saturateContent(m, m.viewport.Width))
			}
			return m, nil
		case key.Matches(msg, DefaultKeyMap.Enter):
			if m.pageIndex == 2 {
				m.projectOpen = true
				m.openProject = m.list.Index()
				m.viewport.GotoTop()
			}
		case key.Matches(msg, DefaultKeyMap.Back):
			if m.pageIndex == 2 {
				m.projectOpen = false
				m.list.Select(m.openProject)
				//m.viewport.GotoTop()
			}
		}
	case tea.WindowSizeMsg:
		// Set new terminal height for proper click areas
		termHeight = msg.Height
		// Setup for viewport sizing
		headerHeight := lipgloss.Height(m.viewportHeader(m.pages[m.pageIndex]))
		footerHeight := lipgloss.Height(m.viewportFooter())
		verticalMarginHeight := headerHeight + footerHeight
		// Project list size
		listMarginWidth, listMarginHeight := listStyle.GetFrameSize()
		m.list.SetSize(msg.Width-listMarginWidth, msg.Height-listMarginHeight-verticalMarginHeight-11)

		// Viewport creation & management
		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight-11)
			m.viewport.SetContent(saturateContent(m, m.viewport.Width))
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight - 11
		}
	}

	if m.pageIndex == 2 && m.projectOpen {
		m.viewport.SetContent(openProject(m.openProject, m.projects, m.viewport.Width))
	}
	// Handle keyboard and mouse events in the viewport
	// Gets viewport update command and map based on the message
	m.viewport, viewportCMD = m.viewport.Update(msg)
	// Update list depending on msg
	// Does the same as viewport but for projects list
	m.list, projectsListCMD = m.list.Update(msg)
	// Append all component commands to cmds
	cmds = append(cmds, viewportCMD, projectsListCMD)

	return m, tea.Batch(cmds...)
}

// Switch case with each page/TUI view
func (m model) View() string {

	// If viewport isn't ready it'll say welcome (this should only be able to happen during startup)
	if !m.ready {
		return "\n  Welcome to my portfolio! \n\n  Loading..."
	}

	nav := `` // Empty to be saturated soon

	// Render/create nav depending on page location
	for i, title := range m.pages {
		if i == m.pageIndex {
			// Highlight the active page
			nav += activePageStyle.Render("• " + title + " ")
		} else {
			nav += inactivePageStyle.Render(title + " ")
		}
	}

	m.content = m.viewportHeader(m.pages[m.pageIndex]) + m.viewport.View() + m.viewportFooter()
	if m.pageIndex == 2 {
		if !m.projectOpen {
			m.projectView = listStyle.Render(m.list.View())
		} else if m.projectOpen {
			m.projectView = m.viewport.View()
		}
		m.content = m.viewportHeader(m.pages[m.pageIndex]) + m.projectView + m.viewportFooter()
	}

	header := lipgloss.PlaceHorizontal(m.viewport.Width, lipgloss.Center, bubbleLettersStyle.Render(ASCII))
	quote := lipgloss.PlaceHorizontal(m.viewport.Width, lipgloss.Center, lipgloss.NewStyle().Italic(true).Render(m.quote))
	nav = lipgloss.PlaceHorizontal(m.viewport.Width, lipgloss.Center, navStyle.Render(nav))

	return header + "\n" + quote + "\n" + nav + m.content + navStyle.Render(m.help.View(m.keys))
}

type Project struct {
	Name        string
	Description string
}

func loadProjects() ([]Project, error) {
	var projects []Project
	files, err := os.ReadDir("assets/configs")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".conf") {
			content, err := os.ReadFile(filepath.Join("assets/configs", file.Name()))
			if err != nil {
				return nil, err
			}

			name := strings.TrimSuffix(file.Name(), ".conf")
			description := strings.TrimSpace(string(content))

			projects = append(projects, Project{
				Name:        name,
				Description: description,
			})
		}
	}

	return projects, nil
}

func generateCMatrix(width, height int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%^&*()"
	var builder strings.Builder
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if rand.Float32() < 0.1 {
				builder.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("green")).Render(string(chars[rand.Intn(len(chars))])))
			} else {
				builder.WriteString(" ")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
