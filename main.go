package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
	"typingtest/api"
)

func GetRandomTarget() string {
	quote, err := api.GetRandomQuote()
	if err != nil {
		panic(err)
	}
	return quote.Content
}

func main() {
	w, h, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}
	fmt.Println("Getting random quote...")
	random_target := GetRandomTarget()

	model := CliTypingGameModel{
		TargetString:   random_target,
		CurrentString:  "",
		CurrentIndex:   0,
		StartTimestamp: 0,
		EndTimestamp:   0,
		term_width:     w,
		term_height:    h,
	}

	p := tea.NewProgram(&model, tea.WithAltScreen())

	_, err = p.Run()
	if err != nil {
		panic(err)
	}
}
