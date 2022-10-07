package post_audio

import "math"

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

// func normalize(xs []float64) []float64 {
// 	length := len(xs)
// 	maxamp := 0.0
// 	for i := 0; i < length; i++ {
// 		amp := math.Abs(xs[i])
// 		if amp > maxamp {
// 			maxamp = amp
// 		}
// 	}

// 	maxamp = 1.0 * maxamp
// 	for i := 0; i < length; i++ {
// 		xs[i] *= maxamp
// 	}
// 	xs[len(xs)-1] = xs[0]
// 	return xs
// }

func (filter Filter) RunFilter(input []float64, delay float32, sr float64, fs int) []float64 {

	return Lowpass5(input, *filter.Cutoff, delay, sr, *filter.Reso, fs, LP)

}

//(maxOut - minOut) * (unscaledNum - min) / (max - min) + minOut
func Lowpass5(input []float64, freq float64, delay float32, sr float64, resoVal float64, fs int, filterType FilterType) []float64 {

	fNorm := freq / (sr)
	// fNorm := (1-0)*(freq-20)/(6000-20) + 0
	resonance := ((resoVal-0)*(4-0)/1000 + 0)
	var buf0, buf1, buf2, buf3, in1, in2, in3, in4 float64
	buf0, buf1, buf2, buf3, in1, in2, in3, in4 = 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	output := make([]float64, len(input))
	// feedbackAmount := resonance * (1 - (0.15 * cutoff * 1.15 * cutoff * 1.15)) //resonance/(1.0-cutoff)
	// feedbackAmount := resonance + resonance/(1.0-cutoff)

	f := math.Sin(math.Pi * fNorm) // * 1.16
	// f := fNorm * 1.16
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
