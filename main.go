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
	go midi.RunMidi(msg)
	go dsp.RunDSP(dsp.DspConf{BufferSize: config.BufferSize}, config.Voices)
	go RunSynth(config.Voices, msg)
	gui.RunGUI(organism.SynthValues{Osc1: &config.OscPanel1, Osc2: &config.OscPanel2})

}
func RunSynth(voices []*generator.Voice, msg chan midi.MidiMsg) {
	pitcChan := make(chan float32)
	for {

		for _, v := range voices {
			for _, t := range v.Tones {

				if t.Type == generator.Regular {
					go t.SendPitch(pitcChan)
					pitcChan <- float32(*v.ControlValues.Pitch)
					t.Osc.SetBaseFreq(float32(*v.ControlValues.Pitch))

					generator.SelectWave(v.ControlValues.Selector.SelectedIndex, t.Osc)
					// *v.Lfo.Rate = *v.ControlValues.LfoR

				}

			}

		}

		select {
		case msg := <-msg:
			config.Voices[0].RunPolly(msg)
			config.Voices[1].RunPolly(msg)

		default:

		}

	}

}
