package shape

import (
	"math"

	"oss.terrastruct.com/d2/lib/geo"
	"oss.terrastruct.com/d2/lib/svg"
	"oss.terrastruct.com/util-go/go2"
)

type shapeManualOperation struct {
	*baseShape
}

const manualOperationBottomReduction = 0.15 // Bottom is 15% narrower on each side

func NewManualOperation(box *geo.Box) Shape {
	shape := shapeManualOperation{
		baseShape: &baseShape{
			Type: MANUAL_OPERATION_TYPE,
			Box:  box,
		},
	}
	shape.FullShape = go2.Pointer(Shape(shape))
	return shape
}

func (s shapeManualOperation) GetInnerBox() *geo.Box {
	// The inner box should fit within the narrower bottom part
	tl := s.Box.TopLeft.Copy()
	width := s.Box.Width
	height := s.Box.Height

	// Adjust for the slanted sides
	bottomInset := width * manualOperationBottomReduction
	tl.X += bottomInset
	width -= 2 * bottomInset

	// Add small vertical padding
	tl.Y += height * 0.1
	height *= 0.8

	return geo.NewBox(tl, width, height)
}

func manualOperationPath(box *geo.Box) *svg.SvgPathContext {
	width := box.Width
	height := box.Height
	bottomInset := width * manualOperationBottomReduction

	pc := svg.NewSVGPathContext(box.TopLeft, 1, 1)
	// Start at top-left
	pc.StartAt(pc.Absolute(0, 0))
	// Top edge to top-right
	pc.L(false, width, 0)
	// Right slanted edge to bottom-right (inset)
	pc.L(false, width-bottomInset, height)
	// Bottom edge to bottom-left (inset)
	pc.L(false, bottomInset, height)
	// Left slanted edge back to top-left
	pc.Z()
	return pc
}

func (s shapeManualOperation) Perimeter() []geo.Intersectable {
	return manualOperationPath(s.Box).Path
}

func (s shapeManualOperation) GetSVGPathData() []string {
	return []string{
		manualOperationPath(s.Box).PathData(),
	}
}

func (s shapeManualOperation) GetDimensionsToFit(width, height, paddingX, paddingY float64) (float64, float64) {
	// Account for the inset at the bottom
	bottomInset := (width + paddingX) * manualOperationBottomReduction
	totalWidth := width + paddingX + 2*bottomInset
	totalHeight := (height + paddingY) / 0.8 // Account for vertical scaling
	return math.Ceil(totalWidth), math.Ceil(totalHeight)
}

func (s shapeManualOperation) GetDefaultPadding() (paddingX, paddingY float64) {
	return defaultPadding / 2, defaultPadding / 2
}
