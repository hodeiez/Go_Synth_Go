package post_audio

import (
	"math"
)

type Filter struct {
}

const tau = (2 * math.Pi)

//TODO: fix filters and add resonance (q)
func Highpass(fs []float32, freq float64, delay float32, sr float64) []float32 {
	output := make([]float32, len(fs))
	copy(output, fs)

	b := 2. - math.Cos(tau*freq/sr)
	coef := b - math.Sqrt(b*b-1.)

	for i, a := range output {
		output[i] = a*(1.-float32(coef)) - delay*float32(coef)
		delay = output[i]
	}

	return output
}
func Lowpass(input []float32, freq float64, delay float32, sr float64) []float32 {
	output := make([]float32, len(input))
	copy(output, input)

	costh := 2. - math.Cos((tau*freq)/sr)
	coef := float32(math.Sqrt(costh*costh-1.)) - float32(costh)

	for i, a := range output {

		output[i] = a*(1+coef) - delay*coef
		delay = output[i]

	}

	return output
}

func Bandpass(input []float32, freq float64, delay float32, sr float64, q float64) []float32 {
	// 200-300
	//	q := 20.
	//return Highpass(input, freq-q, delay, sr) - Lowpass(input, freq, delay, sr)
	return Lowpass(Highpass(input, freq-q, delay, sr), freq, delay, sr)
}

// func Lowpass(input []float64, freq, delay, sr float64) []float64 {
// 	output := make([]float64, len(input))
// 	copy(output, input)

// 	costh := 2. - math.Cos((tau*freq)/sr)
// 	coef := math.Sqrt(costh*costh-1.) - costh

// 	for i, a := range output {
// 		output[i] = a*(1+coef) - delay*coef
// 		delay = output[i]
// 	}

// 	return output
// }
