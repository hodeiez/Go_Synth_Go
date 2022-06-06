package components

import (
	"image"
	"image/color"

	g "github.com/AllenDang/giu"
)

func AdsrGraph(baseX int, baseY int, att int32, dec int32, sus int32, rel int32) *g.CustomWidget {
	return g.Custom(func() {
		canvas := g.GetCanvas()
		pos := g.GetCursorScreenPos()
		color := color.RGBA{255, 0, 0, 255}

		p1 := pos.Add(image.Pt(baseX, baseY))
		p2 := pos.Add(image.Pt(baseX+int(att/2), baseY-100))
		p3 := pos.Add(image.Pt(baseX+int(att/2)+int(dec/2), baseY-(int(sus))))
		p4 := pos.Add(image.Pt(baseX+int(att/2)+int(dec/2), baseY-(int(sus))))
		p5 := pos.Add(image.Pt(baseX+int(att/2)+int(dec/2)+25, baseY-(int(sus))))
		p6 := pos.Add(image.Pt(baseX+int(att/2)+int(dec/2)+int(rel/2)+25, baseY))
		canvas.AddQuadFilled(p1, p2, p3, p4, color)
		canvas.AddQuadFilled(p1, p4, p5, p6, color)

	})
}
