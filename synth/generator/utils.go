package generator

// "os"
// "os/signal"

// "github.com/go-audio/audio"
// "github.com/go-audio/generator"

type Pwm struct {
}

func RescaleMidiValues(value int64, outMin float32, outMax float32) float32 {
	return float32(value-0)*(outMax-outMin)/float32(127) + outMin
}
func RescaleToMilliSeconds(value int32, inMin float32, inMax float32, milisecondsAmount int32) float32 {
	return (float32(value) - inMin) * float32(milisecondsAmount*1000) / inMax
}
