package post_audio

import (
	"math"
)

type Filter struct {
	Cutoff *float64
	Reso   *float64
}
type FilterType int64

const (
	LP FilterType = iota
	HP
	BP
	WTF
)

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
func (filter Filter) RunFilter(input []float64, delay float32, sr float64, fs int) []float64 {

	return Lowpass5(input, *filter.Cutoff, delay, sr, *filter.Reso, fs, LP)

}

func Lowpass4(input []float64, freq float64, delay float32, sr float64, resoVal float64, fs int, filterType FilterType) []float64 {

	freqC := 2.0 * math.Sin(math.Pi*freq/float64(fs))
	cutoff := (freqC)

	resonance := ((resoVal-0)*(1-0)/1000 + 0)
	var buf0, buf1, buf2, buf3 float64
	buf0, buf1, buf2, buf3 = 0.0, 0.0, 0.0, 0.0
	output := make([]float64, len(input)) //EZ AHAZTU
	// feedbackAmount := resonance * (1 - (0.15 * cutoff * 1.15 * cutoff * 1.15)) //resonance/(1.0-cutoff)
	feedbackAmount := resonance + resonance/(1.0-cutoff)

	for i := range input {
		buf0 += cutoff * (input[i] - buf0 + feedbackAmount*(buf0-buf1))
		buf1 += cutoff * (buf0 - buf1)
		buf2 += cutoff * (buf1 - buf2)
		buf3 += cutoff * (buf2 - buf3)

		switch filterType {
		case LP:
			output[i] = buf3
		case HP:
			output[i] = input[i] - buf3
		case BP:
			output[i] = buf0 - buf3
		case WTF:
			output[i] = buf0 - input[i] // - buf3/2 + buf0/2
		}

		//buf3 lp, input-buf3 hp, buf0-buf3 bandpass

	}

	return output

}

//(maxOut - minOut) * (unscaledNum - min) / (max - min) + minOut
func Lowpass5(input []float64, freq float64, delay float32, sr float64, resoVal float64, fs int, filterType FilterType) []float64 {

	fNorm := (1-0)*(freq-20)/(6000-20) + 0
	resonance := ((resoVal-0)*(4-0)/1000 + 0)
	var buf0, buf1, buf2, buf3, in1, in2, in3, in4 float64
	buf0, buf1, buf2, buf3, in1, in2, in3, in4 = 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	output := make([]float64, len(input)) //EZ AHAZTU
	// feedbackAmount := resonance * (1 - (0.15 * cutoff * 1.15 * cutoff * 1.15)) //resonance/(1.0-cutoff)
	// feedbackAmount := resonance + resonance/(1.0-cutoff)

	f := fNorm * 1.16
	fb := resonance * (1 - 0.15*f*f)
	for i := range input {
		input[i] -= buf3 * fb
		input[i] *= 0.35013 * (f * f) * (f * f)
		buf0 = input[i] + 0.3*in1 + (1-f)*buf0
		in1 = input[i]
		buf1 = buf0 + 0.3*in2 + (1-f)*buf1
		in2 = buf0
		buf2 = buf1 + 0.3*in3 + (1-f)*buf2
		in3 = buf1
		buf3 = buf2 + 0.3*in4 + (1-f)*buf3
		in4 = buf2

		switch filterType {
		case LP:
			output[i] = buf3
		case HP:
			output[i] = input[i] - buf3
		case BP:
			output[i] = buf0 - buf3
		case WTF:
			output[i] = buf0 - input[i] // - buf3/2 + buf0/2
		}

		// output[1]=input[1]
		//buf3 lp, input-buf3 hp, buf0-buf3 bandpass

	}
	return output
}
