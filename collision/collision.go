package collision

import "arrangio-core/geometry"

func CheckCollision(a, b *geometry.Footprint) bool {
	// broad-phase: objects do not overlap if their bounds do not
	if !boundsOverlap(a, b) {
		return false
	}

	// narrow phase
	// box optimization -- if two objects are boxes, their
	// AABB overlap equals their overlap
	aIsBox := isPureBox(a.Shape)
	bIsBox := isPureBox(b.Shape)
	if aIsBox && bIsBox {
		return true
	}
	bMin, bMax := b.Shape.Bounds()
	bMinX := b.Anchor.X + int64(bMin.X)
	bMaxX := b.Anchor.X + int64(bMax.X)
	bMinY := b.Anchor.Y + int64(bMin.Y)
	bMaxY := b.Anchor.Y + int64(bMax.Y)
	bMinZ := b.Anchor.Z + int64(bMin.Z)
	bMaxZ := b.Anchor.Z + int64(bMax.Z)

	diffX := a.Anchor.X - b.Anchor.X
	diffY := a.Anchor.Y - b.Anchor.Y
	diffZ := a.Anchor.Z - b.Anchor.Z

	aAnchorX := a.Anchor.X
	aAnchorY := a.Anchor.Y
	aAnchorZ := a.Anchor.Z

	// iterate through first object's points and
	// try to find them in second object
	collision := false

	switch bShape := b.Shape.(type) {
	case geometry.Box:
		a.Shape.ForEachPoint(func(p geometry.Point) bool {
			wx := aAnchorX + int64(p.X)
			wy := aAnchorY + int64(p.Y)
			wz := aAnchorZ + int64(p.Z)

			// bounds check
			if wx >= bMinX && wx < bMaxX &&
				wy >= bMinY && wy < bMaxY &&
				wz >= bMinZ && wz < bMaxZ {

				lx := int16(diffX + int64(p.X))
				ly := int16(diffY + int64(p.Y))
				lz := int16(diffZ + int64(p.Z))

				if bShape.Contains(lx, ly, lz) {
					collision = true
					return false // early exit
				}
			}

			return true // continue iteration
		})
	case *geometry.VoxelShape:
		a.Shape.ForEachPoint(func(p geometry.Point) bool {
			wx := aAnchorX + int64(p.X)
			wy := aAnchorY + int64(p.Y)
			wz := aAnchorZ + int64(p.Z)

			// bounds check
			if wx >= bMinX && wx < bMaxX &&
				wy >= bMinY && wy < bMaxY &&
				wz >= bMinZ && wz < bMaxZ {

				lx := int16(diffX + int64(p.X))
				ly := int16(diffY + int64(p.Y))
				lz := int16(diffZ + int64(p.Z))

				if bShape.Contains(lx, ly, lz) {
					collision = true
					return false // early exit
				}
			}

			return true // continue iteration
		})
	default:
		a.Shape.ForEachPoint(func(p geometry.Point) bool {
			wx := aAnchorX + int64(p.X)
			wy := aAnchorY + int64(p.Y)
			wz := aAnchorZ + int64(p.Z)

			// bounds check
			if wx >= bMinX && wx < bMaxX &&
				wy >= bMinY && wy < bMaxY &&
				wz >= bMinZ && wz < bMaxZ {

				lx := int16(diffX + int64(p.X))
				ly := int16(diffY + int64(p.Y))
				lz := int16(diffZ + int64(p.Z))

				if b.Shape.Contains(lx, ly, lz) {
					collision = true
					return false // early exit
				}
			}

			return true // continue iteration
		})
	}

	return collision
}

func boundsOverlap(a, b *geometry.Footprint) bool {
	minA, maxA := a.Shape.Bounds()
	minB, maxB := b.Shape.Bounds()

	ax1, ay1, az1 := int64(a.Anchor.X)+int64(minA.X), int64(a.Anchor.Y)+int64(minA.Y), int64(a.Anchor.Z)+int64(minA.Z)
	ax2, ay2, az2 := int64(a.Anchor.X)+int64(maxA.X), int64(a.Anchor.Y)+int64(maxA.Y), int64(a.Anchor.Z)+int64(maxA.Z)

	bx1, by1, bz1 := int64(b.Anchor.X)+int64(minB.X), int64(b.Anchor.Y)+int64(minB.Y), int64(b.Anchor.Z)+int64(minB.Z)
	bx2, by2, bz2 := int64(b.Anchor.X)+int64(maxB.X), int64(b.Anchor.Y)+int64(maxB.Y), int64(b.Anchor.Z)+int64(maxB.Z)

	return ax1 < bx2 && ax2 > bx1 &&
		ay1 < by2 && ay2 > by1 &&
		az1 < bz2 && az2 > bz1
}

func isPureBox(s geometry.Shape) bool {
	if s == nil {
		return false
	}

	if rs, ok := s.(*geometry.RotatedShape); ok {
		return isPureBox(rs.Unwrap())
	}

	_, ok := s.(*geometry.Box)
	return ok
}
