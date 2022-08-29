package post_audio

import (
	"math"
)

type Filter struct {
	Cutoff *float64
	Reso   *float64
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
func (filter Filter) RunFilter(input []float32, delay float32, sr float64, fs int) []float32 {

	return Lowpass4(input, *filter.Cutoff, delay, sr, *filter.Reso, fs)

}

//TODO: fix calculations
//MOOG FILTER
func Lowpass(input []float32, freq float64, delay float32, sr float64, resoVal float64, fs int) []float32 {
	var in1, in2, in3, in4, out1, out2, out3, out4 float32
	in1, in2, in3, in4, out1, out2, out3, out4 = 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	output := make([]float32, len(input))
	// copy(output, input)
	// (maxAllowed - minAllowed) * (unscaledNum - min) / (max - min) + minAllowed;
	// newF := float32((freq-0)*(1-0)/5000 + 0) //out 0-1
	newF := float32((freq/float64(fs)/float64(math.Pi*2)-0)*(1-0)/(5000.00/float64(fs)/float64(math.Pi*2)) + 0) //out 0-1
	// newF := freq / 5000
	// newF := freq / 5000

	newR := float32((resoVal-0)*(4-0)/1000 + 0) //out 0-4
	// newR := float32(resoVal / 1000)

	f := float32(newF * 1.16)
	// f := float32(math.Sin(float64(math.Pi*newF)) * 2)
	// f := float32(newF*0.1) * 1.16
	// k := newR * (1 - 0.15*f*f)
	fb := newR * (1 - 0.15*f*f)
	// fb := float32(2*math.Sin(float64(k)*math.Pi/2) - 1)
	magik := float32(0.3)
	magik2 := float32(0.35013)

	// println(newF)
	for i := range output {

		input[i] -= (out4 * fb) + in4
		input[i] *= (magik2 * (f * f) * (f * f)) //0.85013 * (f * f) * (f * f)

		in1 = input[i]
		out1 = (output[i]+magik)*in1 + (1-f)*out1

		in2 = out1 - in1
		out2 = (out1 + magik*in2) + (1-f)*out2
		in3 = out2 - in2
		out3 = (out2 + magik*in3) + (1-f)*out3
		in4 = out3 - in3
		out4 = (out3 + magik*in4) + (1-f)*out4

		// out4 *= (1 - f)
		// output[i] = input[i] - out4 - out3 - out2 - out1
		// output[i] = (out4 + magik*lastOut) + output[i]
		// lastOut = out4

		output[i] = out4
		// output[i] = 3.0 * (out3 - out4)
		// output = Lowpass3(input, 4400, 0.00000001, float32(sr))

	}
	return output
}
func Lowpass4(input []float32, freq float64, delay float32, sr float64, resoVal float64, fs int) []float32 {
	cutoff := float32((freq-0)*(0.99+0.1)/10 + 0.1)
	resonance := float32((resoVal-0)*(1-0)/1000 + 0)
	var buf0, buf1, buf2, buf3 float32
	buf0, buf1, buf2, buf3 = 0.0, 0.0, 0.0, 0.0
	output := make([]float32, len(input))
	feedbackAmount := resonance + resonance/(1.0-cutoff)
	for i := range input {
		buf0 += cutoff * (input[i] - buf0 + feedbackAmount*(buf0-buf1))
		buf1 += cutoff * (buf0 - buf1)
		buf2 += cutoff * (buf1 - buf2)
		buf3 += cutoff * (buf2 - buf3)
		output[i] = buf3
		//buf3 lp, input-buf3 hp, buf0-buf3 bandpass
	}
	return output
}

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
