package shape

import (
	"math"

	"oss.terrastruct.com/d2/lib/geo"
	"oss.terrastruct.com/d2/lib/svg"
)

type shapeMerge struct {
	*baseShape
}

func NewMerge(box *geo.Box) Shape {
	return &shapeMerge{
		baseShape: &baseShape{
			Type: MERGE_TYPE,
			Box:  box,
		},
	}
}

func (s shapeMerge) GetInnerBox() *geo.Box {
	tl := s.Box.TopLeft.Copy()
	width := s.Box.Width
	height := s.Box.Height

	// Triangle pointing down - reduce top and sides
	reduction := 0.15
	tl.X += width * reduction
	tl.Y += height * reduction
	width *= (1 - 2*reduction)
	height *= (1 - 2*reduction)

	return geo.NewBox(tl, width, height)
}

func mergePath(box *geo.Box) *svg.SvgPathContext {
	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)

	// Triangle pointing down: base at top, point at bottom center
	pc.StartAt(pc.Absolute(0, 0))
	pc.L(false, box.Width, 0)
	pc.L(false, box.Width/2, box.Height)
	pc.Z()
	return pc
}

func (s shapeMerge) Perimeter() []geo.Intersectable {
	topLeft := s.Box.TopLeft
	topRight := geo.NewPoint(s.Box.TopLeft.X+s.Box.Width, s.Box.TopLeft.Y)
	bottomCenter := geo.NewPoint(s.Box.TopLeft.X+s.Box.Width/2, s.Box.TopLeft.Y+s.Box.Height)

	return []geo.Intersectable{
		geo.Segment{Start: topLeft, End: topRight},
		geo.Segment{Start: topRight, End: bottomCenter},
		geo.Segment{Start: bottomCenter, End: topLeft},
	}
}

func (s shapeMerge) GetSVGPathData() []string {
	return []string{
		mergePath(s.Box).PathData(),
	}
}

func (s shapeMerge) GetDimensionsToFit(width, height, paddingX, paddingY float64) (float64, float64) {
	return math.Abs(width) + paddingX*2, math.Abs(height) + paddingY*2
}

func (s shapeMerge) GetDefaultPadding() (paddingX, paddingY float64) {
	return defaultPadding, defaultPadding
}
