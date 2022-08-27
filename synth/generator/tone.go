package generator

import (
	"fmt"
	"hodei/gosynthgo/synth/midi"
	"time"
)

type Tone struct {
	Key      int
	IsOn     bool
	Osc      Osc
	Type     OscType
	Gain     float64
	FramePos float64
	Active   bool
	StopTime chan bool
	Vel      float64
}

func NewTone(buffer int, oscType OscType) Tone {
	switch oscType {
	case Regular:
		return Tone{Key: 0, IsOn: false, Osc: Oscillator(buffer), Type: Regular, FramePos: 0, StopTime: make(chan bool, 1), Active: false}
	case Noize:
		return Tone{Key: 0, IsOn: false, Osc: NoiseOsc(buffer), Type: Noize, FramePos: 0, StopTime: make(chan bool, 1), Active: false}
	default:
		return Tone{Key: 0, IsOn: false, Osc: Oscillator(buffer), Type: Regular}

	}
}
func (t *Tone) BindToOSC(message midi.MidiMsg, adsr *Adsr) {
	t.Key = message.Key
	t.IsOn = message.On
	t.Gain = t.Osc.Osc.Amplitude
	t.Osc.ChangeFreq(message)

	if t.IsOn {
		t.Vel = RescaleMidiValues(message.Vel, 0.0, OscMaxAmp) //it gets velocity value
		// if !t.Active {
		t.Active = true
		go startTimer(t, adsr)
		// }

	} else {
		// t.Active = false
		// if t.Active {
		t.StopTime <- true
		t.FramePos = 0.0

		go startTimer(t, adsr)
		// }
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

func (t Tone) ShowStatus() string {
	return fmt.Sprint("STATUS->  ", t.IsOn, t.Key)
}

//TODO: refactor to ADSR all timing function
func startTimer(t *Tone, adsr *Adsr) {
	ticker := time.NewTicker(time.Duration(1) * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			runOnTime(t, adsr)
		case <-t.StopTime:
			t.FramePos = 0.0
			// t.Osc.Osc.Amplitude = 0.0
			ticker.Stop()
			return
		}
	}

}

func runOnTime(t *Tone, adsr *Adsr) {

	t.FramePos += 1
	adsr.RunAdsr(t, EnvelopeAdsr, t.Vel)

}
