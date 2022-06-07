package post_audio

import "math"

type Filter struct {
}

const tau = (2 * math.Pi)

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
