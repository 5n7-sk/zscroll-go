package zscroll

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Scroller struct {
	Text string

	AfterText      string
	BeforeText     string
	Delay          float64
	Length         int
	NewLine        bool
	PaddingText    string
	Reverse        bool
	Scroll         bool
	ScrollRate     int
	Timeout        int
	UpdateCommand  string
	UpdateInterval int

	index int
}

type ScrollerOptions struct {
	AfterText      string
	BeforeText     string
	Delay          float64
	Length         int
	NewLine        bool
	PaddingText    string
	Reverse        bool
	Scroll         bool
	ScrollRate     int
	Timeout        int
	UpdateCommand  string
	UpdateInterval int
}

func NewScroller(text string, opt ScrollerOptions) (*Scroller, error) {
	if len(text) == 0 {
		return nil, fmt.Errorf("text is empty")
	}

	if opt.ScrollRate <= 0 {
		return nil, fmt.Errorf("non-positive scroll rate: %d", opt.ScrollRate)
	}

	if opt.UpdateInterval <= 0 {
		return nil, fmt.Errorf("non-positive interval: %d", opt.UpdateInterval)
	}

	s := &Scroller{
		Text: text,

		AfterText:      opt.AfterText,
		BeforeText:     opt.BeforeText,
		Delay:          opt.Delay,
		Length:         opt.Length,
		NewLine:        opt.NewLine,
		PaddingText:    opt.PaddingText,
		Reverse:        opt.Reverse,
		Scroll:         opt.Scroll,
		ScrollRate:     opt.ScrollRate,
		Timeout:        opt.Timeout,
		UpdateCommand:  opt.UpdateCommand,
		UpdateInterval: opt.UpdateInterval,
	}
	return s, nil
}

func (s *Scroller) getDisplayLength() int {
	textLength := len(s.Text)
	if s.Length < 0 {
		return textLength
	}
	if textLength > s.Length {
		return s.Length
	}
	return textLength
}

func (s *Scroller) CurrentString() string {
	var (
		mainText string
		tempText string = s.Text + s.PaddingText // concat padding text
	)

	if displayLength := s.getDisplayLength(); s.index < len(tempText)-displayLength+1 {
		mainText = tempText[s.index : s.index+displayLength]
	} else {
		head := tempText[s.index:]
		mainText = head + tempText[:displayLength-len(head)]
	}
	return s.BeforeText + mainText + s.AfterText
}

func (s *Scroller) print() {
	if !s.NewLine {
		fmt.Printf("\r%s", s.CurrentString())
		os.Stdout.Sync()
	} else {
		fmt.Println(s.CurrentString())
	}
}

func (s *Scroller) step() *Scroller {
	if !s.Scroll {
		return s
	}

	length := len(s.Text) + len(s.PaddingText)
	if !s.Reverse {
		s.index = (s.index + 1) % length
	} else {
		s.index = (s.index - 1 + length) % length
	}
	return s
}

func (s *Scroller) update() error {
	out, err := exec.Command("/bin/sh", "-c", s.UpdateCommand).Output()
	if err != nil {
		return err
	}

	if text := strings.TrimSuffix(string(out), "\n"); text != s.Text {
		s.Text = text
		s.index = 0
	}
	return nil
}

func (s *Scroller) Run() error {
	ctx := context.Background()
	if s.Timeout >= 0 {
		c, cancel := context.WithTimeout(ctx, time.Duration(s.Timeout)*time.Second)
		defer cancel()
		ctx = c
	}

	ticker := time.NewTicker(time.Duration(s.UpdateInterval) * time.Second)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if s.UpdateCommand == "" {
				continue
			}

			if err := s.update(); err != nil {
				return err
			}
		default:
			s.print()
			time.Sleep(time.Duration(s.Delay*1000) * time.Millisecond)
			for i := 0; i < s.ScrollRate; i++ {
				s.step()
			}
		}
	}
}
