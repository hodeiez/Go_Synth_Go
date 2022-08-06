package generator

import (
	organism "hodei/gosynthgo/gui/components/organisms"
	"hodei/gosynthgo/synth/midi"
	"hodei/gosynthgo/synth/post_audio"
)

type Voice struct {
	Tones         []*Tone
	Filter        *post_audio.Filter
	Adsr          []*Adsr //polyphony
	Lfo           *Lfo
	ControlValues organism.OscPanelValues
}
type VoiceManager struct {
	Voices []*Voice
}

func NewVoice(filter *post_audio.Filter, adsr []*Adsr, lfo *Lfo, controlValues organism.OscPanelValues, polyphony int, bufferSize int) *Voice {

	var tones []*Tone
	for i := 0; i <= polyphony; i++ {
		osc := NewTone(bufferSize, Regular)
		tones = append(tones, &osc)

		noise := NewTone(bufferSize, Noize)
		tones = append(tones, &noise)
	}

	return &Voice{tones, filter, adsr, lfo, controlValues}

}

//TODO: decide if noize is property in TOne as osc
func (vo *Voice) RunPolly(message midi.MidiMsg) {

	oscKey := findWithKey(vo.Tones, message.Key, Regular)
	noizeKey := findWithKey(vo.Tones, message.Key, Noize)
	oscOff := findFirstKeyZeroAndOff(vo.Tones, Regular)
	noizeOff := findFirstKeyZeroAndOff(vo.Tones, Noize)

	if message.On {
		if oscKey != nil && !oscKey.IsOn && noizeKey != nil && !noizeKey.IsOn {
			oscKey.BindToOSC(message)
			noizeKey.BindToOSC(message)

		} else if oscOff != nil && noizeOff != nil {
			oscOff.BindToOSC(message)
			noizeOff.BindToOSC(message)

		}

	} else if !message.On {
		if oscKey != nil && oscKey.IsOn && noizeKey != nil && noizeKey.IsOn {
			oscKey.BindToOSC(message)
			noizeKey.BindToOSC(message)

		}

	}

}

func findFirstKeyZeroAndOff(tones []*Tone, oscType OscType) *Tone {
	for _, tone := range tones {
		if tone.Key == 0 && !tone.IsOn && tone.Type == oscType {
			return tone
		}
	}
	return nil
}
func findWithKey(tones []*Tone, key int, oscType OscType) *Tone {
	for _, tone := range tones {
		if tone.Key == key && tone.Type == oscType {
			return tone
		}
	}
	return nil
}
