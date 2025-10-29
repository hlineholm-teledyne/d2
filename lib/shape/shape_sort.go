package shape

import (
	"math"

	"oss.terrastruct.com/d2/lib/geo"
	"oss.terrastruct.com/d2/lib/svg"
)

type shapeSort struct {
	*baseShape
}

func NewSort(box *geo.Box) Shape {
	return &shapeSort{
		baseShape: &baseShape{
			Type: SORT_TYPE,
			Box:  box,
		},
	}
}

func (s shapeSort) GetInnerBox() *geo.Box {
	tl := s.Box.TopLeft.Copy()
	width := s.Box.Width
	height := s.Box.Height

	// Diamond shape with line - same as regular diamond
	tl.X += width / 4.
	tl.Y += height / 4.
	width /= 2.
	height /= 2.

	return geo.NewBox(tl, width, height)
}

func sortDiamondPath(box *geo.Box) *svg.SvgPathContext {
	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)

	// Diamond shape
	pc.StartAt(pc.Absolute(box.Width/2, 0))
	pc.L(false, box.Width, box.Height/2)
	pc.L(false, box.Width/2, box.Height)
	pc.L(false, 0, box.Height/2)
	pc.Z()
	return pc
}

func sortLinePath(box *geo.Box) *svg.SvgPathContext {
	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)

	// Horizontal line through middle of diamond
	pc.StartAt(pc.Absolute(0, box.Height/2))
	pc.L(false, box.Width, 0)
	return pc
}

func (s shapeSort) Perimeter() []geo.Intersectable {
	topCenter := geo.NewPoint(s.Box.TopLeft.X+s.Box.Width/2, s.Box.TopLeft.Y)
	rightCenter := geo.NewPoint(s.Box.TopLeft.X+s.Box.Width, s.Box.TopLeft.Y+s.Box.Height/2)
	bottomCenter := geo.NewPoint(s.Box.TopLeft.X+s.Box.Width/2, s.Box.TopLeft.Y+s.Box.Height)
	leftCenter := geo.NewPoint(s.Box.TopLeft.X, s.Box.TopLeft.Y+s.Box.Height/2)

	return []geo.Intersectable{
		geo.Segment{Start: topCenter, End: rightCenter},
		geo.Segment{Start: rightCenter, End: bottomCenter},
		geo.Segment{Start: bottomCenter, End: leftCenter},
		geo.Segment{Start: leftCenter, End: topCenter},
	}
}

func (s shapeSort) GetSVGPathData() []string {
	return []string{
		sortDiamondPath(s.Box).PathData(),
		sortLinePath(s.Box).PathData(),
	}
}

func (s shapeSort) GetDimensionsToFit(width, height, paddingX, paddingY float64) (float64, float64) {
	return math.Abs(width) + paddingX*2, math.Abs(height) + paddingY*2
}

func (s shapeSort) GetDefaultPadding() (paddingX, paddingY float64) {
	return defaultPadding, defaultPadding
}
