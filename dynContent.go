package signs

import (
	"fmt"
	"strings"

	"github.com/ev2-1/mt-multiserver-playertools"
)

type DynContent interface {
	Evaluate(string, *SignPos) string
}

type PlayerCnt struct {
	Srv string
}

func (pc *PlayerCnt) Evaluate(text string, pos *SignPos) string {
	return fmt.Sprintf("%d", playerTools.ServerPlayers(pc.Srv))
}

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

type Padding struct {
	Prepend bool
	Append  bool
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

type Text struct {
	text string
}

func (t *Text) Evaluate(text string, pos *SignPos) string {
	return t.text
}


