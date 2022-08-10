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
			t.Osc.Osc.Amplitude = RescaleMidiValues(message.Vel, 0.0, 0.1) //it gets velocity value
			// t.Osc.Osc.Amplitude = 0.01 * float64(message.Vel)
		} else {
			t.Osc.Osc.Amplitude = 0.01
		}
	} else {
		t.Osc.Osc.Amplitude = 0.0000
	}
}

func (t *Tone) SetPitch(pitch float64, freq float64) {

	if t.Osc.BaseFreq != pitch {
		t.Osc.Osc.SetFreq((pitch - t.Osc.BaseFreq) + freq)
	}

}
func (t Tone) SendPitch(pitchTest chan float64) {

	t.SetPitch(<-pitchTest, t.Osc.Osc.Freq)

}
