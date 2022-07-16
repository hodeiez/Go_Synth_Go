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
type VoiceManager struct {
	Voices []*Voice
}

// func NewVoice(oscs []*Osc, filter *post_audio.Filter, adsr []*Adsr, noize []*Osc, lfo *Lfo, controlValues organism.OscPanelValues) *Voice {
// 	return &Voice{oscs, filter, adsr, noize, lfo, controlValues}
// }
func NewVoice(filter *post_audio.Filter, adsr []*Adsr, lfo *Lfo, controlValues organism.OscPanelValues, polyphony int, bufferSize int) *Voice {
	oscs := make([]*Osc, polyphony)
	noizes := make([]*Osc, polyphony)

	for i := range oscs {
		osc := Oscillator(bufferSize)
		noise := NoiseOsc(bufferSize)
		oscs[i] = &osc
		noizes[i] = &noise
	}
	return &Voice{oscs, filter, adsr, noizes, lfo, controlValues}
}
