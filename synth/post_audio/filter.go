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

//MOOG FILTER
func Lowpass(input []float32, freq float64, delay float32, sr float64, resoVal float64) []float32 {
	output := make([]float32, len(input))
	copy(output, input)
	newF := freq / 5000
	newR := float32(resoVal / 1000)
	in1 := float32(0.0)
	in2 := float32(0.0)
	in3 := float32(0.0)
	in4 := float32(0.0)
	out1 := float32(0.0)
	out2 := float32(0.0)
	out3 := float32(0.0)
	out4 := float32(0.0)
	f := float32(newF * 1.16)
	fb := newR * (1 - 0.15*f*f)

	for i, _ := range output {
		input[i] -= out4 * fb
		input[i] *= 0.35013 * (f * f) * (f * f)
		out1 = input[i] + 0.3*in1 + (1-f)*out1
		in1 = input[i]
		out2 = out1 + 0.3*in2 + (1-f)*out2
		in2 = out1
		out3 = out2 + 0.3*in3 + (1-f)*out3
		in3 = out2
		out4 = out3 + 0.3*in4 + (1-f)*out4
		in4 = out3
		output[i] = out4

	}

	return output
}

// func Lowpass(input []float32, freq float64, delay float32, sr float64, resoVal float64) []float32 {
// 	output := make([]float32, len(input))
// 	copy(output, input)
// 	q := float32(resoVal) / 100
// 	costh := 2. - math.Cos((tau*freq)/sr)
// 	coef := float32(math.Sqrt(costh*costh-1.)) - float32(costh)

// 	for i, a := range output {

// 		output[i] = a*(q)*(1+coef) - delay*coef
// 		delay = output[i]

// 	}

// 	return output
// }

func Bandpass(input []float32, freq float64, delay float32, sr float64, q float64) []float32 {
	// 200-300
	//	q := 20.
	//return Highpass(input, freq-q, delay, sr) - Lowpass(input, freq, delay, sr)
	return Lowpass(Highpass(input, freq-q, delay, sr), freq, delay, sr, 100)
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
