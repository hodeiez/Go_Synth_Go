package components

import (
	"image/color"

	g "github.com/AllenDang/giu"
)

type MyButton struct {
	ColorBG     color.RGBA
	ColorAction color.RGBA
	ColorHover  color.RGBA
	InVal       *float64
	SetToVal    float64
	Text        string
}

func Button(colorBG color.RGBA, coloraction color.RGBA, colorHover color.RGBA, inval *float64, setToVal float64, text string) *MyButton {
	return &MyButton{colorBG, coloraction, colorHover, inval, setToVal, text}
}
func (b *MyButton) Build() *g.StyleSetter {
	buttonWithAction := g.Button(b.Text).OnClick(func() {
		*b.InVal = b.SetToVal
	})

	return g.Style().SetColor(g.StyleColorButtonActive, b.ColorAction).To(
		g.Style().SetColor(g.StyleColorButtonHovered, b.ColorHover).To(
			g.Style().SetColor(g.StyleColorButton, b.ColorBG).To(buttonWithAction)))

}

// func myButton(valPitch *float64, initVal float64) *g.ButtonWidget {
// 	return g.Button("Tune me").OnClick(func() {
// 		*valPitch = initVal
// 		println(valPitch)
// 	})
// }

// g.Style().SetColor(g.StyleColorButton, color.RGBA{0, 0, 0, 255}).To(myButton(val.Pitch, *val.PitchInit))
