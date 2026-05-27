package collision

import "arrangio-core/geometry"

func CheckCollision(a, b *geometry.Footprint) bool {
	// broad-phase: objects do not overlap if their bounds do not
	if !boundsOverlap(a, b) {
		return false
	}

	// narrow phase
	// iterate through first object's points and
	// try to find them in second object
	collision := false
	a.Shape.ForEachPoint(func(p geometry.Point) {
		if collision {
			return
		}

		wx := int64(a.Anchor.X) + int64(p.X)
		wy := int64(a.Anchor.Y) + int64(p.Y)
		wz := int64(a.Anchor.Z) + int64(p.Z)

		if b.ContainsPoint(wx, wy, wz) {
			collision = true
		}
	})

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
