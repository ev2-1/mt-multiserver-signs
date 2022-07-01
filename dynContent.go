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
	Filler rune
	Length int

	Content DynContent
}

func (ce *Center) Evaluate(text string, pos *SignPos) string {
	c := ce.Content.Evaluate(text, pos)
	l := len(c)
	
	var p1, p2 string
	if len(c) % 2 == 0 {
		p1 = strings.Repeat(string(ce.Filler), ce.Length/2 -  l/2)
		p2 = p1
	} else {
		p1 = strings.Repeat(string(ce.Filler), ce.Length/2 -  l/2)
		p2 = strings.Repeat(string(ce.Filler), ce.Length/2 - (l/2)+1)		
	}

	return p1 + c + p2
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
