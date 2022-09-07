package dsp

import (
	"hodei/gosynthgo/synth/generator"
	"hodei/gosynthgo/synth/post_audio"

	"log"

	"github.com/gordonklaus/portaudio"
)

//** for now we run just one oscillator
type DspConf struct {
	BufferSize int
}

func RunDSP(dspConf DspConf, voices []*generator.Voice) {

	portaudio.Initialize()
	api, _ := portaudio.HostApis()

	for _, ap := range api {
		log.Println(*ap)
	}
	defer portaudio.Terminate()
	out := make([]float32, dspConf.BufferSize)

	stream, err := portaudio.OpenDefaultStream(0, 2, 48000, len(out), &out)
	if err != nil {
		log.Fatal(err)
	}

	defer stream.Close()

	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}
	defer stream.Stop()
	for {

		fillBuffers(voices)

		Mixing(out, dspConf, voices)
		// Mixing(out, dspConf, voices)
		// write to the stream
		if err := stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
		}

	}

}
func fillBuffers(voices []*generator.Voice) {

	for _, v := range voices {
		for _, o := range v.Tones {
			if err := o.Osc.Osc.Fill(o.Osc.Buf); err != nil {
				log.Printf("error filling up the buffer")
			}

		}
	}

}

//TODO: fix mixing
func Mixing(dst []float32, src DspConf, voices []*generator.Voice) []float32 {

	var audioChannels [][]float64
	buff1 := make([]float64, len(dst))
	buff2 := make([]float64, len(dst))
	audioChannels = append(audioChannels, buff1)
	audioChannels = append(audioChannels, buff2)
	for i, v := range voices {

		premix := PreMix(audioChannels[i], v.Tones, v)
		filtered := v.Filter.RunFilter(premix, 0.0001, 48000, v.Tones[0].Osc.Osc.Fs)
		audioChannels[i] = post_audio.Amp(filtered, (*v.ControlValues.Vol / 100))
		// audioChannels = append(audioChannels, audioChannel)

	}
	// out := make([]float32, len(dst))

	// temp := float32(0.0)
	for i := range dst {

		dst[i] = float32(audioChannels[0][i]) + float32(audioChannels[1][i])
		// dst[i] = out[i]

	}
	// dst = out
	// for _, a := range audioChannels {
	// 	for i := range a {

	// 		dst[i] += a[i]
	// 	}
	// }

	return dst

}

// func Mixing(dst []float32, src DspConf, voices []*generator.Voice) []float32 {

// 	var audioChannels [][]float32

// 	for _, v := range voices {

// 		premix := PreMix(dst, v.Tones, v)
// 		filtered := v.Filter.RunFilter(premix, 0.0001, 48000, v.Tones[0].Osc.Osc.Fs)
// 		audioChannel := post_audio.Amp(filtered, float32(*v.ControlValues.Vol/100))
// 		audioChannels = append(audioChannels, audioChannel)

// 	}

// 	for _, a := range audioChannels {
// 		for i := range a {

// 			dst[i] += a[i]
// 		}
// 	}

// 	return dst

// }
