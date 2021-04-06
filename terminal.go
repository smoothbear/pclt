package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}

	step     int
}

var results map[int]string

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) initFirst() {
	m.selected = map[int]struct{}{}
	m.choices = []string{"Maven Project", "Gradle Project"}
	m.step = 0
}

func (m *model) initSecond() {
	m.selected = map[int]struct{}{}
	m.choices = []string{"Java", "Kotlin", "Groovy"}
	m.step = 1
}

func (m *model) initThird() {
	m.choices = []string{"2.4.4 (Recommended)", "2.5.0 (SNAPSHOT)", "2.5.0 (M3)", "2.4.5 (SNAPSHOT)", "2.3.10 (SNAPSHOT)", "2.3.9"}
	m.step = 2
}

func (m *model) initFourth() {
	m.choices = []string{"안녕하세요"}
	m.step = 3
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.selected[m.cursor] = struct{}{}
			switch m.step {
			case 0:
				m.initSecond()
				switch m.cursor {
				case 0:
					results = append(results, "type=maven-project")
				case 1:
					results = append(results, "type=gradle-project")
				}
				delete(m.selected, m.cursor)
			case 1:
				m.initThird()
				switch m.cursor {
				case 0:
					results = append(results, "&language=Java")
				case 1:
					results = append(results, "&language=Kotlin")
				case 2:
					results = append(results, "&language=Groovy")
				}
			case 2:
				m.initFourth()
				switch m.cursor {
				case 0:
					results = append(results, "&bootVersion=2.4.4.RELEASE")
				case 1:
					results = append(results, "&bootVersion=2.4.4.RELEASE")
				case 2:
					results = append(results, "&bootVersion=2.4.4.RELEASE")
				case 3:
					results = append(results, "&bootVersion=2.4.4.RELEASE")
				case 4:
					results = append(results, "&bootVersion=2.4.4.RELEASE")
				case 5:
					results = append(results, "&bootVersion=2.4.4.RELEASE")
				}
			default:
				fmt.Print(results)
			}

		}
	}

	return m, nil
}

func (m model) View() string {
	var s string
	switch m.step {
	case 0:
		s = "Select project type\n\n"
	case 1:
		s = ""
	}

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"
	return s
}
