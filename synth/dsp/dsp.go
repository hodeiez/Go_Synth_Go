package dsp

import (
	"hodei/gosynthgo/synth/generator"
	"hodei/gosynthgo/synth/post_audio"
	"time"

	"log"

	"github.com/gordonklaus/portaudio"
)

type DspConf struct {
	BufferSize int
}
type MySound struct {
	*portaudio.Stream
}

//TODO: optimize dsp stream and use callback, move all the audio process to its own thread and use dsp thread for just streaming
func RunDSP(dspConf DspConf, voices []*generator.Voice) {

	portaudio.Initialize()
	api, _ := portaudio.HostApis()

	for _, ap := range api {
		log.Println(*ap)
	}
	defer portaudio.Terminate()
	out := make([]float32, dspConf.BufferSize)
	device, err := portaudio.DefaultOutputDevice()
	// outTemp := make([]float32, dspConf.BufferSize*4)
	deviceO := portaudio.StreamDeviceParameters{Device: device, Channels: 1, Latency: time.Duration(000000000)}
	//  deviceI := portaudio.StreamDeviceParameters{Device: device, Channels: 0, Latency: time.Duration(1000000000)}
	// highL := portaudio.LowLatencyParameters(nil, device)
	params := portaudio.StreamParameters{Output: deviceO, SampleRate: 48000, FramesPerBuffer: dspConf.BufferSize, Flags: portaudio.ClipOff}

	// stream, err := portaudio.OpenStream(highL, &out)
	stream, err := portaudio.OpenStream(params, &out)

	//  stream, err := portaudio.OpenDefaultStream(0, 1, 48000, len(out), &out)
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
func Mixing(dst []float32, src DspConf, voices []*generator.Voice) {

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

	for i := range dst {

		dst[i] = float32(audioChannels[0][i]) + float32(audioChannels[1][i])
		// dst[i] = out[i]

	}

	// dst[1]=dst[len(dst)-2]

}
