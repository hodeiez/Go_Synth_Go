package generator

import (
// "os"
// "os/signal"

// "github.com/go-audio/audio"
// "github.com/go-audio/generator"
)

type Lfo struct {
	Rate *float64
	// Width float64
	Main Osc
}

// func NewLFO(bufferSize int, lfoRate *float64) *Lfo {
// 	buf := &audio.FloatBuffer{
// 		Data:   make([]float64, bufferSize),
// 		Format: audio.FormatStereo48000,
// 	}

// 	osc := generator.NewOsc(generator.WaveNoise, 40.0, buf.Format.SampleRate)
// 	osc.Amplitude = 0.1

// 	sig := make(chan os.Signal, 1)
// 	signal.Notify(sig, os.Interrupt, os.Kill)

// 	wraped := Osc{Osc: osc, Buf: buf, BaseFreq: 40.0}
// 	return &Lfo{Main: wraped, Rate: lfoRate}

// }
// func (lfo *Lfo) ChangeLfoRate(val float64) {
// 	lfo.Main.PitchShift(val)
// }
