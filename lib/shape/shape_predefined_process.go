package shape

import (
	"math"

	"oss.terrastruct.com/d2/lib/geo"
	"oss.terrastruct.com/d2/lib/svg"
	"oss.terrastruct.com/util-go/go2"
)

type shapePredefinedProcess struct {
	*baseShape
}

// Predefined process has vertical lines on left and right sides
const predefinedProcessInset = 20.

func NewPredefinedProcess(box *geo.Box) Shape {
	shape := shapePredefinedProcess{
		baseShape: &baseShape{
			Type: PREDEFINED_PROCESS_TYPE,
			Box:  box,
		},
	}
	shape.FullShape = go2.Pointer(Shape(shape))
	return shape
}

func (s shapePredefinedProcess) GetInnerBox() *geo.Box {
	tl := s.Box.TopLeft.Copy()
	width := s.Box.Width

	// Inner box is between the two vertical lines
	tl.X += predefinedProcessInset
	width -= 2 * predefinedProcessInset

	return geo.NewBox(tl, width, s.Box.Height)
}

func (s shapePredefinedProcess) IsRectangular() bool {
	return true
}

func predefinedProcessPath(box *geo.Box) *svg.SvgPathContext {
	// Main rectangle
	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)
	pc.StartAt(pc.Absolute(0, 0))
	pc.L(false, box.Width, 0)
	pc.L(false, box.Width, box.Height)
	pc.L(false, 0, box.Height)
	pc.Z()
	return pc
}

func predefinedProcessLeftLine(box *geo.Box) *svg.SvgPathContext {
	// Left vertical line
	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)
	pc.StartAt(pc.Absolute(predefinedProcessInset, 0))
	pc.L(false, predefinedProcessInset, box.Height)
	return pc
}

func predefinedProcessRightLine(box *geo.Box) *svg.SvgPathContext {
	// Right vertical line
	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)
	pc.StartAt(pc.Absolute(box.Width-predefinedProcessInset, 0))
	pc.L(false, box.Width-predefinedProcessInset, box.Height)
	return pc
}

func (s shapePredefinedProcess) Perimeter() []geo.Intersectable {
	return boxPath(s.Box).Path
}

func (s shapePredefinedProcess) GetSVGPathData() []string {
	return []string{
		predefinedProcessPath(s.Box).PathData(),
		predefinedProcessLeftLine(s.Box).PathData(),
		predefinedProcessRightLine(s.Box).PathData(),
	}
}

func (s shapePredefinedProcess) GetDimensionsToFit(width, height, paddingX, paddingY float64) (float64, float64) {
	totalWidth := width + paddingX + 2*predefinedProcessInset
	return math.Ceil(totalWidth), math.Ceil(height + paddingY)
}
