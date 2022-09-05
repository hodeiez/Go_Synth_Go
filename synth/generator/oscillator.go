package generator

import (
	"hodei/gosynthgo/synth/library"
	"hodei/gosynthgo/synth/midi"
	"log"
	"math"

	"os"
	"os/signal"

	"github.com/go-audio/audio"
	// "github.com/go-audio/generator"
)

type Osc struct {
	// gainControl float32
	Osc      *library.Osc
	Buf      *audio.Float32Buffer
	BaseFreq float32
}
type OscType int64

const (
	Regular OscType = iota
	Noize
	None
)
const (
	OscMaxAmp = 0.03
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

	buf := &audio.Float32Buffer{
		Data:   make([]float32, bufferSize),
		Format: audio.FormatStereo48000,
	}

	osc := library.NewOsc(library.WaveNoise, 440.0, buf.Format.SampleRate)
	osc.Amplitude = 0.0

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	println("noize running")
	return Osc{Osc: osc, Buf: buf, BaseFreq: 440.0}

}
func Oscillator(bufferSize int) Osc {
	// this has to go to a preconf**************

	buf := &audio.Float32Buffer{
		Data:   make([]float32, bufferSize),
		Format: audio.FormatStereo48000,
	}
	//***************************

	currentNote := float32(40.0)
	osc := library.NewOsc(library.WaveSaw, currentNote, buf.Format.SampleRate)

	osc.Amplitude = 0.0
	// osc.Freq = 440.0

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	log.Println("oscillator running")
	return Osc{osc, buf, 440.0}

}
func (osc *Osc) SetBaseFreq(freq float32) {
	osc.BaseFreq = freq
}
func (osc *Osc) ChangeFreq(midimsg midi.MidiMsg) {

	NoteToPitch := (osc.BaseFreq / 32) * (float32(math.Pow(float64(2), float64((float32(midimsg.Key)-9)/12))))

	osc.Osc.SetFreq(NoteToPitch)

}
func (osc *Osc) PitchShift(val float32) {
	temp := osc.Osc.Freq + val
	osc.Osc.Freq = temp

}

func SelectWave(waveName int, o Osc) {

	switch waveName {
	case 0:
		o.Osc.Shape = library.WaveType(library.WaveSaw)
	case 1:
		o.Osc.Shape = library.WaveType(library.WaveTriangle)
	case 2:
		o.Osc.Shape = library.WaveType(library.WaveSqr)
	case 3:
		o.Osc.Shape = library.WaveType(library.WaveSine)

	}

}
