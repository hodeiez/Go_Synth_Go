package midi

import (
	"log"
	"strconv"
	"strings"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"

	driver "gitlab.com/gomidi/rtmididrv"
)

type MidiMsg struct {
	Key int  //
	On  bool //
	// Off    bool //TODO:REMOVE THIS
	Vel    int64
	Cha    int64
	Ctl    int64
	CtlVal int64
	Pitch  int64
}

func RunMidi(val chan MidiMsg) {
	defer func() {
		if error := recover(); error != nil {
			log.Println("NO MIDI!")
		}
	}()

	drv, err := driver.New()

	must(err)

	// make sure to close all open ports at the end
	defer drv.Close()

	ins, err := drv.Ins()
	must(err)

	outs, err := drv.Outs()
	must(err)

	in, out := ins[0], outs[0]

	must(in.Open())
	must(out.Open())

	defer in.Close()
	defer out.Close()

	rd := reader.New(
		reader.NoLogger(),
		// format every message
		reader.Each(func(pos *reader.Position, msg midi.Message) {

			val <- ToMidiMsg(msg.String())

		}),
	)

	r := rd.ListenTo(in)
	log.Print("midi started listening")

	for {

		must(r)

	}

}

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ToMidiMsg(message string) MidiMsg {
	var isOn bool
	var velocity, theKey, ctl, ctlVal, pitch int64 = 0, 0, 0, 0, 0

	println(message)
	channel, _ := strconv.ParseInt(strings.Fields(message)[2], 10, 64)

	switch true {
	case strings.Contains(strings.Fields(message)[0], "NoteOff"):
		isOn = false
		theKey, _ = strconv.ParseInt(strings.Fields(message)[4], 10, 64)
	case strings.Contains(strings.Fields(message)[0], "NoteOn"):
		isOn = true
		velocity, _ = strconv.ParseInt(strings.Fields(message)[6], 10, 64)
		theKey, _ = strconv.ParseInt(strings.Fields(message)[4], 10, 64)
	case strings.Contains(strings.Fields(message)[0], "ControlChange"):
		ctl, _ = strconv.ParseInt(strings.Fields(message)[4], 10, 64)
		ctlVal, _ = strconv.ParseInt(strings.Fields(message)[len(strings.Fields(message))-1], 10, 64)
	case strings.Contains(strings.Fields(message)[0], "Pitchbend"):
		pitch, _ = strconv.ParseInt(strings.Fields(message)[4], 10, 64)
	}

	return MidiMsg{Key: int(theKey), On: isOn, Vel: velocity, Cha: channel, Ctl: ctl, CtlVal: ctlVal, Pitch: pitch}
}
