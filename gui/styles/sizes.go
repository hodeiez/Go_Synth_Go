package styles

import g "github.com/AllenDang/giu"

func VsliderHigh(slider *g.VSliderIntWidget) *g.VSliderIntWidget {
	return slider.Size(30, 300)
}
func VsliderMed(slider *g.VSliderIntWidget) *g.VSliderIntWidget {
	return slider.Size(30, 150)
}
func VsliderSmall(slider *g.VSliderIntWidget) *g.VSliderIntWidget {
	return slider.Size(30, 75)
}
