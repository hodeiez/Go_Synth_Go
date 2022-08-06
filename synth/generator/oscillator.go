package generator

import (
	"hodei/gosynthgo/synth/midi"
	"log"
	"math"

	"os"
	"os/signal"

	"github.com/go-audio/audio"
	"github.com/go-audio/generator"
)

type Osc struct {
	// gainControl float64
	Osc      *generator.Osc
	Buf      *audio.FloatBuffer
	BaseFreq float64
}
type OscType int64

const (
	Regular OscType = iota
	Noize
	None
)

type MyWaveType int64

const (
	Triangle MyWaveType = iota
	Saw
	Square
	Sine
	MyWaveTypeSize
)

func (s MyWaveType) String() string {

	switch s {
	case Triangle:
		return "Triangle"
	case Saw:
		return "Saw"
	case Sine:
		return "Sine"
	case Square:
		return "Square"
	}

	return "-"
}
func NoiseOsc(bufferSize int) Osc {

	buf := &audio.FloatBuffer{
		Data:   make([]float64, bufferSize),
		Format: audio.FormatStereo44100,
	}

	osc := generator.NewOsc(generator.WaveNoise, 440.0, buf.Format.SampleRate)
	osc.Amplitude = 0.0
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	println("noize running")
	return Osc{Osc: osc, Buf: buf, BaseFreq: 440.0}

}
func Oscillator(bufferSize int) Osc {
	// this has to go to a preconf**************

	buf := &audio.FloatBuffer{
		Data:   make([]float64, bufferSize),
		Format: audio.FormatStereo44100,
	}
	//***************************

	currentNote := 440.0
	osc := generator.NewOsc(generator.WaveSaw, currentNote, buf.Format.SampleRate)
	osc.Amplitude = 0.0
	osc.Freq = 440.0
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	log.Println("oscillator running")
	return Osc{osc, buf, 440.0}

}
func (osc *Osc) SetBaseFreq(freq float64) {
	osc.BaseFreq = freq
}
func (osc *Osc) ChangeFreq(midimsg midi.MidiMsg) {

	NoteToPitch := (osc.BaseFreq / 32) * (math.Pow(2, ((float64(midimsg.Key) - 9) / 12)))

	if midimsg.On {
		osc.Osc.SetFreq(NoteToPitch)
	}
	//	return *osc
}

// func ChangeFreq(midimsg midi.MidiMsg, osc *Osc) Osc {

// 	NoteToPitch := (osc.BaseFreq / 32) * (math.Pow(2, ((float64(midimsg.Key) - 9) / 12)))

// 	if midimsg.On {
// 		osc.Osc.SetFreq(NoteToPitch)
// 	}
// 	return *osc
// }
func (o *Osc) ChangeFreq2(midimsg midi.MidiMsg) {

	NoteToPitch := (o.BaseFreq / 32) * (math.Pow(2, ((float64(midimsg.Key) - 9) / 12)))

	o.Osc.SetFreq(NoteToPitch)
}

func SelectWave(waveName int, o Osc) {

	switch waveName {
	case 0:
		o.Osc.Shape = generator.WaveType(generator.WaveSaw)
	case 1:
		o.Osc.Shape = generator.WaveType(generator.WaveTriangle)
	case 2:
		o.Osc.Shape = generator.WaveType(generator.WaveSqr)
	case 3:
		o.Osc.Shape = generator.WaveType(generator.WaveSine)

	}

}
