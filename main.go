package main

import (
	config "hodei/gosynthgo/config"
	gui "hodei/gosynthgo/gui"
	organism "hodei/gosynthgo/gui/components/organisms"
	"hodei/gosynthgo/synth/dsp"
	"hodei/gosynthgo/synth/generator"
	"hodei/gosynthgo/synth/midi"
	"hodei/gosynthgo/synth/post_audio"
)

func main() {
	///main vars
	msg := make(chan midi.MidiMsg)
	const (
		bufferSize = 2048
		polyphony  = 40
	)

	voice1 := generator.NewVoice(&post_audio.Filter{Cutoff: config.OscPanel1.Cut, Reso: config.OscPanel1.Res}, []*generator.Adsr{&generator.Adsr{}}, &generator.Lfo{}, config.OscPanel1, polyphony, bufferSize)
	voice2 := generator.NewVoice(&post_audio.Filter{Cutoff: config.OscPanel2.Cut, Reso: config.OscPanel2.Res}, []*generator.Adsr{&generator.Adsr{}}, &generator.Lfo{}, config.OscPanel2, polyphony, bufferSize)
	voices := []*generator.Voice{voice1, voice2}
	//run processes
	go midi.RunMidi(msg)
	go gui.RunGUI(organism.SynthValues{Osc1: &voice1.ControlValues, Osc2: &config.OscPanel2})
	go dsp.RunDSP(dsp.DspConf{BufferSize: bufferSize}, voices)

	//main loop
	// pitch := 0.0
	pitchTest := make(chan float64)
	for {

		//TODO:refactor to binding/controller function
		for _, v := range voices {
			for _, t := range v.Tones {

				if t.Type == generator.Regular {

					go t.SendPitch(pitchTest)
					pitchTest <- *v.ControlValues.Pitch
					t.Osc.SetBaseFreq(*v.ControlValues.Pitch)

					generator.SelectWave(v.ControlValues.Selector.SelectedIndex, t.Osc)
				}
				if t.Type == generator.Noize {
					if t.Osc.Osc.Amplitude != 0.0 {
						t.Osc.Osc.Amplitude = *v.ControlValues.Noize / 1000
					}

				}

			}
			// for _, n := range v.Noize {

			// 	//n.SetBaseFreq(*v.ControlValues.Pitch * 10)
			// 	n.Osc.Osc.Amplitude = *v.ControlValues.Noize / 1000
			// }
		}

		select {
		case msg := <-msg:
			voice1.RunPolly(msg)
			voice2.RunPolly(msg)

		default:

		}

		// fmt.Println(<-msg)

	}
}
