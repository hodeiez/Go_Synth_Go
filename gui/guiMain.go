package gui

import (
	organism "hodei/gosynthgo/gui/components/organisms"

	g "github.com/AllenDang/giu"
)

var (
	sizeX = 1000
	sizeY = 650
)

func panelWindow() *g.WindowWidget {
	wnd := g.Window("SynthPanel").Flags(g.WindowFlagsNoBackground | g.WindowFlagsNoCollapse | g.WindowFlagsNoTitleBar | g.WindowFlagsNoResize | g.WindowFlagsNoMove)
	wnd.Size(float32(sizeX), float32(sizeY))
	return wnd
}

func loop() {

	panelWindow().Layout(

		organism.VoicePanel(&synthValues),
	)

}

var synthValues = organism.SynthValues{}

func RunGUI(val interface{}) {
	wnd := g.NewMasterWindow("GO SYNTH GO", sizeX-199, sizeY-225, g.MasterWindowFlagsTransparent)
	synthValues = val.(organism.SynthValues)
	wnd.Run(loop)
	defer wnd.Close()
}
