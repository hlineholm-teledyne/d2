package shape

import (
	"math"

	"oss.terrastruct.com/d2/lib/geo"
	"oss.terrastruct.com/d2/lib/svg"
	"oss.terrastruct.com/util-go/go2"
)

type shapeDelay struct {
	*baseShape
}

// Delay has a rounded right edge (half-stadium shape)
const delayRightPadding = 20.

func NewDelay(box *geo.Box) Shape {
	shape := shapeDelay{
		baseShape: &baseShape{
			Type: DELAY_TYPE,
			Box:  box,
		},
	}
	shape.FullShape = go2.Pointer(Shape(shape))
	return shape
}

func (s shapeDelay) GetInnerBox() *geo.Box {
	tl := s.Box.TopLeft.Copy()
	width := s.Box.Width - delayRightPadding
	return geo.NewBox(tl, width, s.Box.Height)
}

func delayPath(box *geo.Box) *svg.SvgPathContext {
	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)
	width := box.Width
	height := box.Height
	radius := height / 2

	// Start at top-left
	pc.StartAt(pc.Absolute(0, 0))
	// Top edge to start of curve
	pc.L(false, width-radius, 0)
	// Right rounded edge (semicircle)
	// Using cubic Bezier to approximate semicircle
	cp1x := width - radius + radius*0.552
	cp1y := 0.0
	cp2x := width
	cp2y := height/2 - radius*0.552
	pc.C(false, cp1x, cp1y, cp2x, cp2y, width, height/2)
	// Bottom half of curve
	cp3x := width
	cp3y := height/2 + radius*0.552
	cp4x := width - radius + radius*0.552
	cp4y := height
	pc.C(false, cp3x, cp3y, cp4x, cp4y, width-radius, height)
	// Bottom edge
	pc.L(false, 0, height)
	// Close path
	pc.Z()
	return pc
}

func (s shapeDelay) Perimeter() []geo.Intersectable {
	return delayPath(s.Box).Path
}

func (s shapeDelay) GetSVGPathData() []string {
	return []string{
		delayPath(s.Box).PathData(),
	}
}

func (s shapeDelay) GetDimensionsToFit(width, height, paddingX, paddingY float64) (float64, float64) {
	totalWidth := width + paddingX + delayRightPadding
	return math.Ceil(totalWidth), math.Ceil(height + paddingY)
}
