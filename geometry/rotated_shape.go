package geometry

type Rotatable interface {
	SetRotation(matrixIdx uint8)
	GetRotation() uint8
}

// decorator that virtualize rotations
// instead of changing underlying `Shape`, it dynamically
// applies rotation matrices to the `Shape` points
type RotatedShape struct {
	underlying Shape
	matrixIdx  uint8
}

func NewRotatedShape(base Shape) *RotatedShape {
	return &RotatedShape{
		underlying: base,
		matrixIdx:  0,
	}
}

func (r *RotatedShape) SetRotation(matrixIdx uint8) {
	if matrixIdx >= 24 {
		matrixIdx = 0
	}
	r.matrixIdx = matrixIdx
}

func (r *RotatedShape) GetRotation() uint8 {
	return r.matrixIdx
}

// *trick*: instead of rotating shape, we rotate the incoming point backward
// using the inverse matrix and query unrotated base shape
func (r *RotatedShape) Contains(lx, ly, lz int16) bool {
	if r.matrixIdx == 0 {
		return r.underlying.Contains(lx, ly, lz)
	}

	matrix := rotationMatrices[r.matrixIdx]

	invX := matrix[0][0]*lx + matrix[1][0]*ly + matrix[2][0]*lz
	invY := matrix[0][1]*lx + matrix[1][1]*ly + matrix[2][1]*lz
	invZ := matrix[0][2]*lx + matrix[1][2]*ly + matrix[2][2]*lz

	return r.underlying.Contains(invX, invY, invZ)
}

// return `RotatedShape` bounding box
func (r *RotatedShape) Bounds() (min, max Point) {
	minBase, maxBase := r.underlying.Bounds()

	if r.matrixIdx == 0 {
		return minBase, maxBase
	}

	matrix := rotationMatrices[r.matrixIdx]
	return transformBounds(minBase, maxBase, matrix)
}

// iterate over all points of the shape, applying a forward rotation to each 
func (r *RotatedShape) ForEachPoint(fn func(p Point)) {
	if r.matrixIdx == 0 {
		r.underlying.ForEachPoint(fn)
		return
	}

	matrix := rotationMatrices[r.matrixIdx]
	r.underlying.ForEachPoint(func(p Point) {
		fn(applyForwardRotation(p, matrix))
	})
}
