package generator

import "hodei/gosynthgo/synth/post_audio"

type Voice struct {
	Osc    []*Osc //polyphony
	Filter *post_audio.Filter
	Adsr   []*Adsr  //polyphony
	Noize  []*Noise //polyphony
	Lfo    *Lfo
}
