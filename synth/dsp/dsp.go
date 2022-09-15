package dsp

import (
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

func RunDSP(p *ProcessAudio, dspConf DspConf) {

	portaudio.Initialize()

	defer portaudio.Terminate()

	device, err := portaudio.DefaultOutputDevice()

	output := make([]float32, dspConf.BufferSize)
	deviceO := portaudio.StreamDeviceParameters{Device: device, Channels: 1, Latency: time.Duration(0)}

	params := portaudio.StreamParameters{Output: deviceO, SampleRate: 48000, FramesPerBuffer: dspConf.BufferSize, Flags: portaudio.PrimeOutputBuffersUsingStreamCallback}

	p.Stream, err = portaudio.OpenStream(params, &output)
	// p.Stream = stream
	if err != nil {
		log.Fatal(err)
	}

	defer p.Stream.Close()

	if err := p.Stream.Start(); err != nil {
		log.Fatal(err)
	}
	defer p.Stream.Stop()
	for {

		p.RunProcess(output)

		if err := p.Stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
		}

	}

}
