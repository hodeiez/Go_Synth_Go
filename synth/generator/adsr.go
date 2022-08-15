package generator

type Adsr struct {
	AttackTime  *int32
	DecayTime   *int32
	SustainAmp  *int32
	ReleaseTime *int32
	MinValue    float64
	MaxValue    float64
	Type        AdsrType
}
type AdsrType int64
type AdsrAction int64

const (
	IncreaseAction AdsrAction = iota
	DecreaseAction
	SustainAction
)
const (
	EnvelopeAdsr AdsrType = iota
	FilterAdsr
	PitchAdsr
)

func (t *Tone) increaseAmp(amp float64, maxValue float64) {
	if t.Osc.Osc.Amplitude < maxValue {
		t.Osc.Osc.Amplitude += amp
	}
}
func (t *Tone) decreaseAmp(amp float64, minValue float64) {
	if t.Osc.Osc.Amplitude > minValue {
		t.Osc.Osc.Amplitude -= amp
	}
}
func (t *Tone) sustainAmp(amp float64) {
	t.Osc.Osc.Amplitude = amp / 1000
}

func (t *Tone) adsrAction(adsrType AdsrType, adsrAction AdsrAction, rate float64, controlValue float64) {
	switch adsrAction {
	case IncreaseAction:
		switch adsrType {
		case EnvelopeAdsr:
			t.increaseAmp(rate, controlValue)
		}
	case DecreaseAction:
		switch adsrType {
		case EnvelopeAdsr:
			t.decreaseAmp(rate, controlValue)
		}
	case SustainAction:
		switch adsrType {
		case EnvelopeAdsr:
			t.sustainAmp(rate)
		}
	}
}
func (adsr Adsr) RunAdsr(t *Tone, adsrType AdsrType, rateValue float64) {
	if t.IsOn {
		if t.FramePos < float64(*adsr.AttackTime) {
			t.adsrAction(adsrType, IncreaseAction, rateValue, adsr.MaxValue)
		} else if t.FramePos > float64(*adsr.AttackTime) && t.FramePos > float64(*adsr.AttackTime+*adsr.DecayTime) {

			t.adsrAction(adsrType, DecreaseAction, rateValue, float64(*adsr.SustainAmp/1000))
		} else {
			t.adsrAction(adsrType, SustainAction, float64(*adsr.SustainAmp), 0.0)
		}

	} else {
		t.adsrAction(adsrType, DecreaseAction, rateValue, adsr.MinValue)
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
