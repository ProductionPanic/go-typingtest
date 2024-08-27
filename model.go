package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"math"
	"time"
)

const (
	DEFAULT_TEXT_COLOR = lg.Color("#FDFFFC")
	TYPED_TEXT_COLOR   = lg.Color("#4CB944")
	CURSOR_TEXT_COLOR  = lg.Color("#FFA62B")
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
		// is playing
		is_playing := m.CurrentIndex < len(m.TargetString)
		if is_playing {
			switch msg.String() {
			case string(m.TargetString[m.CurrentIndex]):
				m.CurrentString += string(m.TargetString[m.CurrentIndex])
				m.CurrentIndex++
			case "ctrl+c":
				return m, tea.Quit
			}
		} else {
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

	optimal_width := 50
	width := int(math.Min(float64(m.term_width), float64(optimal_width)))

	out := lg.JoinVertical(
		lg.Left,
		m.GetCurrentString(),
		"",
		m.GetTimeString(),
	)

	out = lg.NewStyle().Width(width).Render(out)
	return lg.Place(m.term_width, m.term_height, lg.Center, lg.Center, out)
}

func (m CliTypingGameModel) GetCurrentString() string {
	s := []string{}
	currentStringStyle := lg.NewStyle().Underline(false).Foreground(DEFAULT_TEXT_COLOR)
	targetStringStyle := lg.NewStyle().Underline(false).Foreground(TYPED_TEXT_COLOR)
	nextCharStyle := lg.NewStyle().Underline(false).Foreground(CURSOR_TEXT_COLOR)

	for i, c := range m.TargetString {
		if i < m.CurrentIndex {
			s = append(s, targetStringStyle.Render(string(c)))
		} else if i == m.CurrentIndex {
			s = append(s, nextCharStyle.Render(string(c)))
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
	out := fmt.Sprintf("%.2f Seconds", timeTaken)
	return lg.NewStyle().Foreground(DEFAULT_TEXT_COLOR).Render(out)
}

func (m CliTypingGameModel) GetWpmString() string {
	currentTime := time.Now().UnixMilli()
	timeTaken := float64(currentTime-m.StartTimestamp) / 1000
	wpm := float64(m.CurrentIndex) / timeTaken * 60 / 5
	out := fmt.Sprintf("%.2f WPM", wpm)
	return lg.NewStyle().Foreground(DEFAULT_TEXT_COLOR).Render(out)
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
		lg.NewStyle().Foreground(DEFAULT_TEXT_COLOR).Render("Game Over!"),
		"",
		m.GetTimeString(),
		m.GetWpmString(),
	)
	return lg.Place(m.term_width, m.term_height, lg.Center, lg.Center, out)
}
