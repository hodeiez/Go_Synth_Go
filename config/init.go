package config

import (
	organism "hodei/gosynthgo/gui/components/organisms"
	"hodei/gosynthgo/synth/generator"
	"hodei/gosynthgo/synth/post_audio"
)

var (
	pitchInit = float64(440.0)
)

var (
	vol1   = 100.0
	pitch1 = 440.0
	cut1   = 546.0
	res1   = 0.0
	pwm1   = 0.0
	lfoR1  = 0.0
	lfoW1  = 0.0
	noize1 = 0.0
)
var (
	vol2   = 0.0
	pitch2 = 440.0
	cut2   = 0.0
	res2   = 0.0
	pwm2   = 0.0
	lfoR2  = 0.0
	lfoW2  = 0.0
	noize2 = 0.0
)
var adsr = organism.AdsrValues{Att: 0, Dec: 0, Sus: 100, Rel: 0}
var adsr2 = organism.AdsrValues{Att: 0, Dec: 0, Sus: 0, Rel: 0}
var selector1 = organism.SelectorValues{TextValues: testSelector(), SelectedIndex: 0}
var selector2 = organism.SelectorValues{TextValues: testSelector(), SelectedIndex: 0}
var OscPanel1 = organism.OscPanelValues{Adsr: &adsr,
	Vol:       &vol1,
	Pitch:     &pitch1,
	PitchInit: &pitchInit,
	Cut:       &cut1,
	Res:       &res1,
	Pwm:       &pwm1,
	LfoR:      &lfoR1,
	LfoW:      &lfoW1,
	Noize:     &noize1,
	Selector:  &selector1,
}
var OscPanel2 = organism.OscPanelValues{Adsr: &adsr2,
	Vol:       &vol2,
	Pitch:     &pitch2,
	PitchInit: &pitchInit,
	Cut:       &cut2,
	Res:       &res2,
	Pwm:       &pwm2,
	LfoR:      &lfoR2,
	LfoW:      &lfoW2,
	Noize:     &noize2,
	Selector:  &selector2,
}
var ValToPass = organism.SynthValues{Osc1: &OscPanel1, Osc2: &OscPanel2}

func testSelector() []string {
	var string_first []string
	string_first = append(string_first, "saw", "tri", "squ", "sin")
	return string_first
}

const (
	SampleRate   = 48000
	DelaySeconds = 0.1
)
const (
	BufferSize = SampleRate * DelaySeconds //
	polyphony  = 40
)

var voice1 = generator.NewVoice(&post_audio.Filter{Cutoff: OscPanel1.Cut, Reso: OscPanel1.Res},
	&generator.Adsr{
		AttackTime:  &OscPanel1.Adsr.Att,
		DecayTime:   &OscPanel1.Adsr.Dec,
		ReleaseTime: &OscPanel1.Adsr.Rel,
		SustainAmp:  &OscPanel1.Adsr.Sus,
		Type:        generator.EnvelopeAdsr,
		MinValue:    0.0,
		MaxValue:    0.0,
		StopTime:    make(chan bool, 1)},
	generator.NewLFO(BufferSize, OscPanel1.LfoR),
	// &generator.Lfo{generator.NewLFO(BufferSize)},
	OscPanel1,
	polyphony,
	BufferSize)

var voice2 = generator.NewVoice(&post_audio.Filter{Cutoff: OscPanel2.Cut, Reso: OscPanel2.Res},
	&generator.Adsr{
		AttackTime:  &OscPanel2.Adsr.Att,
		DecayTime:   &OscPanel2.Adsr.Dec,
		ReleaseTime: &OscPanel2.Adsr.Rel,
		SustainAmp:  &OscPanel2.Adsr.Sus,
		Type:        generator.EnvelopeAdsr,
		MinValue:    0.0,
		MaxValue:    0.0,
		StopTime:    make(chan bool, 1)},
	generator.NewLFO(BufferSize, OscPanel2.LfoR),
	// &generator.Lfo{},
	OscPanel2,
	polyphony,
	BufferSize)

var Voices = []*generator.Voice{voice1, voice2}
