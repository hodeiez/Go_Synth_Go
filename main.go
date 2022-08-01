package main

import (
	"fmt"
	gui "hodei/gosynthgo/gui"
	organism "hodei/gosynthgo/gui/components/organisms"
	"hodei/gosynthgo/synth/dsp"
	"hodei/gosynthgo/synth/generator"
	"hodei/gosynthgo/synth/midi"
	"hodei/gosynthgo/synth/post_audio"
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
	string_first = append(string_first, "saw", "tri", "squ", "sin")
	return string_first
}

func main() {
	///main vars
	msg := make(chan midi.MidiMsg)
	const (
		bufferSize = 2048
		polyphony  = 2
	)

	voice1 := generator.NewVoice(&post_audio.Filter{oscPanel1.Cut, oscPanel1.Res}, []*generator.Adsr{&generator.Adsr{}}, &generator.Lfo{}, oscPanel1, polyphony, bufferSize)
	voice2 := generator.NewVoice(&post_audio.Filter{oscPanel2.Cut, oscPanel2.Res}, []*generator.Adsr{&generator.Adsr{}}, &generator.Lfo{}, oscPanel2, polyphony, bufferSize)
	voices := []*generator.Voice{voice1, voice2}

	//run processes
	go midi.RunMidi(msg)
	go gui.RunGUI(organism.SynthValues{Osc1: &voice1.ControlValues, Osc2: &oscPanel2})
	go dsp.RunDSP(dsp.DspConf{BufferSize: bufferSize}, voices)
	//main loop
	for {
		//TODO:refactor to binding/controller function
		for _, v := range voices {

			for _, o := range v.Osc {
				o.Osc.Amplitude = *v.ControlValues.Vol / 1000
				o.SetBaseFreq(*v.ControlValues.Pitch * 10) //TODO: make pitchShift function
				// o.Osc.SetFreq(*v.ControlValues.Pitch * 10)
				generator.SelectWave(v.ControlValues.Selector.SelectedIndex, *o)

			}
			for _, n := range v.Noize {

				n.SetBaseFreq(*v.ControlValues.Pitch * 10)
				n.Osc.Amplitude = *v.ControlValues.Noize / 1000
			}
		}

		select {
		case msg := <-msg:
			for _, v := range voices {
				for _, o := range v.Osc {
					*o = generator.ChangeFreq(msg, o)
				}
				for _, n := range v.Noize {
					generator.ChangeFreq(msg, n)
				}
			}
			fmt.Println(msg)
		default:

		}

		// fmt.Println(<-msg)

	}
}
