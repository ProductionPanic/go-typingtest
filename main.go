package main

import (
	"flag"
	"fmt"
	"typingtest/api"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
)

var minLengthArg = flag.Int("min", 80, "Minimum length of the quote")
var maxLengthArg = flag.Int("max", -1, "Maximum length of the quote")

func GetRandomTarget(min, max int) string {
	quote, err := api.GetRandomQuote(api.ApiArguments{
		MaxLength: max,
		MinLength: min,
	})
	if err != nil {
		panic(err)
	}
	return quote.Content
}

func main() {
	flag.Parse()

	w, h, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	random_target := GetRandomTarget(*minLengthArg, *maxLengthArg)

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

	m, err := p.Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	finished_model := m.(CliTypingGameModel)

	fmt.Printf("WPM: %s\n", finished_model.GetWpmString())
	fmt.Printf("Time: %s\n", finished_model.GetTimeString())
}
