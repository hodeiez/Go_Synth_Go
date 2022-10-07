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
	go RunSynth(config.Voices, msg)

	go dsp.RunDSP(p, dspConfig)
	gui.RunGUI(organism.SynthValues{Osc1: &config.OscPanel1, Osc2: &config.OscPanel2})

}

func RunSynth(voices []*generator.Voice, msg chan midi.MidiMsg) {
	pitcChan := make(chan float64)
	// pitcChan2 := make(chan float64)
	for {

		for _, v := range voices {
			for _, t := range v.Tones {

				if t.Type == generator.Regular {
					go t.SendPitch(pitcChan)
					pitcChan <- *v.ControlValues.Pitch
					t.Osc.SetBaseFreq(*v.ControlValues.Pitch)

					generator.SelectWave(v.ControlValues.Selector.SelectedIndex, t.Osc)
					*v.Lfo.Rate = *v.ControlValues.LfoR
					// println(v.Lfo.Main.Osc.Sample())

					// if *v.Lfo.Rate > 0 {
					// 	// go t.SendPitch(pitcChan2)
					// 	*v.ControlValues.Pitch = generator.RescaleThis(v.Lfo.Main.Osc.Sample())
					// 	// t.Osc.SetBaseFreq(generator.RescaleThis(v.Lfo.Main.Osc.Sample()))
					// 	// pitcChan2 <- generator.RescaleThis(v.Lfo.Main.Osc.Sample())
					// 	// t.Osc.SetBaseFreq(generator.RescaleThis(v.Lfo.Main.Osc.Sample()))
					// }

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
