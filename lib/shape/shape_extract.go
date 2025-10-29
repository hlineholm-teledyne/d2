package shape

import (
	"math"

	"oss.terrastruct.com/d2/lib/geo"
	"oss.terrastruct.com/d2/lib/svg"
)

type shapeExtract struct {
	*baseShape
}

func NewExtract(box *geo.Box) Shape {
	return &shapeExtract{
		baseShape: &baseShape{
			Type: EXTRACT_TYPE,
			Box:  box,
		},
	}
}

func (s shapeExtract) GetInnerBox() *geo.Box {
	tl := s.Box.TopLeft.Copy()
	width := s.Box.Width
	height := s.Box.Height

	// Triangle pointing up - reduce bottom and sides
	reduction := 0.15
	tl.X += width * reduction
	tl.Y += height * reduction
	width *= (1 - 2*reduction)
	height *= (1 - 2*reduction)

	return geo.NewBox(tl, width, height)
}

func extractPath(box *geo.Box) *svg.SvgPathContext {
	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)

	// Triangle pointing up: base at bottom, point at top center
	pc.StartAt(pc.Absolute(box.Width/2, 0))
	pc.L(false, box.Width, box.Height)
	pc.L(false, 0, box.Height)
	pc.Z()
	return pc
}

func (s shapeExtract) Perimeter() []geo.Intersectable {
	topCenter := geo.NewPoint(s.Box.TopLeft.X+s.Box.Width/2, s.Box.TopLeft.Y)
	bottomLeft := geo.NewPoint(s.Box.TopLeft.X, s.Box.TopLeft.Y+s.Box.Height)
	bottomRight := geo.NewPoint(s.Box.TopLeft.X+s.Box.Width, s.Box.TopLeft.Y+s.Box.Height)

	return []geo.Intersectable{
		geo.Segment{Start: topCenter, End: bottomRight},
		geo.Segment{Start: bottomRight, End: bottomLeft},
		geo.Segment{Start: bottomLeft, End: topCenter},
	}
}

func (s shapeExtract) GetSVGPathData() []string {
	return []string{
		extractPath(s.Box).PathData(),
	}
}

func (s shapeExtract) GetDimensionsToFit(width, height, paddingX, paddingY float64) (float64, float64) {
	return math.Abs(width) + paddingX*2, math.Abs(height) + paddingY*2
}

func (s shapeExtract) GetDefaultPadding() (paddingX, paddingY float64) {
	return defaultPadding, defaultPadding
}
