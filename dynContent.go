package signs

import (
	"fmt"
	"strings"

	"github.com/ev2-1/mt-multiserver-playertools"
)

type DynContent interface {
	Evaluate(string, *SignPos) string
}

// PlayerCnt is a `DynContent` replaced by the PlayerCount on Srv
type PlayerCnt struct {
	Srv string
}

func (pc *PlayerCnt) Evaluate(text string, pos *SignPos) string {
	return fmt.Sprintf("%d", playerTools.ServerPlayers(pc.Srv))
}

// Padding is a `DynContent`
type Center struct {
	Line   int
	Length int
	Rune   rune
	Sub    int
}

func (ce *Center) Evaluate(text string, pos *SignPos) string {
	p := (ce.Length - len(strings.Split(text, "\n")[ce.Line-1])) / 2

	return strings.Repeat(string(ce.Rune), p)
}

// Padding is a `DynContent`, padding Content to be `Length` long filleed up by `Filler` while either `Prepending` or else Appending
type Padding struct {
	Prepend bool
	Length  int
	Filler  rune
	Content DynContent
}

func (pa *Padding) Evaluate(text string, pos *SignPos) string {
	text = pa.Content.Evaluate(text, pos)

	if pa.Prepend {
		return strings.Repeat(string(pa.Filler), pa.Length-len(text)) + text
	} else {
		return text + strings.Repeat(string(pa.Filler), pa.Length-len(text))
	}
}

// Text is a `DynContent` thats static any always returns `Text`
type Text struct {
	Text string
}

func (t *Text) Evaluate(text string, pos *SignPos) string {
	return t.Text
}
