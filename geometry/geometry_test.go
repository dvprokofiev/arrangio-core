package geometry

import (
	"testing"
)

func TestBox(t *testing.T) {
	b := Box{W: 10, H: 10, D: 10}

	tests := []struct {
		lx, ly, lz int16
		want       bool
	}{
		{0, 0, 0, true},
		{9, 9, 9, true}, // [0, 10)
		{-1, 0, 0, false},
		{10, 0, 0, false},
	}

	for _, tt := range tests {
		if got := b.Contains(tt.lx, tt.ly, tt.lz); got != tt.want {
			t.Errorf("Box.Contains(%d, %d, %d) = %v; want %v", tt.lx, tt.ly, tt.lz, got, tt.want)
		}
	}
}

func TestVoxelShape(t *testing.T) {
	// create L-shaped object
	pts := []Point{{1, 1, 1}, {2, 2, 2}}
	shape := NewVoxelShape(pts)

	t.Run("Valid points", func(t *testing.T) {
		if !shape.Contains(1, 1, 1) || !shape.Contains(2, 2, 2) {
			t.Error("VoxelShape failed to find existing points")
		}
	})

	t.Run("Invalid points", func(t *testing.T) {
		if shape.Contains(0, 0, 0) {
			t.Error("VoxelShape found non-existent point")
		}
	})

	t.Run("Bounds check", func(t *testing.T) {
		min, max := shape.Bounds()
		if min != (Point{1, 1, 1}) || max != (Point{3, 3, 3}) {
			t.Errorf("Wrong bounds: got %v to %v", min, max)
		}
	})
}

func TestFootprint(t *testing.T) {
	fp := Footprint{
		Anchor: Point64{X: 1000, Y: 2000, Z: 3000},
		Shape:  Box{W: 5, H: 5, D: 5},
	}

	tests := []struct {
		name       string
		wx, wy, wz int64
		want       bool
	}{
		{"Center of object", 1002, 2002, 3002, true},
		{"Boundary (min)", 1000, 2000, 3000, true},
		{"Boundary (max-1)", 1004, 2004, 3004, true},
		{"Outside", 1005, 2000, 3000, false},
		{"Far away", 0, 0, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fp.ContainsPoint(tt.wx, tt.wy, tt.wz); got != tt.want {
				t.Errorf("%s: ContainsPoint(%d,%d,%d) = %v; want %v",
					tt.name, tt.wx, tt.wy, tt.wz, got, tt.want)
			}
		})
	}
}

// test for two-step "is point inside" check
func TestFootprint_OverflowProtection(t *testing.T) {
	fp := Footprint{
		Anchor: Point64{X: 1_000_000_000, Y: 1_000_000_000, Z: 1_000_000_000},
		Shape:  Box{W: 5, H: 5, D: 5},
	}

	if fp.ContainsPoint(0, 0, 0) {
		t.Error("Should not find point far away")
	}
}

func Test_ForEachPoint(t *testing.T) {
	t.Run("Iterate through Box", func(t *testing.T) {
		box := Box{W: 2, H: 2, D: 2}
		count := 0
		box.ForEachPoint(func(p Point) {
			count++
		})
		if count != 8 {
			t.Errorf("Box expected 8 points, got %d", count)
		}
	})

	t.Run("Iterate through VoxelShape", func(t *testing.T) {
		pts := []Point{{0, 0, 0}, {1, 1, 1}, {2, 2, 2}}
		shape := NewVoxelShape(pts)

		count := 0
		shape.ForEachPoint(func(p Point) {
			count++
		})

		if count != 3 {
			t.Errorf("VoxelShape expected 3 points, got %d", count)
		}
	})
}
