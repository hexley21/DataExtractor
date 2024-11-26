package multi_select

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hexley21/data_extractor/pkg/config"
)

type model struct {
	options []string
	choice  map[int]struct{}
	cursor  int
	colors  config.MultiSelect
	exited bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func initialModel(options []string, choice map[int]struct{}, colors config.MultiSelect) model {
	return model{
		options: options,
		choice:  choice,
		colors:  colors,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.exited = true
			return m, tea.Quit

		case "y":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}

		case "enter", " ":
			_, ok := m.choice[m.cursor]
			if ok {
				delete(m.choice, m.cursor)
			} else {
				m.choice[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	focusedStyle := getFocusedStyle(m.colors.Focused)
	selectedStyle := getSelectedStyle(m.colors.Selected)

	var sb strings.Builder
	sb.WriteString("Use Up/Down to navigate, Space/Enter to select, Q to quit.\n")

	for i, option := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = focusedStyle.Render(">")
			option = selectedStyle.Render(option)

		}

		checked := " "
		if _, ok := m.choice[i]; ok {
			checked = focusedStyle.Render("*")
		}

		title := focusedStyle.Render(option)

		sb.WriteString(cursor)
		sb.WriteString(" [")
		sb.WriteString(checked)
		sb.WriteString("] ")
		sb.WriteString(title)
		sb.WriteString("\n")
	}

	sb.WriteString("Press ")
	sb.WriteString(focusedStyle.Render("y"))
	sb.WriteString(" to confirm choice.\n")

	return sb.String()
}

func getFocusedStyle(color string) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Bold(true)
}

func getSelectedStyle(color string) lipgloss.Style {
	return lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color(color)).Bold(true)
}

func DisplayChecklist(keys []string, colors config.MultiSelect) ([]string, error) {
	selected := make(map[int]struct{})

	model := initialModel(keys, selected, colors)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		return nil, fmt.Errorf("failed to start checklist UI: %w", err)
	}

	if model.exited || len(selected) == 0 {
		return nil, nil
	}

	var result []string
	for s := range selected {
		result = append(result, keys[s])
	}

	return result, nil
}
