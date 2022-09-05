package library

import (
	"math"
	"math/rand"
)

// WaveType is an alias type for the type of waveforms that can be generated
type WaveType uint16

const (
	WaveSine     WaveType = iota // 0
	WaveTriangle                 // 1
	WaveSaw                      // 2
	WaveSqr                      //3
	WaveNoise                    //4
)

//
const (
	TwoPi = float32(2 * math.Pi)
)

const (
	SineB = 4.0 / math.Pi
	SineC = -4.0 / (math.Pi * math.Pi)
	Q     = 0.775
	SineP = 0.225
)

func Noise(x float32) float32 {
	return (rand.Float32() * -2) + 1
}

// Sine takes an input value from -Pi to Pi
// and returns a value between -1 and 1
func Sine(x32 float32) float32 {
	x := x32
	y := SineB*x + SineC*x*(float32(math.Abs(float64(x))))
	y = SineP*(y*(float32(math.Abs(float64(y))))-y) + y
	return y
}

const TringleA = 2.0 / math.Pi

// Triangle takes an input value from -Pi to Pi
// and returns a value between -1 and 1
func Triangle(x float32) float32 {
	return TringleA*x - 1.0
}

// Square takes an input value from -Pi to Pi
// and returns -1 or 1
func Square(x float32) float32 {
	if x >= 0.0 {
		return 1
	}
	return -1.0
}

const SawtoothA = 1.0 / math.Pi

// Triangle takes an input value from -Pi to Pi
// and returns a value between -1 and 1
func Sawtooth(x float32) float32 {
	return SawtoothA * x
}
