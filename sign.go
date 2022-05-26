package signs

import (
	"bufio"
	"fmt"
	"github.com/anon55555/mt"
	"net/http"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

var charMap map[rune]string
var charMapMu sync.Once

const (
	SIGN_WIDTH         = 115
	LINE_LENGTH        = 15
	NUMBER_OF_LINES    = 4
	LINE_HEIGHT        = 14
	CHAR_WIDTH         = 5
	PRINTED_CHAR_WIDTH = CHAR_WIDTH + 1

	SignUrl = "https://git.minetest.land/MineClone2/MineClone2/raw/commit/d887a9731055fb041624cb6a1a09fa4ee7365bf6/mods/ITEMS/mcl_signs/characters.txt"
)

// LoadCharMap downloads the character map from constant SignUrl (if not already have)
func LoadCharMap() {
	charMapMu.Do(func() {
		charMap = make(map[rune]string)

		resp, err := http.Get(SignUrl)
		if err != nil {
			fmt.Println("[SIGNS] couldn't download sign from", SignUrl)
			os.Exit(-1)
		}

		s := bufio.NewScanner(resp.Body)
		var state uint8 = 1
		var char string
		var eChar string
		for s.Scan() {
			switch state {
			case 1:
				char = s.Text()
			case 2:
				eChar = s.Text()
			case 3:
				ru, _ := utf8.DecodeRuneInString(char)
				charMap[ru] = "micl2_" + eChar
				state = 0
			}

			state++
		}
	})
}

// TestSign generates a test sign (wow)
func TestSign(wall bool) mt.Texture {
	return GenerateSignTexture(
		"O"+CenterLine("", '-', 14)+"O\n|"+CenterLine("test", ' ', 14)+"|\n|"+CenterLine("odd", ' ', 14)+"|\nO--------------O",
		wall,
		"black",
	)
}

// center line centers a line, by padding the same left and right (right one more with odd length strings)
func CenterLine(line string, filler rune, width int) string {
	if len(line) > width {
		return line
	}

	// calculate padding
	width = (width - len(line)) / 2

	if len(line)%2 == 0 { // if even
		return strings.Repeat(string(filler), width) + line + strings.Repeat(string(filler), width)
	} else { // not even
		return strings.Repeat(string(filler), width) + line + strings.Repeat(string(filler), width+1)
	}
}

// GenerateSignTexture generates a sign texture
func GenerateSignTexture(text string, wall bool, color string) mt.Texture {
	LoadCharMap()

	texture := fmt.Sprintf("[combine:%dx%d", SIGN_WIDTH, SIGN_WIDTH)
	ypos := 0
	if wall {
		ypos = 30
	}

	for _, line := range strings.Split(text, "\n") {
		xpos := 10
		for _, letter := range line {
			if charMap[letter] != "" {
				texture += fmt.Sprintf(":%d,%d=%s.png", xpos, ypos, charMap[letter])
			}

			xpos += PRINTED_CHAR_WIDTH
		}

		ypos += LINE_HEIGHT
	}

	return mt.Texture(texture + "^[colorize:" + color + ":128")
}
