package organism

import (
	components "hodei/gosynthgo/gui/components/molecules"

	g "github.com/AllenDang/giu"
)

type Configpanel struct {
	SampleRate   *int32
	DelaySeconds *int32
}

var (
	mins = int32(0.0) //TODO:check in generator.Adsr this value to be 0
	maxs = int32(100.0)
)

func configPanelInit(sampleRate *int32, delay *int32) *g.ColumnWidget {
	sr := g.Column(components.InitSlideH(sampleRate, 22050, 48000),
		components.InitSlideLabel("Sample Rate"),
	)
	dl := g.Column(components.InitSlideH(delay, 0, 10000),
		components.InitSlideLabel("LatencyMs"),
	)
	column := g.Column(
		g.Row(sr, dl),
		// components.AdsrGraph(g.GetCursorPos().X+150, g.GetCursorPos().Y-70, *attackVal, *decayval, *sustainVal, *releaseVal),
	)
	return column
}
