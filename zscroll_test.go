package zscroll

import (
	"testing"
)

func TestScroller_getDisplayLength(t *testing.T) {
	type fields struct {
		Text   string
		Length int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			fields: fields{
				Text:   "0123456789",
				Length: 5,
			},
			want: 5,
		},
		{
			fields: fields{
				Text:   "0123456789",
				Length: 20,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scroller{
				Text:   tt.fields.Text,
				Length: tt.fields.Length,
			}
			if got := s.getDisplayLength(); got != tt.want {
				t.Errorf("Scroller.getDisplayLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScroller_CurrentString(t *testing.T) {
	type fields struct {
		Text        string
		AfterText   string
		BeforeText  string
		Length      int
		PaddingText string
		index       int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{
				Text:        "0123456789",
				AfterText:   "<<<",
				BeforeText:  ">>>",
				Length:      10,
				PaddingText: "~~~",
				index:       0,
			},
			want: ">>>0123456789<<<",
		},
		{
			fields: fields{
				Text:        "0123456789",
				AfterText:   "<<<",
				BeforeText:  ">>>",
				Length:      10,
				PaddingText: "~~~",
				index:       5,
			},
			want: ">>>56789~~~01<<<",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scroller{
				Text:        tt.fields.Text,
				AfterText:   tt.fields.AfterText,
				BeforeText:  tt.fields.BeforeText,
				Length:      tt.fields.Length,
				PaddingText: tt.fields.PaddingText,
				index:       tt.fields.index,
			}
			if got := s.CurrentString(); got != tt.want {
				t.Errorf("Scroller.CurrentString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScroller_step(t *testing.T) {
	type fields struct {
		Text    string
		Reverse bool
		Scroll  bool
		index   int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			fields: fields{
				Text:    "0123456789",
				Reverse: false,
				Scroll:  true,
				index:   9,
			},
			want: 2,
		},
		{
			fields: fields{
				Text:    "0123456789",
				Reverse: true,
				Scroll:  true,
				index:   0,
			},
			want: 7,
		},
		{
			fields: fields{
				Text:    "0123456789",
				Reverse: false,
				Scroll:  false,
				index:   0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scroller{
				Text:    tt.fields.Text,
				Reverse: tt.fields.Reverse,
				Scroll:  tt.fields.Scroll,
				index:   tt.fields.index,
			}
			if s.step().step().step(); s.index != tt.want {
				t.Errorf("Scroller.index = %v, want %v", s, tt.want)
			}
		})
	}
}
