package signs

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/anon55555/mt"
)

var charMap map[rune]string
var charMapMu sync.Once

// Prefix is the prefix before the chars
var Prefix string = "micl2_"

// CharUrl is the Url the characters.txt file will be loaded from (default is official mineclone2 repos)
var CharUrl = "https://git.minetest.land/MineClone2/MineClone2/raw/commit/d887a9731055fb041624cb6a1a09fa4ee7365bf6/mods/ITEMS/mcl_signs/characters.txt"

const (
	SIGN_WIDTH         = 115
	LINE_LENGTH        = 15
	NUMBER_OF_LINES    = 4
	LINE_HEIGHT        = 13
	CHAR_WIDTH         = 5
	PRINTED_CHAR_WIDTH = CHAR_WIDTH + 1
)

// LoadCharMap downloads the character map from constant SignUrl (if not already have)
func LoadCharMap() {
	charMapMu.Do(func() {
		charMap = make(map[rune]string)

		resp, err := http.Get(CharUrl)
		if err != nil {
			fmt.Println("[SIGNS] couldn't download sign from", CharUrl)
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
				charMap[ru] = Prefix + eChar
				state = 0
			}

			state++
		}
	})
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

// GenerateSignTexture generates a sign texture (color has to be mt colorspec)
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

func GenerateTextureAOMod(text string, wall bool, color string) *mt.AOCmdTextureMod {
	return &mt.AOCmdTextureMod{
		Mod: "^" + GenerateSignTexture(text, wall, color),
	}
}
