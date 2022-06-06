package components

import (
	styles "hodei/gosynthgo/gui/styles"

	g "github.com/AllenDang/giu"
)

func SlideLabel(labelText string) *g.LabelWidget {
	label := g.Label(labelText)

	return label

}
func slideLabelStyled(label *g.LabelWidget) *g.StyleSetter {

	return g.Style().SetColor(g.StyleColorText, styles.Redhite).To(label)
}

func InitSlideLabel(label string) *g.StyleSetter {
	return slideLabelStyled(SlideLabel(label))
}
