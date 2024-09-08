package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
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

func (m *Menu) GetString(prompt string) string {
	fmt.Print(prompt)
	m.Scanner.Scan()
	return strings.Trim(m.Scanner.Text(), " ")
}

func (m *Menu) GetInt(prompt string) (int, error) {
	str := m.GetString(prompt)
	int, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return int, nil
}

func (m *Menu) GetFloat(prompt string) (float64, error) {
	str := m.GetString(prompt)
	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return float, nil
}
