package generator

import "hodei/gosynthgo/synth/midi"

type Tone struct {
	Key  int
	IsOn bool
	Osc  Osc
	Type OscType
}

func NewTone(buffer int, oscType OscType) Tone {
	switch oscType {
	case Regular:
		return Tone{Key: 0, IsOn: false, Osc: Oscillator(buffer), Type: Regular}
	case Noize:
		return Tone{Key: 0, IsOn: false, Osc: NoiseOsc(buffer), Type: Noize}
	default:
		return Tone{Key: 0, IsOn: false, Osc: Oscillator(buffer), Type: Regular}

	}
}
func (t *Tone) BindToOSC(message midi.MidiMsg) {
	t.Key = message.Key
	t.IsOn = message.On
	t.Osc.ChangeFreq(message)
	//TODO: change this when ADSR is on
	if t.IsOn {
		if t.Type == Regular {
			t.Osc.Osc.Amplitude = 0.01
		} else {
			t.Osc.Osc.Amplitude = 0.01
		}
	} else {
		t.Osc.Osc.Amplitude = 0.0000
	}
}
