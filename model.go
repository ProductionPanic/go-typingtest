package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"math"
	"time"
)

type CliTypingGameModel struct {
	TargetString   string
	CurrentString  string
	CurrentIndex   int
	StartTimestamp int64
	EndTimestamp   int64

	term_width  int
	term_height int
}

func (m CliTypingGameModel) Init() tea.Cmd {
	return nil
}

func (m CliTypingGameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	started := false
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.CurrentIndex == 0 {
			m.StartTimestamp = time.Now().UnixMilli()
			started = true
		}
		if m.CurrentIndex == len(m.TargetString) {
			m.EndTimestamp = time.Now().UnixMilli()
		}
		switch msg.String() {
		case string(m.TargetString[m.CurrentIndex]):
			m.CurrentString += string(m.TargetString[m.CurrentIndex])
			m.CurrentIndex++
		case "ctrl+c":
			return m, tea.Quit
		}
	case tickMsg:
		if m.CurrentIndex < len(m.TargetString) {
			return m, GameTick()
		}
	}
	if started {
		return m, GameTick()
	}
	return m, nil
}

func GameTick() tea.Cmd {
	return tea.Tick(time.Millisecond*10, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

type tickMsg struct{}

func (m CliTypingGameModel) View() string {
	if m.HasEnded() {
		return m.ViewEnded()
	}

	out := lg.JoinVertical(
		lg.Left,
		m.GetCurrentString(),
		"",
		m.GetTimeString(),
	)

	optimal_width := 50
	width := math.Min(float64(m.term_width), float64(optimal_width))

	out = lg.NewStyle().Width(int(width)).Render(out)
	return lg.Place(m.term_width, m.term_height, lg.Center, lg.Center, out)
}

func (m CliTypingGameModel) GetCurrentString() string {
	s := []string{}
	currentStringStyle := lg.NewStyle().Underline(false).Foreground(lg.Color("#ffffff"))
	targetStringStyle := lg.NewStyle().Underline(false).Foreground(lg.Color("#ff0033"))

	for i, c := range m.TargetString {
		if i < m.CurrentIndex {
			s = append(s, targetStringStyle.Render(string(c)))
		} else if i == m.CurrentIndex {
			s = append(s, currentStringStyle.UnsetUnderline().Underline(true).Render(string(c)))
		} else {
			s = append(s, currentStringStyle.Render(string(c)))
		}
	}
	return lg.JoinHorizontal(lg.Left, s...)
}

func (m CliTypingGameModel) GetTimeString() string {
	if !m.HasStarted() {
		return ""
	}
	currentTime := time.Now().UnixMilli()
	timeTaken := float64(currentTime-m.StartTimestamp) / 1000
	return fmt.Sprintf("%.2f Seconds", timeTaken)
}

func (m CliTypingGameModel) GetWpmString() string {
	currentTime := time.Now().UnixMilli()
	timeTaken := float64(currentTime-m.StartTimestamp) / 1000
	wpm := float64(m.CurrentIndex) / timeTaken * 60 / 5
	return fmt.Sprintf("%.2f WPM", wpm)
}

func (m CliTypingGameModel) HasStarted() bool {
	return m.CurrentIndex > 0
}

func (m CliTypingGameModel) HasEnded() bool {
	return m.CurrentIndex == len(m.TargetString)
}

func (m CliTypingGameModel) ViewEnded() string {
	out := lg.JoinVertical(
		lg.Left,
		lg.NewStyle().Foreground(lg.Color("#ffffff")).Render("Game Over!"),
		"",
		lg.NewStyle().Foreground(lg.Color("#ffffff")).Render(m.GetTimeString()),
		lg.NewStyle().Foreground(lg.Color("#ffffff")).Render(m.GetWpmString()),
	)
	return lg.Place(m.term_width, m.term_height, lg.Center, lg.Center, out)
}
