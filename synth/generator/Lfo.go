package generator

import (
	"os"
	"os/signal"

	"github.com/go-audio/audio"
	"github.com/go-audio/generator"
)

type Lfo struct {
	Rate *float64
	Buf  audio.FloatBuffer
	Main Osc
}

func NewLFO(bufferSize int, lfoRate *float64) *Lfo {
	buf := &audio.FloatBuffer{
		Data:   make([]float64, bufferSize),
		Format: audio.FormatStereo48000,
	}

	osc := generator.NewOsc(generator.WaveSine, 10.0, buf.Format.SampleRate)
	osc.Amplitude = 0.1

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	wraped := Osc{Osc: osc, Buf: buf, BaseFreq: 4.0}
	return &Lfo{Main: wraped, Rate: lfoRate}

}
func (lfo *Lfo) ChangeLfoRate(val float64) {
	lfo.Main.PitchShift(val)
}
func (lfo *Lfo) LfoToHz() float64 {
	return lfo.Main.Osc.Sample() * 100
}

/*
1.-Panel set lfo rate
2.-get lfo sample
3.-map each sample to oscillators pitch
*/
