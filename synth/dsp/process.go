package dsp

import (
	"hodei/gosynthgo/synth/generator"
	"hodei/gosynthgo/synth/post_audio"
	"log"

	"github.com/gordonklaus/portaudio"
)

type ProcessAudio struct {
	Stream        *portaudio.Stream
	Voices        []*generator.Voice
	out           []float64
	output        []float32
	dspConf       DspConf
	audioChannels [][]float64
}

func (p ProcessAudio) RunProcess(out []float32) {
	// p := NewProcessAudio(out, voices, dspConf)

	go p.fillBuffers()
	go p.mixing()

	for i := range out {
		out[i] = p.output[i]
	}

}

func NewProcessAudio(out *float64, voices []*generator.Voice, dspConf DspConf) *ProcessAudio {
	var audioChannels [][]float64
	// audioChannels := make([][]float64, 2)
	buf1 := make([]float64, dspConf.BufferSize)
	buf2 := make([]float64, dspConf.BufferSize)
	audioChannels = append(audioChannels, buf1)
	audioChannels = append(audioChannels, buf2)
	return &ProcessAudio{Voices: voices, out: make([]float64, dspConf.BufferSize), output: make([]float32, dspConf.BufferSize), dspConf: dspConf, audioChannels: audioChannels}
}
func (p *ProcessAudio) fillBuffers() {

	for _, v := range p.Voices {
		for _, o := range v.Tones {
			if err := o.Osc.Osc.Fill(o.Osc.Buf); err != nil {
				log.Printf("error filling up the buffer")
			}

		}
	}

}

func (p *ProcessAudio) mixing() {

	// var audioChannels [][]float64
	// audioChannels[0]=make([]float64, p.dspConf.BufferSize)
	// audioChannels[1]=make([]float64, p.dspConf.BufferSize)
	// buff1 := make([]float64, p.dspConf.BufferSize)
	// buff2 := make([]float64, p.dspConf.BufferSize)
	// audioChannels = append(audioChannels, buff1)
	// audioChannels = append(audioChannels, buff2)
	// for i, v := range p.Voices {

	// premix := PreMix(p.audioChannels[0], p.Voices[0].Tones, p.Voices[0])
	// premix1 := PreMix(p.audioChannels[1], p.Voices[1].Tones, p.Voices[1])
	// filtered := p.Voices[0].Filter.RunFilter(premix, 0.0001, 48000, p.Voices[0].Tones[0].Osc.Osc.Fs)
	// filtered1 := p.Voices[0].Filter.RunFilter(premix1, 0.0001, 48000, p.Voices[0].Tones[0].Osc.Osc.Fs)
	// // audioChannels[i] = post_audio.Amp(premix, (*v.ControlValues.Vol / 100))
	// post_audio.Amp(filtered, (*p.Voices[0].ControlValues.Vol / 100))
	// post_audio.Amp(filtered1, (*p.Voices[1].ControlValues.Vol / 100))
	// // audioChannels = append(audioChannels, audioChannel)

	// }
	// var audioChannels [][]float64
	// buff1 := make([]float64, p.dspConf.BufferSize)
	// buff2 := make([]float64, p.dspConf.BufferSize)
	// audioChannels = append(audioChannels, buff1)
	// audioChannels = append(audioChannels, buff2)
	for i, v := range p.Voices {

		premix := PreMix(p.audioChannels[i], v.Tones, v)
		// filtered := v.Filter.RunFilter(premix, 0.0001, 48000, v.Tones[0].Osc.Osc.Fs)
		// audioChannels[i] = post_audio.Amp(premix, (*v.ControlValues.Vol / 100))
		p.audioChannels[i] = post_audio.Amp(premix, (*v.ControlValues.Vol / 100))
		// audioChannels = append(audioChannels, audioChannel)

	}

	for i := range p.out {

		p.output[i] = float32(p.audioChannels[0][i] + p.audioChannels[1][i])
		// dst[i] = out[i]

	}

}
