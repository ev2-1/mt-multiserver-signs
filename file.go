package signs

import (
	"bufio"
	"strconv"
	"strings"
)

// ParseScanner parses a bufio.Scanner into a []*Sign
func ParseScanner(scan *bufio.Scanner) (s []*Sign) {
	var buffer = &Sign{}
	for scan.Scan() {
		text := scan.Text()

		// comments
		if len(text) != 0 && text[0] == '#' {
			continue
		}

		cmd := strings.SplitN(text, " ", 2)

		switch cmd[0] {
		case "pos":
			buffer.Pos = ParsePos(cmd[1])

		case "text":
			buffer.Text = strings.ReplaceAll(cmd[1], "\\n", "\n")

		case "color":
			buffer.Color = cmd[1]

		case "dyn":
			buffer.Dyn = append(buffer.Dyn, ParseDyn(cmd[1]))

		case "click":
			buffer.OnClick = ParseClick(cmd[1])

		case "end":
			if buffer.Text != "" {
				s = append(s, buffer)
				buffer = &Sign{}
			}
		}
	}

	return
}

// ParseClick parses a Stringrepresentation of a ClickEvent
func ParseClick(s string) ClickEvent {
	split := strings.SplitN(s, ":", 2)

	switch split[0] {
	case "Hop":
		return &Hop{
			Srv: split[1],
		}
	}

	return nil
}

// ParseDyn parses string into dyncontent Interface
func ParseDyn(s string) DynContent {
	split := strings.SplitN(s, ":", 2)

	switch split[0] {
	case "Padding":
		arg := strings.SplitN(split[1], ",", 4)
		l, _ := strconv.Atoi(arg[1])
		return &Padding{
			Prepend: arg[0] == "prepend",
			Length:  l,
			Filler:  []rune(arg[2])[0],
			Content: ParseDyn(arg[3]),
		}

	case "PlayerCnt":
		return &PlayerCnt{
			Srv: split[1],
		}

	case "Center":
		arg := strings.SplitN(split[1], ",", 3)
		l, _ := strconv.Atoi(arg[1])
		c := arg[2]
		return &Center{
			Filler:  []rune(arg[0])[0],
			Length: l,

			Content: ParseDyn(c),
		}

	case "Text":
		return &Text{
			Text: strings.ReplaceAll(split[1], "\\n", "\n"),
		}
	}

	return &Text{Text: ""}
}

// ParsePos parses a stringpos like:
// 2 10 -5@hub wall South
// into a SignPos
func ParsePos(s string) (pos *SignPos) {
	pos = &SignPos{}

	split := strings.SplitN(s, "@", 2)
	posStr := strings.SplitN(split[0], " ", 3)

	for k, _ := range pos.Pos {
		p, _ := strconv.Atoi(posStr[k])
		pos.Pos[k] = int16(p)
	}

	metaPos := strings.SplitN(split[1], " ", 3)
	pos.Server = metaPos[0]
	pos.Wall = metaPos[1] == "wall"
	pos.Rotation = ParseRotationString(metaPos[2])

	return
}

// ParseRotationString parses rotation string
func ParseRotationString(s string) Rotate {
	return rotationMap[strings.ToLower(s)]
}
