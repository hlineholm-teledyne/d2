package shape

import (
	"math"

	"oss.terrastruct.com/d2/lib/geo"
	"oss.terrastruct.com/d2/lib/svg"
)

type shapeSummingJunction struct {
	*baseShape
}

func NewSummingJunction(box *geo.Box) Shape {
	return &shapeSummingJunction{
		baseShape: &baseShape{
			Type: SUMMING_JUNCTION_TYPE,
			Box:  box,
		},
	}
}

func (s shapeSummingJunction) GetInnerBox() *geo.Box {
	tl := s.Box.TopLeft.Copy()
	width := s.Box.Width
	height := s.Box.Height

	// Circle with X - reduce for text placement
	reduction := 0.25
	tl.X += width * reduction
	tl.Y += height * reduction
	width *= (1 - 2*reduction)
	height *= (1 - 2*reduction)

	return geo.NewBox(tl, width, height)
}

func summingJunctionCirclePath(box *geo.Box) *svg.SvgPathContext {
	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)

	// Use same circle logic as existing circle shape
	radiusX := box.Width / 2
	radiusY := box.Height / 2

	// Circle using four cubic Bezier curves
	// Control point distance for circle approximation
	c := 0.552

	pc.StartAt(pc.Absolute(box.Width/2, 0))

	// Top to right
	pc.C(false, c*radiusX, 0, radiusX, (1-c)*radiusY, radiusX, radiusY)
	// Right to bottom
	pc.C(false, 0, c*radiusY, -(1-c)*radiusX, radiusY, -radiusX, radiusY)
	// Bottom to left
	pc.C(false, -c*radiusX, 0, -radiusX, -(1-c)*radiusY, -radiusX, -radiusY)
	// Left to top
	pc.C(false, 0, -c*radiusY, (1-c)*radiusX, -radiusY, radiusX, -radiusY)

	pc.Z()
	return pc
}

func summingJunctionXPath(box *geo.Box) []string {
	pc1 := svg.NewSVGPathContext(box.TopLeft, 1, 1)
	pc2 := svg.NewSVGPathContext(box.TopLeft, 1, 1)

	// X lines from corner to corner
	// Diagonal from top-left to bottom-right
	pc1.StartAt(pc1.Absolute(0, 0))
	pc1.L(false, box.Width, box.Height)

	// Diagonal from top-right to bottom-left
	pc2.StartAt(pc2.Absolute(box.Width, 0))
	pc2.L(false, 0, box.Height)

	return []string{pc1.PathData(), pc2.PathData()}
}

func (s shapeSummingJunction) Perimeter() []geo.Intersectable {
	return []geo.Intersectable{geo.NewEllipse(s.Box.Center(), s.Box.Width/2, s.Box.Height/2)}
}

func (s shapeSummingJunction) GetSVGPathData() []string {
	paths := []string{summingJunctionCirclePath(s.Box).PathData()}
	paths = append(paths, summingJunctionXPath(s.Box)...)
	return paths
}

func (s shapeSummingJunction) GetDimensionsToFit(width, height, paddingX, paddingY float64) (float64, float64) {
	return math.Abs(width) + paddingX*2, math.Abs(height) + paddingY*2
}

func (s shapeSummingJunction) GetDefaultPadding() (paddingX, paddingY float64) {
	return defaultPadding, defaultPadding
}
