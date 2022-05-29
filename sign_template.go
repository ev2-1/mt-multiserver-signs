package signs

import (
	"github.com/anon55555/mt"

	"image/color"
	"math"
	"sync"
)

var offsets [16][3]float32
var offsetsW [4][3]float32
var mathsMu sync.Once

func Maths() {
	mathsMu.Do(func() {
		m := (-1.0/16.0 + 1.0/64.0) * 10.0 * 1.5
		var angles [16]float32

		for i := 0; i < 16; i++ {
			angles[i] = math.Pi*2.0 - ((math.Pi*2.0)/16.0)*float32(i)
			offsets[i] = [3]float32{float32(math.Sin(float64(angles[i])) * m), float32(2.0 / 28.0), float32(math.Cos(float64(angles[i])) * m)}
		}

		offsetsW[0] = [3]float32{0, 0, 4}
		offsetsW[1] = [3]float32{-4, 0, 0}
		offsetsW[2] = [3]float32{0, 0, -4}
		offsetsW[3] = [3]float32{4, 0, 0}
	})
}

func SignProps() mt.AOProps {
	return mt.AOProps{
		MaxHP:            10,
		Pointable:        false,
		Visual:           "upright_sprite",
		VisualSize:       [3]float32{1.0, 1.0, 1.0},
		Textures:         []mt.Texture{"[combine:115x115"},
		DmgTextureMod:    "^[brighten",
		Shaded:           true,
		SpriteSheetSize:  [2]int16{1, 1},
		SpritePos:        [2]int16{0, 0},
		Visible:          true,
		Colors:           []color.NRGBA{color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}},
		BackfaceCull:     true,
		NametagColor:     color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		NametagBG:        color.NRGBA{R: 0x01, G: 0x01, B: 0x01, A: 0x00},
		FaceRotateSpeed:  -1,
		Infotext:         "mcl_signs:standing_sign",
		Itemstring:       "",
	}
}

type Rotate uint8

const (
	North Rotate = iota
	North22_5
	North45
	North67_5
	East
	East22_5
	East45
	East67_5
	South
	South22_5
	South45
	South67_5
	West
	West22_5
	West45
	West67_5
)

func rot2Vec(r Rotate) mt.Vec {
	return mt.Vec{0, float32(r) / 4 * 90, 0}
}

func toPos(p [3]int16, r Rotate, w bool) (f [3]float32) {
	Maths()
	if w {
		for k := range f {
			f[k] = float32(p[k]*10) + offsetsW[r/4][k]
		}
	} else {
		for k := range f {
			f[k] = float32(p[k]*10) + offsets[r][k]
		}
	}
	return
}

func GenerateSignAOAdd(text, color string, pos [3]int16, rotation Rotate, wall bool, id mt.AOID) mt.AOAdd {
	return mt.AOAdd{
		ID: id,
		InitData: mt.AOInitData{
			ID:  id,
			Pos: toPos(pos, rotation, wall),
			Rot: rot2Vec(rotation),
			HP:  10,
			Msgs: []mt.AOMsg{
				&mt.AOCmdProps{
					Props: SignProps(),
				},
				&mt.AOCmdArmorGroups{
					Armor: []mt.Group{
						mt.Group{
							Name:   "immortal",
							Rating: 1,
						},
					},
				},
				&mt.AOCmdAttach{
					Attach: mt.AOAttach{
						ParentID:     0,
						Bone:         "",
						Pos:          mt.Vec{0, 0, 0},
						Rot:          mt.Vec{0, 0, 0},
						ForceVisible: false,
					},
				},
				GenerateTextureAOMod(text, wall, color),
			},
		},
	}
}
