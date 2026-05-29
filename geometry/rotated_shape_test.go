package geometry

import (
	"testing"
)

func TestRotatedShape_BoundsAndContains(t *testing.T) {
	testShapes := []Shape{
		&Box{W: 2, H: 2, D: 2},
		NewVoxelShape([]Point{{0, 0, 0}, {1, 1, 1}}),
	}

	for _, base := range testShapes {
		rotated := NewRotatedShape(base)

		for i := uint8(0); i < 24; i++ {
			rotated.SetRotation(i)

			// check if bounds are valid
			min, max := rotated.Bounds()
			if min.X >= max.X || min.Y >= max.Y || min.Z >= max.Z {
				t.Errorf("Shape %T collapsed or inverted bounds at rotation %d: min%v max%v",
					base, i, min, max)
			}

			// RotatedShape should contain all points that it gives through ForEachPoint iterator
			rotated.ForEachPoint(func(wp Point) {
				if !rotated.Contains(wp.X, wp.Y, wp.Z) {
					t.Errorf("Rotation %d: RotatedShape failed to contain its own point %v", i, wp)
				}
			})
		}
	}
}

func TestRotatedShape_Identity(t *testing.T) {
	base := &Box{W: 5, H: 5, D: 5}
	rotated := NewRotatedShape(base)

	// rotation 0 equals underlying object
	rotated.SetRotation(0)

	minB, maxB := base.Bounds()
	minR, maxR := rotated.Bounds()

	if minR != minB || maxR != maxB {
		t.Errorf("Rotation 0 failed identity check: expected %v-%v, got %v-%v", minB, maxB, minR, maxR)
	}
}

func TestRotatedShape_Guardrail(t *testing.T) {
	base := &Box{W: 1, H: 1, D: 1}
	rotated := NewRotatedShape(base)

	rotated.SetRotation(24) // no such rotation indexed 24
	if rotated.GetRotation() != 0 {
		t.Errorf("Expected index 24 to reset to 0, got %d", rotated.GetRotation())
	}
}

func TestRotatedShape_NegativeContains(t *testing.T) {
	base := &Box{W: 2, H: 2, D: 2}
	rotated := NewRotatedShape(base)

	for i := uint8(0); i < 24; i++ {
		rotated.SetRotation(i)

		min, max := rotated.Bounds()

		outsidePoints := []Point{ // this points are 100% outside Box
			{min.X - 1, min.Y, min.Z},
			{max.X + 1, min.Y, min.Z},
			{min.X, min.Y - 1, min.Z},
			{min.X, max.Y + 1, min.Z},
			{min.X, min.Y, min.Z - 1},
			{min.X, min.Y, max.Z + 1},
		}

		for _, p := range outsidePoints {
			if rotated.Contains(p.X, p.Y, p.Z) {
				t.Errorf("Rotation %d: RotatedShape reported true for genuine outside point %v (Bounds: %v - %v)",
					i, p, min, max)
			}
		}
	}
}

// rotation shouldn't change the volume of the shape
func TestRotatedShape_PointConservation(t *testing.T) {
	points := []Point{{0, 0, 0}, {1, 0, 0}, {1, 1, 1}}
	base := NewVoxelShape(points)
	rotated := NewRotatedShape(base)

	expectedCount := len(points)

	for i := uint8(0); i < 24; i++ {
		rotated.SetRotation(i)

		actualCount := 0
		uniquePoints := make(map[Point]bool)

		rotated.ForEachPoint(func(p Point) {
			actualCount++
			uniquePoints[p] = true
		})

		if actualCount != expectedCount {
			t.Errorf("Rotation %d changed point count: expected %d, got %d", i, expectedCount, actualCount)
		}

		if len(uniquePoints) != expectedCount {
			t.Errorf("Rotation %d generated duplicate points: unique points count is %d", i, len(uniquePoints))
		}
	}
}

// end-to-end rotation logic test
// prevents "symmetrically broken" math where both forward and inverse
// matrices are wrong in identical ways
func TestRotatedShape_SpecificRotationCheck(t *testing.T) {
	base := &Box{W: 1, H: 2, D: 3}
	rotated := NewRotatedShape(base)

	foundTargetRotation := false

	for i := uint8(0); i < 24; i++ {
		rotated.SetRotation(i)
		min, max := rotated.Bounds()

		dx := max.X - min.X
		dy := max.Y - min.Y
		dz := max.Z - min.Z

		if dx == 2 && dy == 1 && dz == 3 {
			foundTargetRotation = true
			matrix := rotationMatrices[i]

			srcPoint := Point{X: 0, Y: 1, Z: 2}

			dstPoint := applyForwardRotation(srcPoint, matrix)

			absX := dstPoint.X
			if absX < 0 {
				absX = -absX
			}
			absY := dstPoint.Y
			if absY < 0 {
				absY = -absY
			}
			absZ := dstPoint.Z
			if absZ < 0 {
				absZ = -absZ
			}

			if absX != 1 || absY != 0 || absZ != 2 {
				t.Errorf("Rotation %d: Matrix failed structural transformation check. Got magnitudes %v, expected {1, 0, 2}",
					i, []int16{absX, absY, absZ})
			}

			if !rotated.Contains(dstPoint.X, dstPoint.Y, dstPoint.Z) {
				t.Errorf("Rotation %d: failed to contain expected mapped point %v", i, dstPoint)
			}
			break
		}
	}

	if !foundTargetRotation {
		t.Error("Could not find a valid 90-degree Z-axis rotation matrix in precomputed pool")
	}
}
