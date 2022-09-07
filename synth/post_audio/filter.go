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
	// oversamped := oversampling(input, 2000, delay, sr*2, 80, fs)
	// filtered := Lowpass4(oversamped, *filter.Cutoff, delay, sr*2, *filter.Reso, fs, LP)
	// return downSamp(filtered)
	return Lowpass4(input, *filter.Cutoff, delay, sr, *filter.Reso, fs, LP)
	// return Lowpass4(input, 3230, delay, sr, 0, fs, BP)
	//return Lowpass3(input, *filter.Cutoff, 0.1, float32(sr))
	//6000 380
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

// func oversampling(input []float32, freq float64, delay float32, sr float64, reso float64, fs int) []float32 {
// 	oversamp := make([]float32, 2*len(input))
// 	for a := 0; a < len(input)-1; a++ {
// 		oversamp[a] = input[a]
// 		oversamp[a+1] = input[a]
// 	}
// 	// Lowpass4(oversamped, *filter.Cutoff, delay, sr, *filter.Reso, fs)
// 	return Lowpass4(oversamp, freq, delay, sr, reso, fs, BP)
// 	// return oversamp
// }
func downSamp(oversampOut []float32) []float32 {
	output := make([]float32, len(oversampOut)/2)

	for b := 0; b < (len(oversampOut)/2)-1; b++ {
		output[b] = oversampOut[b]
		output[b+1] = oversampOut[b+2]
	}
	return output
}

// func oversampling(input []float32) []float32 {
// 	oversamp := make([]float32, 2*len(input))
// 	for a := 0; a < len(input)-1; a++ {
// 		oversamp[a] = input[a]
// 		oversamp[a+1] = input[a]
// 	}
// 	return oversamp
// }
// func downSamp(oversampOut []float32) []float32 {
// 	output := make([]float32, len(oversampOut)/2)

// 	for b := 0; b < (len(oversampOut)/2)-1; b++ {
// 		output[b] = oversampOut[b]
// 		output[b+1] = oversampOut[b+2]
// 	}
// 	return output
// }

// func Lowpass2(input []float32, freq float64, delay float32, sr float64, resoVal float64) []float32 {
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

// func Bandpass(input []float32, freq float64, delay float32, sr float64, q float64) []float32 {
// 	// 200-300
// 	//	q := 20.
// 	//return Highpass(input, freq-q, delay, sr) - Lowpass(input, freq, delay, sr)
// 	return Lowpass(Highpass(input, freq-q, delay, sr), freq, delay, sr, 100)
// }

func Lowpass3(input []float32, freq float64, delay, sr float32) []float32 {
	output := make([]float32, len(input))
	copy(output, input)

	costh := 2. - float32(math.Cos(float64(tau*freq)))/sr
	coef := float32(math.Sqrt(float64(costh*costh-1.))) - costh

	for i, a := range output {
		output[i] = a*(1+coef) - delay*coef
		delay = output[i]
	}

	return output
}
