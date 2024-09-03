package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Menu struct {
	Prompt  string
	Options map[string]Option
	active  bool
	Scanner *bufio.Scanner
}

type Option struct {
	Description string
	Function    func()
}

func (m *Menu) AddOption(key, description string, function func()) error {
	option, ok := m.Options[key]
	if ok {
		return errors.New("key already exists")
	}
	option.Description = description
	option.Function = function
	m.Options[key] = option
	return nil
}

func (m *Menu) ShowOptions() {
	for key, option := range m.Options {
		fmt.Printf("%s: %s\n", key, option.Description)
	}
}

func (m *Menu) Start() {
	m.active = true
	fmt.Println("Press ? to show options")
	fmt.Printf("\n[%s] >> ", m.Prompt)
	for m.Scanner.Scan() {
		text := strings.Trim(m.Scanner.Text(), " ")
		text = strings.ToLower(text)
		if option, ok := m.Options[text]; ok {
			option.Function()
		}
		if !m.active {
			break
		}
		fmt.Printf("\n[%s] >> ", m.Prompt)
	}
}

func NewMenu(prompt string) *Menu {
	m := Menu{}
	m.Prompt = prompt
	m.Options = make(map[string]Option)
	m.AddOption("?", "show options", func() {
		m.ShowOptions()
	})
	m.AddOption("x", "close menu", func() {
		m.active = false
	})
	m.Scanner = bufio.NewScanner(os.Stdin)
	return &m
}
