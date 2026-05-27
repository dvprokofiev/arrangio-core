package geometry

type Footprint struct {
	Anchor Point64
	Shape  Shape
}

// broad check
func (f *Footprint) IsInsideBound(wx, wy, wz int64) bool {
	min, max := f.Shape.Bounds()

	return wx >= f.Anchor.X+int64(min.X) && wx < f.Anchor.X+int64(max.X) &&
		wy >= f.Anchor.Y+int64(min.Y) && wy < f.Anchor.Y+int64(max.Y) &&
		wz >= f.Anchor.Z+int64(min.Z) && wz < f.Anchor.Z+int64(max.Z)
}

func (f *Footprint) ContainsPoint(wx, wy, wz int64) bool {
	if !f.IsInsideBound(wx, wy, wz) {
		return false
	}

	// narrowed down check
	lx := int16(wx - f.Anchor.X)
	ly := int16(wy - f.Anchor.Y)
	lz := int16(wz - f.Anchor.Z)

	return f.Shape.Contains(lx, ly, lz)
}
