package generator

// import "time"

type Adsr struct {
	AttackTime  float64
	DecayTime   float64
	SustainAmp  float64
	ReleaseTime float64
	MinValue    float64
	MaxValue    float64
	Type        AdsrType
}
type AdsrType int64
type AdsrAction int64

const (
	IncreaseAction AdsrAction = iota
	DecreaseAction
)
const (
	EnvelopeAdsr AdsrType = iota
	FilterAdsr
	PitchAdsr
)

func (t *Tone) increaseAmp(amp float64) {
	t.Osc.Osc.Amplitude += amp
}
func (t *Tone) decreaseAmp(amp float64) {
	t.Osc.Osc.Amplitude -= amp
}

func (t *Tone) adsrAction(adsrType AdsrType, adsrAction AdsrAction, rate float64) {
	switch adsrAction {
	case IncreaseAction:
		switch adsrType {
		case EnvelopeAdsr:
			t.increaseAmp(rate)
		}
	case DecreaseAction:
		switch adsrType {
		case EnvelopeAdsr:
			t.decreaseAmp(rate)
		}
	}
}

// func ticker() {
// 	ticker := time.NewTicker(1 * time.Millisecond)

// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				// runner(voice, parameter, controller, controlRate, actionType)
// 			case <-voice.Quit:
// 				ticker.Stop()
// 				return
// 			}
// 		}
// 	}()

// }
