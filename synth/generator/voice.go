package generator

import (
	organism "hodei/gosynthgo/gui/components/organisms"
	"hodei/gosynthgo/synth/midi"
	"hodei/gosynthgo/synth/post_audio"
)

type Voice struct {
	Tones         []*Tone
	Filter        *post_audio.Filter
	Adsr          *Adsr //polyphony
	Lfo           *Lfo
	ControlValues organism.OscPanelValues
}
type VoiceManager struct {
	Voices []*Voice
}

func NewVoice(filter *post_audio.Filter, adsr *Adsr, lfo *Lfo, controlValues organism.OscPanelValues, polyphony int, bufferSize int) *Voice {

	var tones []*Tone
	for i := 0; i <= polyphony; i++ {
		osc := NewTone(bufferSize, Regular)
		tones = append(tones, &osc)

		noise := NewTone(bufferSize, Noize)
		tones = append(tones, &noise)
	}

	return &Voice{tones, filter, adsr, lfo, controlValues}

}

func (vo *Voice) RunPolly(message midi.MidiMsg, pitcChan chan float64) {

	for _, t := range vo.Tones {

		if t.Type == Regular {
			go t.SendPitch(pitcChan)
			pitcChan <- *vo.ControlValues.Pitch
			t.Osc.SetBaseFreq(*vo.ControlValues.Pitch)

			SelectWave(vo.ControlValues.Selector.SelectedIndex, t.Osc)
			*vo.Lfo.Rate = *vo.ControlValues.LfoR
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

	oscKey := findWithKey(vo.Tones, message.Key, Regular)
	noizeKey := findWithKey(vo.Tones, message.Key, Noize)
	oscOff := findFirstOff(vo.Tones, Regular)
	noizeOff := findFirstOff(vo.Tones, Noize)
	// for _, tone := range vo.Tones {

	// 	println(tone.ShowStatus())

	// }
	if message.On {

		if oscOff != nil && noizeOff != nil {
			oscOff.BindToOSC(message, vo.Adsr)
			noizeOff.BindToOSC(message, vo.Adsr)

		}

	} else if !message.On {
		if oscKey != nil && oscKey.IsOn && noizeKey != nil && noizeKey.IsOn {

			// message.Key = 0
			oscKey.BindToOSC(message, vo.Adsr)
			noizeKey.BindToOSC(message, vo.Adsr)

		}

	}

}

// func findFirstKeyZeroAndOff(tones []*Tone, oscType OscType) *Tone {
// 	for _, tone := range tones {
// 		if tone.Key == 0 && !tone.IsOn && tone.Type == oscType {
// 			return tone
// 		}
// 	}
// 	return nil
// }
func findWithKey(tones []*Tone, key int, oscType OscType) *Tone {
	for _, tone := range tones {
		if tone.Key == key && tone.Type == oscType && tone.Active && tone.IsOn {
			return tone
		}
	}
	return nil
}
func findFirstOff(tones []*Tone, oscType OscType) *Tone {
	for _, tone := range tones {
		if !tone.IsOn && tone.Type == oscType && !tone.Active {
			return tone
		}
	}
	return nil
}
