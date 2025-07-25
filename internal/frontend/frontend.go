package frontend

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 1 2 3 C
// 4 5 6 D
// 7 8 9 E
// A 0 B F
var keymap = [...]ebiten.Key{
	ebiten.KeyX,      // 0
	ebiten.KeyDigit1, // 1
	ebiten.KeyDigit2, // 2
	ebiten.KeyDigit3, // 3
	ebiten.KeyQ,      // 4
	ebiten.KeyW,      // 5
	ebiten.KeyE,      // 6
	ebiten.KeyA,      // 7
	ebiten.KeyS,      // 8
	ebiten.KeyD,      // 9
	ebiten.KeyZ,      // A
	ebiten.KeyC,      // B
	ebiten.KeyDigit4, // C
	ebiten.KeyR,      // D
	ebiten.KeyF,      // E
	ebiten.KeyV,      // F
}

// var keymap = [...]ebiten.Key{
// 	ebiten.KeyDigit1,
// 	ebiten.KeyDigit2,
// 	ebiten.KeyDigit3,
// 	ebiten.KeyDigit4,
// 	ebiten.KeyQ,
// 	ebiten.KeyW,
// 	ebiten.KeyE,
// 	ebiten.KeyR,
// 	ebiten.KeyA,
// 	ebiten.KeyS,
// 	ebiten.KeyD,
// 	ebiten.KeyF,
// 	ebiten.KeyZ,
// 	ebiten.KeyX,
// 	ebiten.KeyC,
// 	ebiten.KeyV,
// }

var (
	FOREGROUND = color.RGBA{0xf8, 0xff, 0xc0, 255} // #f8ffc0
	BACKGROUND = color.RGBA{0x32, 0x2f, 0x2f, 255} // #322f2f
)

type game struct {
	width, height int
	screenBuf     []bool
	keypad        []bool
}

func (g *game) updateKeypad() {
	for i, k := range keymap {
		g.keypad[i] = ebiten.IsKeyPressed(k)
	}
}

func (g *game) Update() error {
	g.updateKeypad()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 5.0, 5.0, 1, 1, color.RGBA{0, 0, 255, 255}, false)

	for y := range g.height {
		for x := range g.width {
			c := BACKGROUND
			if g.screenBuf[y*g.width+x] {
				c = FOREGROUND
			}

			vector.DrawFilledRect(screen, float32(x), float32(y), 1, 1, c, false)
		}
	}

}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}

func Start(screen []bool, width, height, blocksize int, keypad []bool) {

	ebiten.SetWindowSize(width*blocksize, height*blocksize)

	g := &game{width: width, height: height, screenBuf: screen, keypad: keypad}

	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
