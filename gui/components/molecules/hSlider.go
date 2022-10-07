package components

import (
	g "github.com/AllenDang/giu"
	styles "hodei/gosynthgo/gui/styles"
)

func SlideH(start *int32, min *int32, max *int32) *g.SliderIntWidget {
	slider := g.SliderInt(start, *min, *max)

	return slider

}
func hslideStyled(slider g.SliderIntWidget) *g.StyleSetter {
	styles.HsliderMed(&slider)

	return g.Style().SetColor(g.StyleColorSliderGrabActive, styles.Red).
		To(
			g.Style().SetColor(g.StyleColorSliderGrab, styles.Redhite).
				To(
					g.Style().SetColor(g.StyleColorBorderShadow, styles.Redhite).
						To(&slider)))

}

func InitSlideH(start *int32, min int32, max int32) *g.StyleSetter {
	return hslideStyled(*SlideH(start, &min, &max))
}
