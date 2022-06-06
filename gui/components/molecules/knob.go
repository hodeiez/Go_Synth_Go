package components

import (
	"image"
	"image/color"

	"math"
	"strconv"

	g "github.com/AllenDang/giu"
)

type KnobWidget struct {
	pos       image.Point
	pValue    *float64
	label     string
	min       float64
	max       float64
	size      float32
	colorOut  color.RGBA
	colorIn   color.RGBA
	lineColor color.RGBA
}

func Knob(pos image.Point, pValue *float64, label string) *KnobWidget {
	return &KnobWidget{
		pos:       pos,
		pValue:    pValue,
		label:     label,
		min:       0,
		max:       100,
		size:      30,
		colorOut:  color.RGBA{200, 200, 200, 255},
		colorIn:   color.RGBA{0, 0, 0, 255},
		lineColor: color.RGBA{255, 0, 0, 255},
	}
}

func (k *KnobWidget) SetColors(colorOut color.RGBA, colorIn color.RGBA, lineColor color.RGBA) *KnobWidget {
	k.colorOut = colorOut
	k.colorIn = colorIn
	k.lineColor = lineColor
	return k
}
func (k *KnobWidget) SetLabel(label string) *KnobWidget {
	k.label = label
	return k
}
func (k *KnobWidget) SetMinMax(min float64, max float64) *KnobWidget {
	k.min = min
	k.max = max
	return k
}
func (k *KnobWidget) SetSize(size float32) *KnobWidget {
	k.size = size
	return k
}
func (k *KnobWidget) GetValue() float64 {
	return *k.pValue
}

//TODO:center text, use tooltip instead window
func (k *KnobWidget) Build() {

	ANGLE_MIN := 3.141592 * 0.75
	ANGLE_MAX := 3.141592 * 2.25
	t := (*k.pValue - k.min) / (k.max - k.min)
	angle := ANGLE_MIN + (ANGLE_MAX-ANGLE_MIN)*t
	angle_cos := math.Cos(angle)
	angle_sin := math.Sin(angle)
	innerSpaceX, innerSpaceY := g.GetItemInnerSpacing()
	padX, padY := g.GetWindowPadding()

	canvas := g.GetCanvas()

	g.SetCursorScreenPos(k.pos)
	pos := g.GetCursorScreenPos()

	center := image.Pt(pos.X+int(k.size), pos.Y+int(k.size))
	radiusOut := k.size
	radiusIn := radiusOut * 0.4

	g.InvisibleButton().Size(radiusOut*2, radiusOut*2).Build()
	active := g.IsItemActive()
	hovered := g.IsItemHovered()

	if active && g.Context.IO().GetMouseDelta().X != 0 {
		step := (k.max - k.min) / 200.00
		*k.pValue += float64(g.Context.IO().GetMouseDelta().X) * step
		if *k.pValue < k.min {
			*k.pValue = k.min
		}
		if *k.pValue > k.max {
			*k.pValue = k.max
		}

	}

	canvas.AddCircleFilled(center, radiusOut, k.colorOut)
	canvas.AddCircleFilled(center, radiusIn, k.colorIn)
	canvas.AddCircle(center, radiusOut, color.RGBA{169, 169, 169, 255}, 0, 2)
	canvas.AddLine(image.Pt(int(float32(center.X)+float32(angle_cos)*(radiusIn)), int(float32(center.Y)+float32(angle_sin)*(radiusIn))), image.Pt(int(float32(center.X)+(float32(angle_cos)*(radiusOut-2))), int(float32(center.Y)+(float32(angle_sin)*(radiusOut-2)))), k.lineColor, 5)
	canvas.AddText(image.Pt(pos.X+int(innerSpaceX)+centered(k.label), int(float32(pos.Y)+(radiusOut*2)-innerSpaceY)), k.lineColor, k.label)
	//canvas.AddText(image.Pt(int(float32(k.pos.X)+innerSpaceX), int(float32(k.pos.Y)+(radiusOut*2)-innerSpaceY)), k.lineColor, k.label)
	if active || hovered {

		//wnd:=g.SetNextWindowPos(float32(pos.X)-padX, float32(pos.Y)-padY)
		//g.Tooltip(strconv.FormatFloat(*k.pValue, 'f', 2, 64))
		wnd := g.SingleWindow()
		wnd.Size(50, 20)
		// wndX, wndY := wnd.CurrentSize()
		// wnd.Pos(float32(k.pos.X)-wndX/3, float32(k.pos.Y)+wndY)

		wnd.Pos(float32(pos.X)-padX, float32(pos.Y)-padY)
		wnd.Layout(g.Label(strconv.FormatFloat(*k.pValue, 'f', 0, 64)))
		//wnd.Layout(g.Tooltip("tip").Layout(g.BulletText(strconv.FormatFloat(*k.pValue, 'f', 0, 64))))

	}

}
func centered(s string) int {
	if len(s) <= 3 {
		return len(s) * len(s)
	} else {
		return 0
	}

}
