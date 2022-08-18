package generator

const (
	tenMsconds = 10
)

type Adsr struct {
	AttackTime  *int32
	DecayTime   *int32
	SustainAmp  *int32
	ReleaseTime *int32
	MinValue    float64
	MaxValue    float64
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
	// println(amp)
	t.Osc.Osc.Amplitude = amp
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
			t.sustainAmp(controlValue)
		}
	}
}

//TODO: normalize values, add release ticker, fix
func (adsr Adsr) RunAdsr(t *Tone, adsrType AdsrType, gain float64) {
	// println(*adsr.AttackTime)
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
		t.adsrAction(adsrType, DecreaseAction, getRateValue(releaseT, gain), 0.0)
		println(t.FramePos)
		if t.FramePos >= releaseT {

			t.StopTime <- true
		}
		//t.adsrAction(adsrType, DecreaseAction, getRateValue(RescaleToMilliSeconds(*adsr.ReleaseTime, 0, 100, 10), rateValue), 0.0)
		// t.adsrAction(adsrType, DecreaseAction, rateValue, adsr.MinValue)
	}

}

func getSustainAmpValue(maxValue float64, sustainValue int32) float64 {
	sustainMax := 100.00
	return maxValue * (float64(sustainValue) / sustainMax)
}
func getRateValue(timeScale float64, rateValue float64) float64 {
	if timeScale == RescaleToMilliSeconds(0, 0, 100, 10) {
		return rateValue
	} else {
		return rateValue / timeScale
	}
	/*
		attack= 4000 => 0.0
	*/
}

// func starAdsrtTimer(t *Tone, adsr *Adsr, gain float64, adsrType AdsrType) {
// 	releaseT := RescaleToMilliSeconds(*adsr.ReleaseTime, 0, 100, 10)
// 	ticker := time.NewTicker(time.Duration(1) * time.Millisecond)

// 	for {
// 		select {
// 		case <-ticker.C:
// 			t.adsrAction(adsrType, DecreaseAction, getRateValue(releaseT, gain), getSustainAmpValue(gain, *adsr.SustainAmp))
// 		case <-t.StopTime:
// 			ticker.Stop()
// 			return
// 		}
// 	}

// }

//
// func ticker() {
// 	ticker := time.NewTicker(1 * time.Millisecond)
// sus 100, max 100=> 100, sus 25, max 100=> 25. sus 40, max 100=> 40, sus 25, max 40=> 10 => 40*(sus/100)
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
