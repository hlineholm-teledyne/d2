package shape

import (
	"math"

	"oss.terrastruct.com/d2/lib/geo"
	"oss.terrastruct.com/d2/lib/svg"
)

type shapeCollate struct {
	*baseShape
}

func NewCollate(box *geo.Box) Shape {
	return &shapeCollate{
		baseShape: &baseShape{
			Type: COLLATE_TYPE,
			Box:  box,
		},
	}
}

func (s shapeCollate) GetInnerBox() *geo.Box {
	tl := s.Box.TopLeft.Copy()
	width := s.Box.Width
	height := s.Box.Height

	// Hourglass/X shape - reduce all sides
	reduction := 0.25
	tl.X += width * reduction
	tl.Y += height * reduction
	width *= (1 - 2*reduction)
	height *= (1 - 2*reduction)

	return geo.NewBox(tl, width, height)
}

func collatePath(box *geo.Box) *svg.SvgPathContext {
	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)

	// Hourglass: two triangles forming X
	// Draw hourglass as single path
	pc.StartAt(pc.Absolute(0, 0))
	pc.L(false, box.Width, 0)
	pc.L(false, 0, box.Height)
	pc.L(false, box.Width, box.Height)
	pc.Z()
	return pc
}

func (s shapeCollate) Perimeter() []geo.Intersectable {
	topLeft := s.Box.TopLeft
	topRight := geo.NewPoint(s.Box.TopLeft.X+s.Box.Width, s.Box.TopLeft.Y)
	bottomLeft := geo.NewPoint(s.Box.TopLeft.X, s.Box.TopLeft.Y+s.Box.Height)
	bottomRight := geo.NewPoint(s.Box.TopLeft.X+s.Box.Width, s.Box.TopLeft.Y+s.Box.Height)

	return []geo.Intersectable{
		geo.Segment{Start: topLeft, End: topRight},
		geo.Segment{Start: topRight, End: bottomLeft},
		geo.Segment{Start: bottomLeft, End: bottomRight},
		geo.Segment{Start: bottomRight, End: topLeft},
	}
}

func (s shapeCollate) GetSVGPathData() []string {
	return []string{
		collatePath(s.Box).PathData(),
	}
}

func (s shapeCollate) GetDimensionsToFit(width, height, paddingX, paddingY float64) (float64, float64) {
	return math.Abs(width) + paddingX*2, math.Abs(height) + paddingY*2
}

func (s shapeCollate) GetDefaultPadding() (paddingX, paddingY float64) {
	return defaultPadding, defaultPadding
}
