package components

import (
	"image/color"

	g "github.com/AllenDang/giu"
)

type SelectorWidget struct {
	pValues    []string
	leftArrow  *g.ArrowButtonWidget
	rightArrow *g.ArrowButtonWidget
	label      *g.LabelWidget
	selected   *int
}

func Selector(pValues []string, s *int) *SelectorWidget {
	return &SelectorWidget{pValues: pValues, leftArrow: g.ArrowButton(g.DirectionLeft), rightArrow: g.ArrowButton(g.DirectionRight), selected: s, label: g.Label(pValues[*s])}

}
func (s *SelectorWidget) Build() *g.RowWidget {

	s.rightArrow.OnClick(func() {
		if *s.selected+1 == len(s.pValues) {
			*s.selected = 0
		} else {
			*s.selected++
		}
	})
	s.leftArrow.OnClick(func() {
		if *s.selected <= 0 {
			*s.selected = len(s.pValues) - 1
		} else {
			*s.selected--
		}
	})
	styledLabel := (g.Style().SetColor(g.StyleColorText, color.RGBA{0, 0, 0, 255}).To(g.Style().SetFontSize(2).To(s.label)))
	return g.Row(s.leftArrow, styledLabel, s.rightArrow)
}
