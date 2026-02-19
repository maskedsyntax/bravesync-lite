package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type passwordModel struct {
	textInput textinput.Model
	password  string
	err       error
}

func NewPasswordModel(prompt string) passwordModel {
	ti := textinput.New()
	ti.Placeholder = prompt
	ti.EchoMode = textinput.EchoPassword
	ti.EchoCharacter = '•'
	ti.Focus()

	return passwordModel{
		textInput: ti,
	}
}

func (m passwordModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m passwordModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			m.password = m.textInput.Value()
			return m, tea.Quit
		}
	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m passwordModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n(esc to quit)\n",
		m.textInput.Placeholder,
		m.textInput.View(),
	)
}

// GetPasswordTUI prompts the user for a password in TUI mode.
func GetPasswordTUI(prompt string) (string, error) {
	m := NewPasswordModel(prompt)
	p := tea.NewProgram(m)
	finalModel, err := p.Run()
	if err != nil {
		return "", err
	}
	return finalModel.(passwordModel).password, nil
}
