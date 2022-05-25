package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"sync"
	"strconv"
)

type charMapE struct {
	encoded string
	width   int64
}

var charMap   map[rune]*charMapE
var charMapMu sync.Once

const (
	SIGN_WIDTH = 115
	LINE_LENGTH = 15
	NUMBER_OF_LINES = 4
	LINE_HEIGHT = 14
	CHAR_WIDTH = 5
	PRINTED_CHAR_WIDTH = CHAR_WIDTH + 1
)

func loadCharMap() {
	charMapMu.Do(func() {
		f, err := os.Open("characters.txt")
		if err != nil {
			fmt.Println("[MT_SIGNS] Cant read characters.txt file!")
			os.Exit(-1)
		}

		s := bufio.NewScanner(f)
		var state uint8 = 0
		var char string
		var eChar string
		
		for s.Scan() {
			switch state {
			case 0: // every third line
				char = s.Text()
			case 1:
				eChar = s.Text()
			case 2:
				w, _ := strconv.ParseInt(s.Text(), 10, 64)
			
				charMap[rune(char[0])] = &charMapE{
					encoded: eChar,
					width: w,
				}

				state = 0
			}
		
			state++
		}
	})
}

func generateTexture(text string) string {
	texture := fmt.Sprintf("[combine:%nx%n", SIGN_WIDTH, SIGN_WIDTH)

	// TODO: "mcl_signs:wall_sign" starts at ypos = 30 else 0

	ypos := 0

	for _, line := range strings.Split(text, "\n") {
		xpos := 0
	
		for _, letter := range line {
			if charMap[letter] != nil {
				// ":"..xpos..","..ypos.."="..parsed[i]..".png"
				texture += fmt.Sprintf(":%n,%n=%s.png", xpos, ypos, charMap[letter].encoded)

				xpos += PRINTED_CHAR_WIDTH
			}
		}

		ypos += LINE_HEIGHT
	}

	return texture
}
