package generator

import (
	organism "hodei/gosynthgo/gui/components/organisms"
	"hodei/gosynthgo/synth/midi"
	"hodei/gosynthgo/synth/post_audio"
)

type Voice struct {
	Tones []*Tone
	// Osc    []*Tone //polyphony
	Filter *post_audio.Filter
	Adsr   []*Adsr //polyphony
	// Noize  []*Tone //polyphony

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

	}
	noise := NewTone(bufferSize, Noize)
	tones = append(tones, &noise)

	return &Voice{tones, filter, adsr, lfo, controlValues}

}

func (vo *Voice) RunPolly(message midi.MidiMsg) {

	key := findWithKey(vo.Tones, message.Key)
	off := findFirstKeyZeroAndOff(vo.Tones)

	if message.On {
		if key != nil && !key.IsOn {
			key.BindToOSC(message)
			vo.Tones[len(vo.Tones)-1].BindToOSC(message)
		} else if off != nil {
			off.BindToOSC(message)
			vo.Tones[len(vo.Tones)-1].BindToOSC(message)
		}

	} else if !message.On {
		if key != nil && key.IsOn {
			key.BindToOSC(message)
			vo.Tones[len(vo.Tones)-1].BindToOSC(message)

		}

	}

}

func findFirstKeyZeroAndOff(tones []*Tone) *Tone {
	for _, tone := range tones {
		if tone.Key == 0 && !tone.IsOn && tone.Type == Regular {
			return tone
		}
	}
	return nil
}
func findWithKey(tones []*Tone, key int) *Tone {
	for _, tone := range tones {
		if tone.Key == key && tone.Type == Regular {
			return tone
		}
	}
	return nil
}
