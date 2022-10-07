package dsp

import (
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
	if err != nil {
		log.Fatal(err)
	}
	println(device.Name)
	output := make([]float32, dspConf.BufferSize)

	deviceO := portaudio.StreamDeviceParameters{Device: device, Channels: 1, Latency: device.DefaultHighOutputLatency}

	params := portaudio.StreamParameters{Output: deviceO, SampleRate: device.DefaultSampleRate, FramesPerBuffer: portaudio.FramesPerBufferUnspecified, Flags: portaudio.ClipOff}

	p.Stream, err = portaudio.OpenStream(params, &output)
	if err != nil {
		log.Fatal(err)
	}

	defer p.Stream.Close()

	if err := p.Stream.Start(); err != nil {
		log.Fatal(err)
	}
	defer p.Stream.Stop()
	for {
		if err := p.Stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
		}
		p.FillBuffers()
		p.RunProcess(output)

	}

}
