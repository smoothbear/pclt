package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}

	textInput textinput.Model
	step     int
}

type selector map[int]string

var results map[int]string
var steps = map[int]selector{
	0: {
		0: "Maven Project",
		1: "Gradle Project",
	},
	1: {
		0: "Java",
		1: "Kotlin",
		2: "Groovy",
	},
	2: {
		0: "2.4.4.RELEASE",
		1: "2.5.0.BUILD-SNAPSHOT",
		2: "2.5.0.M3",
		3: "2.4.5.BUILD-SNAPSHOT",
		4: "2.3.10.BUILD-SNAPSHOT",
		5: "2.3.9.RELEASE",
	},
	/*
		Steps that setting names.
	*/
	8: {
		0: "jar",
		1: "war",
	},
	9: {
		0: "16",
		1: "11",
		2: "8",
	},
}

var guide = map[int]string{
	1: "Select project type",
	2: "Please select a language to use",
	3: "Please enter your group name",
	4: "Please enter your group artifact name",
	5: "Please enter your project name",
	6: "Please enter your description",
	7: "Please enter your package name",
	8: "Please select packaging ways to use",
	9: "Please select java version to use",
	10: "Please select dependencies to use",
}

var dValue = map[int]string {
	3: "com.example",
	4: "demo",
	5: "demo",
	6: "Demo project for Spring Boot",
	7: "com.example.demo",
}

func (m model) Init() tea.Cmd {
	results = map[int]string{}
	return textinput.Blink
}

func (m *model) initMenu() {
	ti := textinput.NewModel()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	m.choices = []string{}
	m.textInput = ti
	m.selected = map[int]struct{}{}
	for _, choice := range steps[m.step] {
		m.choices = append(m.choices, choice)
	}
}


func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
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

			print(m.step)
			if m.step > 2 && m.step < 8 {
				results[m.step] = m.textInput.Value()
				if results[m.step] == "" {
					results[m.step] = dValue[m.step]
				}

				delete(m.selected, m.cursor)
			} else if m.step > 10 {
				for i, result := range results {
					print(i, ":", result, "\n")
				}
			} else {
				results[m.step] = steps[m.step][m.cursor]
				delete(m.selected, m.cursor)
			}

			m.step++
			m.initMenu()
		}

		m.textInput, cmd = m.textInput.Update(msg)

		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	var s string
	s += guide[m.step]
	s += "\n\n"

	if m.step > 2 && m.step < 8 {
		s += m.textInput.View()
	} else {
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
	}

	s += "\nPress ctrl+c to quit.\n"

	return s
}
