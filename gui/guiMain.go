package gui

import (
	organism "hodei/gosynthgo/gui/components/organisms"

	g "github.com/AllenDang/giu"
)

var (
	sizeX = 1000
	sizeY = 540
)

var wnd *g.MasterWindow

func panelWindow() *g.WindowWidget {
	wnd := g.Window("SynthPanel").Flags(g.WindowFlagsNoBackground | g.WindowFlagsNoCollapse | g.WindowFlagsNoTitleBar | g.WindowFlagsNoMove | g.WindowFlagsAlwaysAutoResize)

	return wnd
}
func resizeMe() {
	scale := g.Context.GetPlatform().GetContentScale()

	if scale > 1 {
		wnd.SetSize(int(float32(sizeX)/scale)+sizeX-int(float32(sizeX)/scale), int(float32(sizeY)/scale)+sizeY-int(float32(sizeY)/scale))
	} else {
		wnd.SetSize(int(float32(sizeX)/scale), int(float32(sizeY)/scale))
	}
}

func loop() {

	panelWindow().Layout(

		organism.VoicePanel(&synthValues),
	)

	resizeMe()

}

var synthValues = organism.SynthValues{}

func RunGUI(val interface{}) {
	wnd = g.NewMasterWindow("GO SYNTH GO", sizeX, sizeY, g.MasterWindowFlagsNotResizable)
	synthValues = val.(organism.SynthValues)
	wnd.Run(loop)

	defer wnd.Close()
}
