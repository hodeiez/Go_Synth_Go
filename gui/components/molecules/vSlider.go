package components

import (
	g "github.com/AllenDang/giu"
	styles "hodei/gosynthgo/gui/styles"
)

func Slide(start *int32, min *int32, max *int32) *g.VSliderIntWidget {
	slider := g.VSliderInt(start, *min, *max)

	return slider

}
func slideStyled(slider g.VSliderIntWidget) *g.StyleSetter {
	styles.VsliderMed(&slider)

	return g.Style().SetColor(g.StyleColorSliderGrabActive, styles.Red).
		To(
			g.Style().SetColor(g.StyleColorSliderGrab, styles.Redhite).
				To(
					g.Style().SetColor(g.StyleColorBorderShadow, styles.Redhite).
						To(&slider)))

}

func InitSlide(start *int32, min int32, max int32) *g.StyleSetter {
	return slideStyled(*Slide(start, &min, &max))
}
