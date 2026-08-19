package collision

import (
	"arrangio-core/geometry"
	"testing"
)

func BenchmarkCheckCollision_Boxes(b *testing.B) {
	objA := &geometry.Footprint{
		Anchor: geometry.Point64{X: 0, Y: 0, Z: 0},
		Shape:  geometry.Box{W: 10, H: 10, D: 10},
	}
	objB := &geometry.Footprint{
		Anchor: geometry.Point64{X: 5, Y: 5, Z: 5},
		Shape:  geometry.Box{W: 10, H: 10, D: 10},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckCollision(objA, objB)
	}
}

func BenchmarkCheckCollision_VoxelShapes(b *testing.B) {
	ptsA := []geometry.Point{}
	for x := int16(0); x < 10; x++ {
		for y := int16(0); y < 10; y++ {
			for z := int16(0); z < 10; z++ {
				ptsA = append(ptsA, geometry.Point{X: x, Y: y, Z: z})
			}
		}
	}

	ptsB := []geometry.Point{}
	for x := int16(0); x < 10; x++ {
		for y := int16(0); y < 10; y++ {
			for z := int16(0); z < 10; z++ {
				ptsB = append(ptsB, geometry.Point{X: x, Y: y, Z: z})
			}
		}
	}

	objA := &geometry.Footprint{
		Anchor: geometry.Point64{X: 0, Y: 0, Z: 0},
		Shape:  geometry.NewVoxelShape(ptsA),
	}
	objB := &geometry.Footprint{
		Anchor: geometry.Point64{X: 9, Y: 9, Z: 9},
		Shape:  geometry.NewVoxelShape(ptsB),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckCollision(objA, objB)
	}
}
