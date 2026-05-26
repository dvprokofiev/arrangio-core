package geometry

type Footprint struct {
	Anchor Point
	Shape  Shape
}

func (f *Footprint) ContainsPoint(wx, wy, wz int32) bool {
	// calculate (l)ocal coordinates relatively to Anchor
	lx := int16(wx - int32(f.Anchor.X))
	ly := int16(wy - int32(f.Anchor.Y))
	lz := int16(wz - int32(f.Anchor.Z))

	return f.Shape.Contains(lx, ly, lz)
}
