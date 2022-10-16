package main

import (
	config "hodei/gosynthgo/config"
	gui "hodei/gosynthgo/gui"
	organism "hodei/gosynthgo/gui/components/organisms"
	"hodei/gosynthgo/synth/dsp"
	"hodei/gosynthgo/synth/generator"
	"hodei/gosynthgo/synth/midi"
)

func main() {
	///main vars
	msg := make(chan midi.MidiMsg)
	out := 0.0
	dspConfig := dsp.DspConf{BufferSize: config.BufferSize}
	p := dsp.NewProcessAudio(&out, config.Voices, dspConfig)
	go midi.RunMidi(msg)

	go dsp.RunDSP(p, dspConfig)
	go gui.RunGUI(organism.SynthValues{Osc1: &config.OscPanel1, Osc2: &config.OscPanel2})
	pitcChan := make(chan float64)
	RunSynth(config.Voices, msg, pitcChan)

}

func RunSynth(voices []*generator.Voice, msg chan midi.MidiMsg, pitcChan chan float64) {
	for {

		select {
		case msg := <-msg:
			config.Voices[0].RunPolly(msg, pitcChan)
			config.Voices[1].RunPolly(msg, pitcChan)

		default:

		}

	}

}
