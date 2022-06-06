package organism

import (
	components "hodei/gosynthgo/gui/components/molecules"

	g "github.com/AllenDang/giu"
)

type ADSRpanel struct {
	Attack  *int32
	Decay   *int32
	Sustain *int32
	Release *int32
}

var (
	min = int32(0.0)
	max = int32(100.0)
)

func ADSRpanelInit(attackVal *int32, decayval *int32, sustainVal *int32, releaseVal *int32) *g.ColumnWidget {
	attack := g.Column(components.InitSlide(attackVal, min, max),
		components.InitSlideLabel("A"),
	)
	decay := g.Column(components.InitSlide(decayval, min, max),
		components.InitSlideLabel("D"))
	sustain := g.Column(components.InitSlide(sustainVal, min, max),
		components.InitSlideLabel("S"))
	release := g.Column(components.InitSlide(releaseVal, min, max),
		components.InitSlideLabel("R"),
	)
	column := g.Column(
		g.Row(attack, decay, sustain, release),
		components.AdsrGraph(g.GetCursorPos().X+150, g.GetCursorPos().Y-70, *attackVal, *decayval, *sustainVal, *releaseVal),
	)
	return column
}
