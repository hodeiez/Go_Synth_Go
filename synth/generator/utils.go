package generator

import (
// "os"
// "os/signal"

// "github.com/go-audio/audio"
// "github.com/go-audio/generator"
)

type Pwm struct {
}

// func NoiseOsc(bufferSize int) Osc {

// 	buf := &audio.FloatBuffer{
// 		Data:   make([]float64, bufferSize),
// 		Format: audio.FormatStereo44100,
// 	}

// 	osc := generator.NewOsc(generator.WaveNoise, 440.0, buf.Format.SampleRate)
// 	osc.Amplitude = 1
// 	sig := make(chan os.Signal, 1)
// 	signal.Notify(sig, os.Interrupt, os.Kill)
// 	println("noize running")
// 	return Osc{Osc: osc, Buf: buf}

// }
