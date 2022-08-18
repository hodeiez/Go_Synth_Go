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
	go gui.RunGUI(organism.SynthValues{Osc1: &config.OscPanel1, Osc2: &config.OscPanel2})
	go dsp.RunDSP(dsp.DspConf{BufferSize: config.BufferSize}, config.Voices)

	pitcChan := make(chan float64)
	for {

		for _, v := range config.Voices {
			for _, t := range v.Tones {

				if t.Type == generator.Regular {

					go t.SendPitch(pitcChan)
					pitcChan <- *v.ControlValues.Pitch
					t.Osc.SetBaseFreq(*v.ControlValues.Pitch)

					generator.SelectWave(v.ControlValues.Selector.SelectedIndex, t.Osc)
				}
				//TODO:move to dsp
				// if t.Type == generator.Noize {
				// 	t.Osc.Osc.Amplitude = 0.0
				// 	// if t.Osc.Osc.Amplitude != 0.0 {
				// 	// 	t.Osc.Osc.Amplitude -= *v.ControlValues.Noize / 1000
				// 	// }

				// }

			}
			// for _, n := range v.Noize {

			// 	//n.SetBaseFreq(*v.ControlValues.Pitch * 10)
			// 	n.Osc.Osc.Amplitude = *v.ControlValues.Noize / 1000
			// }
		}

		select {
		case msg := <-msg:
			config.Voices[0].RunPolly(msg)
			config.Voices[1].RunPolly(msg)

		default:

		}

		// fmt.Println(<-msg)

	}
}
