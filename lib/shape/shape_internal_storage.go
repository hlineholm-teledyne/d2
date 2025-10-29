package shape

import (
	"math"

	"oss.terrastruct.com/d2/lib/geo"
	"oss.terrastruct.com/d2/lib/svg"
)

const (
	internalStorageTopLine  = 20.
	internalStorageLeftLine = 20.
)

type shapeInternalStorage struct {
	*baseShape
}

func NewInternalStorage(box *geo.Box) Shape {
	return &shapeInternalStorage{
		baseShape: &baseShape{
			Type: INTERNAL_STORAGE_TYPE,
			Box:  box,
		},
	}
}

func (s shapeInternalStorage) GetInnerBox() *geo.Box {
	tl := s.Box.TopLeft.Copy()
	width := s.Box.Width
	height := s.Box.Height

	// Account for the top and left lines
	tl.X += internalStorageLeftLine
	tl.Y += internalStorageTopLine
	width -= internalStorageLeftLine
	height -= internalStorageTopLine

	return geo.NewBox(tl, width, height)
}

func internalStoragePath(box *geo.Box) *svg.SvgPathContext {
	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)
	// Main rectangle
	pc.StartAt(pc.Absolute(0, 0))
	pc.L(false, box.Width, 0)
	pc.L(false, 0, box.Height)
	pc.L(false, -box.Width, 0)
	pc.Z()
	return pc
}

func internalStorageTopLinePath(box *geo.Box) *svg.SvgPathContext {
	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)
	// Horizontal line near top
	pc.StartAt(pc.Absolute(0, internalStorageTopLine))
	pc.L(false, box.Width, 0)
	return pc
}

func internalStorageLeftLinePath(box *geo.Box) *svg.SvgPathContext {
	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)
	// Vertical line on left
	pc.StartAt(pc.Absolute(internalStorageLeftLine, 0))
	pc.L(false, 0, box.Height)
	return pc
}

func (s shapeInternalStorage) Perimeter() []geo.Intersectable {
	return boxPath(s.Box).Path
}

func (s shapeInternalStorage) GetSVGPathData() []string {
	return []string{
		internalStoragePath(s.Box).PathData(),
		internalStorageTopLinePath(s.Box).PathData(),
		internalStorageLeftLinePath(s.Box).PathData(),
	}
}

func (s shapeInternalStorage) GetDimensionsToFit(width, height, paddingX, paddingY float64) (float64, float64) {
	return math.Abs(width) + paddingX + internalStorageLeftLine,
		math.Abs(height) + paddingY + internalStorageTopLine
}

func (s shapeInternalStorage) GetDefaultPadding() (paddingX, paddingY float64) {
	return defaultPadding, defaultPadding
}
