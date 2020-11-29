package zscroll

import (
	"context"
	"fmt"
	"os"
	"time"
)

type Scroller struct {
	Text string

	AfterText   string
	BeforeText  string
	Delay       float64
	Length      int
	NewLine     bool
	PaddingText string
	Reverse     bool
	Scroll      bool
	Timeout     int

	index int
}

type ScrollerOptions struct {
	AfterText   string
	BeforeText  string
	Delay       float64
	Length      int
	NewLine     bool
	PaddingText string
	Reverse     bool
	Scroll      bool
	Timeout     int
}

func NewScroller(text string, opt ScrollerOptions) (*Scroller, error) {
	if len(text) == 0 {
		return nil, fmt.Errorf("text is empty")
	}

	s := &Scroller{
		Text: text,

		AfterText:   opt.AfterText,
		BeforeText:  opt.BeforeText,
		Delay:       opt.Delay,
		Length:      opt.Length,
		NewLine:     opt.NewLine,
		PaddingText: opt.PaddingText,
		Reverse:     opt.Reverse,
		Scroll:      opt.Scroll,
		Timeout:     opt.Timeout,
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

func (s *Scroller) Run() error {
	ctx := context.Background()
	if s.Timeout >= 0 {
		c, cancel := context.WithTimeout(ctx, time.Duration(s.Timeout)*time.Second)
		defer cancel()
		ctx = c
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			s.print()
			time.Sleep(time.Duration(s.Delay*1000) * time.Millisecond)
			s.step()
		}
	}
}
