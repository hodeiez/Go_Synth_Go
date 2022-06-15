package generator

import (
	organism "hodei/gosynthgo/gui/components/organisms"
	"hodei/gosynthgo/synth/post_audio"
)

type Voice struct {
	Osc           []*Osc //polyphony
	Filter        *post_audio.Filter
	Adsr          []*Adsr //polyphony
	Noize         []*Osc  //polyphony
	Lfo           *Lfo
	ControlValues organism.OscPanelValues
}

func NewVoice(oscs []*Osc, filter *post_audio.Filter, adsr []*Adsr, noize []*Osc, lfo *Lfo, controlValues organism.OscPanelValues) *Voice {
	return &Voice{oscs, filter, adsr, noize, lfo, controlValues}
}
