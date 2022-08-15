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
	//TODO: change this when ADSR is on

	if t.IsOn {
		if !t.Active {
			go startTimer(t)
			t.Active = true
		}
		// println(t.FramePos)

		if t.Type == Regular {
			t.Osc.Osc.Amplitude = RescaleMidiValues(message.Vel, 0.0, 0.1) //it gets velocity value
			//adsr.MaxValue = t.Osc.Osc.Amplitude
			//run adsr, get adsr values?? put tone,
			// t.Osc.Osc.Amplitude = 0.01 * float64(message.Vel)
		} else {
			// t.Osc.Osc.Amplitude = RescaleMidiValues(message.Vel, 0.0, 0.1) //it gets velocity value

			t.Osc.Osc.Amplitude = 0.01
		}
	} else {
		t.Active = false
		// println(t.FramePos)
		t.StopTime <- true
		t.FramePos = 0.0
		t.Osc.Osc.Amplitude = 0.0000
	}
	// adsr.RunAdsr(t, EnvelopeAdsr, 0.00001)
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

func startTimer(t *Tone) {
	ticker := time.NewTicker(time.Duration(1) * time.Millisecond)

	// go func() {
	for {
		select {
		case <-ticker.C:
			runOnTime(t)
		case <-t.StopTime:
			ticker.Stop()
			return
		}
	}
	// }()

}

func runOnTime(t *Tone) {
	// for {
	t.FramePos += 1
	// println(t.FramePos)
	// t.Active <- false
	// println(t.FramePos)
	// }
}
