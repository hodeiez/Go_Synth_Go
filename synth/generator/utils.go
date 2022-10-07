package generator



type Pwm struct {
}

func RescaleMidiValues(value int64, outMin float64, outMax float64) float64 {
	return float64(value-0)*(outMax-outMin)/float64(127) + outMin
}
func RescaleToMilliSeconds(value int32, inMin float64, inMax float64, milisecondsAmount int32) float64 {
	return (float64(value) - inMin) * float64(milisecondsAmount*1000) / inMax
}
func RescaleThis(val float64) float64 {

	return (1-20)*(val-(-0.01))/(0.01-(-0.01)) + 1
}
