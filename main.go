package main

import (
	gui "hodei/gosynthgo/gui"
	organism "hodei/gosynthgo/gui/components/organisms"
	"hodei/gosynthgo/synth/dsp"
	"hodei/gosynthgo/synth/generator"
	//"log"
)

var (
	vol1   = 0.0
	pitch1 = 0.0
	cut1   = 0.0
	res1   = 0.0
	pwm1   = 0.0
	lfoR1  = 0.0
	lfoW1  = 0.0
	noize1 = 0.0
)
var (
	vol2   = 0.0
	pitch2 = 0.0
	cut2   = 0.0
	res2   = 0.0
	pwm2   = 0.0
	lfoR2  = 0.0
	lfoW2  = 0.0
	noize2 = 0.0
)
var adsr = organism.AdsrValues{Att: 0, Dec: 0, Sus: 0, Rel: 0}
var adsr2 = organism.AdsrValues{Att: 0, Dec: 0, Sus: 0, Rel: 0}
var selector1 = organism.SelectorValues{TextValues: testSelector(), SelectedIndex: 0}
var selector2 = organism.SelectorValues{TextValues: testSelector(), SelectedIndex: 0}
var oscPanel1 = organism.OscPanelValues{Adsr: &adsr,
	Vol:      &vol1,
	Pitch:    &pitch1,
	Cut:      &cut1,
	Res:      &res1,
	Pwm:      &pwm1,
	LfoR:     &lfoR1,
	LfoW:     &lfoW1,
	Noize:    &noize1,
	Selector: &selector1,
}
var oscPanel2 = organism.OscPanelValues{Adsr: &adsr2,
	Vol:      &vol2,
	Pitch:    &pitch2,
	Cut:      &cut2,
	Res:      &res2,
	Pwm:      &pwm2,
	LfoR:     &lfoR2,
	LfoW:     &lfoW2,
	Noize:    &noize2,
	Selector: &selector2,
}
var valToPass = organism.SynthValues{Osc1: &oscPanel1, Osc2: &oscPanel2}

func testSelector() []string {
	var string_first []string
	string_first = append(string_first, "saw")
	string_first = append(string_first, "tri")
	string_first = append(string_first, "squ")
	string_first = append(string_first, "sin")
	return string_first
}

func main() {
	osc := generator.Oscillator(2048)
	go gui.RunGUI(valToPass)
	go dsp.RunDSP(dsp.DspConf{BufferSize: 2048}, osc)
	for {
		osc.Osc.Amplitude = *valToPass.Osc1.Vol / 1000
		osc.Osc.SetFreq(*valToPass.Osc1.Pitch * 10)
		//log.Println(*valToPass.Osc1.Vol)
	}
}
