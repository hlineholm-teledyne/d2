package shape

import (
	"math"

	"oss.terrastruct.com/d2/lib/geo"
	"oss.terrastruct.com/d2/lib/svg"
	"oss.terrastruct.com/util-go/go2"
)

type shapeDisplay struct {
	*baseShape
}

const displayPointAngleTan = 0.3

func NewDisplay(box *geo.Box) Shape {
	shape := shapeDisplay{
		baseShape: &baseShape{
			Type: DISPLAY_TYPE,
			Box:  box,
		},
	}
	shape.FullShape = go2.Pointer(Shape(shape))
	return shape
}

func (s shapeDisplay) GetInnerBox() *geo.Box {
	pointWidth := s.Box.Height * displayPointAngleTan

	tl := s.Box.TopLeft.Copy()
	tl.X += pointWidth
	width := s.Box.Width - pointWidth
	height := s.Box.Height
	return geo.NewBox(tl, width, height)
}

func displayPath(box *geo.Box) *svg.SvgPathContext {
	pointWidth := box.Height * displayPointAngleTan

	w := box.Width
	h := box.Height
	midY := h / 2.0

	tipX, tipY := 0.0, midY
	topLeftX, topLeftY := pointWidth, 0.0
	topRightX, topRightY := w, 0.0
	bottomRightX, bottomRightY := w, h
	bottomLeftX, bottomLeftY := pointWidth, h

	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)

	pc.StartAt(pc.Absolute(tipX, tipY))
	pc.L(false, topLeftX, topLeftY)
	pc.L(false, topRightX, topRightY)
	pc.L(false, bottomRightX, bottomRightY)
	pc.L(false, bottomLeftX, bottomLeftY)
	pc.L(false, tipX, tipY)
	pc.Z()

	return pc
}

func (s shapeDisplay) Perimeter() []geo.Intersectable {
	return displayPath(s.Box).Path
}

func (s shapeDisplay) GetSVGPathData() []string {
	return []string{
		displayPath(s.Box).PathData(),
	}
}

func (s shapeDisplay) GetDimensionsToFit(width, height, paddingX, paddingY float64) (float64, float64) {
	baseHeight := height + paddingY
	pointWidth := baseHeight * displayPointAngleTan

	totalWidth := width + paddingX + pointWidth
	totalHeight := baseHeight
	return math.Ceil(totalWidth), math.Ceil(totalHeight)
}
