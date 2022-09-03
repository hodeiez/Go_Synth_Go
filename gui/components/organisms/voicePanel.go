package organism

import (
	components "hodei/gosynthgo/gui/components/molecules"
	"hodei/gosynthgo/gui/styles"
	"image"
	"image/color"

	g "github.com/AllenDang/giu"
)

type AdsrValues struct {
	Att int32
	Dec int32
	Sus int32
	Rel int32
}
type SelectorValues struct {
	TextValues    []string
	SelectedIndex int
}
type OscPanelValues struct {
	Adsr      *AdsrValues
	Vol       *float64
	Pitch     *float64
	PitchInit *float64
	Cut       *float64
	Res       *float64
	Pwm       *float64
	LfoR      *float64
	LfoW      *float64
	Noize     *float64
	Selector  *SelectorValues
}

type SynthValues struct {
	Osc1 *OscPanelValues
	Osc2 *OscPanelValues
}

//TODO: abstract and make dynamic
func VoicePanel(val *SynthValues) *g.ColumnWidget {

	return g.Column(panelStyled(oscPanel("OSC1", 0, val.Osc1), styles.Silver), panelStyled(oscPanel("OSC2", 220, val.Osc2), styles.Goldish))
}
func oscPanel(title string, ypos int, val *OscPanelValues) *g.ChildWidget {
	return oscPanelStyled().Layout(
		titleStyled(title),
		g.Row(g.Column(
			g.Dummy(400, 70),
			components.Selector(val.Selector.TextValues, &val.Selector.SelectedIndex).Build()), g.Dummy(50, 10),
			g.Column(g.Column(components.Button(styles.Red, styles.Blackish, styles.Redhite, val.Pitch, *val.PitchInit, "Tune me").Build())),
			// g.Column(components.Button(styles.Red, styles.Blackish, styles.Redhite, val.Pitch, *val.PitchInit, "Reset").Build())),//TODO: implement reset everything
			g.AlignManually(g.AlignRight, ADSRpanelInit(&val.Adsr.Att, &val.Adsr.Dec, &val.Adsr.Sus, &val.Adsr.Rel), 400, true)),

		components.Knob(image.Pt(g.GetCursorPos().X+200, g.GetCursorPos().Y+ypos), val.Vol, "VOL"),
		components.Knob(image.Pt(g.GetCursorPos().X+270, g.GetCursorPos().Y+ypos), val.Pitch, "PITCH").SetMinMax(220, 660),
		components.Knob(image.Pt(g.GetCursorPos().X+340, g.GetCursorPos().Y+ypos), val.Cut, "CUT").SetMinMax(20, 6000),
		components.Knob(image.Pt(g.GetCursorPos().X+410, g.GetCursorPos().Y+ypos), val.Res, "RES").SetMinMax(0, 1000),
		components.Knob(image.Pt(g.GetCursorPos().X+200, g.GetCursorPos().Y+100+ypos), val.Pwm, "PWM"),

		components.Knob(image.Pt(g.GetCursorPos().X+270, g.GetCursorPos().Y+100+ypos), val.LfoR, "LFO-R"),
		components.Knob(image.Pt(g.GetCursorPos().X+340, g.GetCursorPos().Y+100+ypos), val.LfoW, "LFO-W"),
		components.Knob(image.Pt(g.GetCursorPos().X+410, g.GetCursorPos().Y+100+ypos), val.Noize, "NOIZE"),
	)
}

func oscPanelStyled() *g.ChildWidget {
	return g.Child().Size(1000, 220).Flags(g.WindowFlags(g.StyleVarChildRounding))
}
func titleStyled(title string) *g.StyleSetter {
	return g.Style().SetColor(g.StyleColorText, color.RGBA{0, 0, 0, 255}).To(
		g.Label(title))
}
func panelStyled(children *g.ChildWidget, color color.RGBA) *g.StyleSetter {
	return g.Style().SetColor(g.StyleColorChildBg, color).To(children)
}
