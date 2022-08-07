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

	voice1 := generator.NewVoice(&post_audio.Filter{config.OscPanel1.Cut, config.OscPanel1.Res}, []*generator.Adsr{&generator.Adsr{}}, &generator.Lfo{}, config.OscPanel1, polyphony, bufferSize)
	voice2 := generator.NewVoice(&post_audio.Filter{config.OscPanel2.Cut, config.OscPanel2.Res}, []*generator.Adsr{&generator.Adsr{}}, &generator.Lfo{}, config.OscPanel2, polyphony, bufferSize)
	voices := []*generator.Voice{voice1, voice2}

	//run processes
	go midi.RunMidi(msg)
	go gui.RunGUI(organism.SynthValues{Osc1: &voice1.ControlValues, Osc2: &config.OscPanel2})
	go dsp.RunDSP(dsp.DspConf{BufferSize: bufferSize}, voices)
	//main loop
	for {
		//TODO:refactor to binding/controller function
		for _, v := range voices {
			for _, o := range v.Tones {
				//o.SetBaseFreq(*v.ControlValues.Pitch * 10) //TODO: make pitchShift function
				// o.Osc.SetFreq(*v.ControlValues.Pitch * 10)

				if o.Type == generator.Regular {

					//temp := o.Osc.Osc.Amplitude + (*v.ControlValues.Vol / 6000)

					// if o.Osc.Osc.Amplitude != temp && o.Osc.Osc.Amplitude != 0.0 {
					// 	o.Osc.Osc.Amplitude = temp
					// }

					generator.SelectWave(v.ControlValues.Selector.SelectedIndex, o.Osc)
				}
				if o.Type == generator.Noize {
					if o.Osc.Osc.Amplitude != 0.0 {
						o.Osc.Osc.Amplitude = *v.ControlValues.Noize / 1000
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
