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
	playhead      int
}

func (p ProcessAudio) RunProcess(out []float32) {
	//we do blocking way
	p.Mixing()
	for i := range out {
		out[i] = float32(p.audioChannels[0][i]) + float32(p.audioChannels[1][i])
		// out[i] = p.output[(p.playhead+1)%len(p.output)]
	}
	// p.playhead += len(out) % len(p.output) //p.output[(p.playhead+1)%len(p.output)]
	// for a := 0; a < 1024; a++ {
	// 	out[a] = 0
	// }

	// for b := 1; b < 1024; b++ {
	// 	out[len(out)-b] = 0
	// }
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

func (p *ProcessAudio) FillBuffers() {

	for _, v := range p.Voices {

		for _, o := range v.Tones {
			if err := o.Osc.Osc.Fill(o.Osc.Buf); err != nil {
				log.Printf("error filling up the buffer")
			}

		}
	}

}

func (p *ProcessAudio) Mixing() {

	for i, v := range p.Voices {
		// min := v.Lfo.Main.Buf.Data[0]
		// max := v.Lfo.Main.Buf.Data[0]
		// if err := v.Lfo.Main.Osc.Fill(v.Lfo.Main.Buf); err != nil {
		// 	log.Printf("error filling up the buffer")
		// }
		// for k, _ := range v.Lfo.Buf.Data {
		// 	if v.Lfo.Main.Buf.Data[k] > max {
		// 		max = v.Lfo.Main.Buf.Data[k]
		// 	}
		// 	if v.Lfo.Main.Buf.Data[k] < min {
		// 		min = v.Lfo.Main.Buf.Data[k]
		// 	}
		// if *v.Lfo.Rate > 0 {

		// 	*v.ControlValues.Pitch = generator.RescaleThis(v.Lfo.Main.Buf.Data[k])

		// }
		// }

		premix := PreMix(p.audioChannels[i], v.Tones, v)
		filtered := v.Filter.RunFilter(premix, 0.0001, 48000, v.Tones[0].Osc.Osc.Fs)
		// audioChannels[i] = post_audio.Amp(premix, (*v.ControlValues.Vol / 100))
		p.audioChannels[i] = post_audio.Amp(filtered, (*v.ControlValues.Vol / 100))
		// audioChannels = append(audioChannels, audioChannel)

	}

}
