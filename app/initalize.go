package app

import (
	"oh-my-gossh/internal"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) Init() tea.Cmd {
	return tea.SetWindowTitle(GlobalConfig.Title.Tab)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		viewportCmd tea.Cmd
		listCmd     tea.Cmd
		cmds        []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.MouseMsg:
		switch msg.Button {
		case tea.MouseButtonLeft:
			// Navigation menu clicks
			for i, title := range m.Pages {
				x, y := m.CalculateNavItemPosition(title)
				width, height := internal.CalculateNavItemSize(title)

				if msg.X >= x && msg.X <= x+width && msg.Y >= y && msg.Y <= y+height {
					m.PageIndex = i
					m.Viewport.SetContent(SaturateContent(m, m.Viewport.Width))
					return m, nil
				}
			}

			// Help toggle
			if msg.Y >= TermHeight-3 {
				m.Help.ShowAll = !m.Help.ShowAll
				return m, nil
			}

			// Project list clicks
			if m.PageIndex == 1 && !m.ProjectOpen {
				if msg.Y >= 16 && msg.Y < TermHeight-3 {
					projectIndex := 0
					for i := 16; projectIndex < len(m.Projects); i += 3 {
						if i <= msg.Y && msg.Y <= i+1 {
							if m.List.Index() == projectIndex {
								m.ClickCounter++
							} else {
								m.ClickCounter = 0
							}
							m.List.Select(projectIndex)
						} else {
							projectIndex++
						}
						if m.ClickCounter >= 2 {
							m.ClickCounter = 0
							m.ProjectOpen = true
							m.OpenProject = m.List.Index()
						}
					}
				}
			}

		case tea.MouseButtonWheelUp:
			if m.PageIndex == 1 && !m.ProjectOpen {
				if m.List.Index() == 0 {
					m.List.Select(len(m.Projects) - 1)
				} else {
					m.List.Select(m.List.Index() - 1)
				}
			}

		case tea.MouseButtonWheelDown:
			if m.PageIndex == 1 && !m.ProjectOpen {
				if m.List.Index() == len(m.Projects)-1 {
					m.List.Select(0)
				} else {
					m.List.Select(m.List.Index() + 1)
				}
			}
		}

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.Keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll
		case key.Matches(msg, m.Keys.RCycle):
			cycled := m.CyclePage("right")
			cycled.Viewport.SetContent(SaturateContent(cycled, m.Viewport.Width))
			return cycled, nil
		case key.Matches(msg, m.Keys.LCycle):
			cycled := m.CyclePage("left")
			cycled.Viewport.SetContent(SaturateContent(cycled, m.Viewport.Width))
			return cycled, nil
		case key.Matches(msg, m.Keys.Left):
			if m.PageIndex > 0 {
				m.PageIndex--
				m.Viewport.SetContent(SaturateContent(m, m.Viewport.Width))
			}
		case key.Matches(msg, m.Keys.Right):
			if m.PageIndex < len(m.Pages)-1 {
				m.PageIndex++
				m.Viewport.SetContent(SaturateContent(m, m.Viewport.Width))
			}
		case key.Matches(msg, m.Keys.Enter):
			if m.PageIndex == 1 {
				m.ProjectOpen = true
				m.OpenProject = m.List.Index()
				m.Viewport.GotoTop()
			}
		case key.Matches(msg, m.Keys.Back):
			if m.PageIndex == 1 {
				m.ProjectOpen = false
				m.List.Select(m.OpenProject)
			}
		}

	case tea.WindowSizeMsg:
		TermHeight = msg.Height

		headerHeight := lipgloss.Height(m.ViewportHeader(m.Pages[m.PageIndex]))
		footerHeight := lipgloss.Height(m.ViewportFooter())
		verticalMarginHeight := headerHeight + footerHeight

		listMarginWidth, listMarginHeight := ListStyle.GetFrameSize()
		m.List.SetSize(
			msg.Width-listMarginWidth,
			msg.Height-listMarginHeight-verticalMarginHeight-11,
		)

		if !m.Ready {
			m.Viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight-11)
			m.Viewport.SetContent(SaturateContent(m, m.Viewport.Width))
			m.Ready = true
		} else {
			m.Viewport.Width = msg.Width
			m.Viewport.Height = msg.Height - verticalMarginHeight - 11
		}
	}

	if m.PageIndex == 1 && m.ProjectOpen {
		m.Viewport.SetContent(internal.OpenProject(m.OpenProject, m.Projects, m.Viewport.Width))
	}

	m.Viewport, viewportCmd = m.Viewport.Update(msg)
	m.List, listCmd = m.List.Update(msg)
	cmds = append(cmds, viewportCmd, listCmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if !m.Ready {
		return "\n  Welcome..."
	}

	// Navigation menu
	nav := ""
	for i, title := range m.Pages {
		if i == m.PageIndex {
			nav += ActivePageStyle.Render("• " + title + " ")
		} else {
			nav += InactivePageStyle.Render(title + " ")
		}
	}

	// Main content rendering
	m.Content = m.ViewportHeader(m.Pages[m.PageIndex]) + m.Viewport.View() + m.ViewportFooter()

	if m.PageIndex == 1 {
		if !m.ProjectOpen {
			m.ProjectView = ListStyle.Render(m.List.View())
		} else {
			m.ProjectView = m.Viewport.View()
		}
		m.Content = m.ViewportHeader(m.Pages[m.PageIndex]) + m.ProjectView + m.ViewportFooter()
	}

	// Render header from ASCII art file
	header := strings.Repeat("\n", 2) + // Add 2 blank lines at top
        lipgloss.PlaceHorizontal(m.Viewport.Width, lipgloss.Center,
            BubbleLetterStyle.Render(GlobalConfig.Title.AsciiArt))

	// Add newline before subtitle for spacing
	subtitle := "\n" + lipgloss.PlaceHorizontal(m.Viewport.Width, lipgloss.Center,
		BubbleLetterStyle.Render(GlobalConfig.Title.Subtitle))

	// Add newline after subtitle for spacing from nav
	nav = "\n" + lipgloss.PlaceHorizontal(m.Viewport.Width, lipgloss.Center,
		NavStyle.Render(nav))

	// Combine all elements with proper spacing
	return header + subtitle + nav + m.Content + NavStyle.Render(m.Help.View(m.Keys))
}
