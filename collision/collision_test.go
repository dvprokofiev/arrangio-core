package collision

import (
	"arrangio-core/geometry"
	"testing"
)

func TestCheckCollision(t *testing.T) {
	// cube 2 x 2 x 2
	box := geometry.Box{W: 2, H: 2, D: 2}

	objA := &geometry.Footprint{
		Anchor: geometry.Point64{X: 0, Y: 0, Z: 0},
		Shape:  box,
	}

	objB := &geometry.Footprint{
		Anchor: geometry.Point64{X: 1, Y: 0, Z: 0}, // overlaps with A
		Shape:  box,
	}

	objC := &geometry.Footprint{
		Anchor: geometry.Point64{X: 10, Y: 10, Z: 10}, // too far
		Shape:  box,
	}

	t.Run("Should detect collision", func(t *testing.T) {
		if !CheckCollision(objA, objB) {
			t.Error("Collision between A and B was not detected")
		}
	})

	t.Run("Should not detect collision", func(t *testing.T) {
		if CheckCollision(objA, objC) {
			t.Error("Collision between A and C detected, but should not have")
		}
	})
}

func TestVoxelShapeCollision(t *testing.T) {
	// two L-shaped objects
	lShape := []geometry.Point{
		{X: 0, Y: 0, Z: 0},
		{X: 1, Y: 0, Z: 0},
		{X: 0, Y: 1, Z: 0},
	}

	shapeA := geometry.NewVoxelShape(lShape)
	shapeB := geometry.NewVoxelShape(lShape)

	objA := &geometry.Footprint{
		Anchor: geometry.Point64{X: 0, Y: 0, Z: 0},
		Shape:  shapeA,
	}

	t.Run("Collision with itself", func(t *testing.T) {
		targetB := &geometry.Footprint{
			Anchor: geometry.Point64{X: 0, Y: 0, Z: 0},
			Shape:  shapeB,
		}
		if !CheckCollision(objA, targetB) {
			t.Error("VoxelShape should collide with itself at same position")
		}
	})

	t.Run("Collision at offset", func(t *testing.T) {
		// move objects B to offset (1, 0, 0)
		// should collide at (0 ,0, 0)
		targetB := &geometry.Footprint{
			Anchor: geometry.Point64{X: 1, Y: 0, Z: 0},
			Shape:  shapeB,
		}
		if !CheckCollision(objA, targetB) {
			t.Error("VoxelShape should collide when shifted")
		}
	})

	t.Run("No collision", func(t *testing.T) {
		targetB := &geometry.Footprint{
			Anchor: geometry.Point64{X: 10, Y: 10, Z: 10},
			Shape:  shapeB,
		}
		if CheckCollision(objA, targetB) {
			t.Error("VoxelShape should not collide when far away")
		}
	})
}
