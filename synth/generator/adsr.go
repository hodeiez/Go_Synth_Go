package generator

const (
	tenMsconds = 10
)

type Adsr struct {
	AttackTime  *int32
	DecayTime   *int32
	SustainAmp  *int32
	ReleaseTime *int32
	MinValue    float32
	MaxValue    float32
	Type        AdsrType
	StopTime    chan bool
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

func (t *Tone) increaseAmp(amp float32, maxValue float32) {
	if t.Osc.Osc.Amplitude < maxValue {

		t.Osc.Osc.Amplitude += amp
	}
}
func (t *Tone) decreaseAmp(amp float32, minValue float32) {
	if t.Osc.Osc.Amplitude > minValue {
		t.Osc.Osc.Amplitude -= amp
	}
}
func (t *Tone) sustainAmp(amp float32) {
	t.Osc.Osc.Amplitude = amp
}

func (t *Tone) adsrAction(adsrType AdsrType, adsrAction AdsrAction, rate float32, controlValue float32) {
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
			t.sustainAmp(controlValue)
		}
	}
}

//TODO: normalize values
func (adsr Adsr) RunAdsr(t *Tone, adsrType AdsrType, gain float32) {

	attackT := RescaleToMilliSeconds(*adsr.AttackTime, 0, 100, 10)

	decayT := RescaleToMilliSeconds(*adsr.DecayTime, 0, 100, 10)
	releaseT := RescaleToMilliSeconds(*adsr.ReleaseTime, 0, 100, 10)

	if t.IsOn {

		if attackT == 0.0 && t.FramePos < tenMsconds {

			t.adsrAction(adsrType, SustainAction, 0.0, gain)
		}
		if t.FramePos <= attackT {

			t.adsrAction(adsrType, IncreaseAction, getRateValue(attackT, gain), gain)

		} else if t.FramePos >= attackT && t.FramePos < attackT+decayT {

			t.adsrAction(adsrType, DecreaseAction, getRateValue(decayT, gain), getSustainAmpValue(gain, *adsr.SustainAmp))
		} else if t.FramePos >= attackT+decayT {

			t.adsrAction(adsrType, SustainAction, 0.0, getSustainAmpValue(gain, *adsr.SustainAmp))
		}

	} else {

		if t.FramePos >= releaseT {
			t.FramePos = 0.0
			t.Active = false
			t.StopTime <- true
		} else {
			t.adsrAction(adsrType, DecreaseAction, getRateValue(releaseT, gain), 0.0)
		}

	}

}

func getSustainAmpValue(maxValue float32, sustainValue int32) float32 {
	sustainMax := float32(100.00)
	return maxValue * (float32(sustainValue) / sustainMax)
}
func getRateValue(timeScale float32, rateValue float32) float32 {
	if timeScale == RescaleToMilliSeconds(0, 0, 100, 10) { //if time is set to 0, then rateValue is full
		return rateValue
	} else {
		return rateValue / timeScale
	}

}
