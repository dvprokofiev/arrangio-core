package grid

import (
	"arrangio-core/entity"
	"arrangio-core/geometry"
	"testing"
)

type mockShape struct {
	minBounds geometry.Point
	maxBounds geometry.Point
}

func (m *mockShape) Bounds() (geometry.Point, geometry.Point) {
	return m.minBounds, m.maxBounds
}

func (m *mockShape) Contains(x, y, z int16) bool {
	return true
}

func (m *mockShape) ForEachPoint(f func(geometry.Point)) {
	// ...
}

func newTestEntity(id uint64, min, max geometry.Point) *entity.Entity {
	return &entity.Entity{
		ID: id,
		Footprint: geometry.Footprint{
			Anchor: geometry.Point64{X: 0, Y: 0, Z: 0},
			Shape:  &mockShape{minBounds: min, maxBounds: max},
		},
	}
}

func TestInsertAndQuery(t *testing.T) {
	g := NewGrid(3, 64, 64, 64, 100)
	e := newTestEntity(1, geometry.Point{X: 10, Y: 10, Z: 10}, geometry.Point{X: 20, Y: 20, Z: 20})
	g.Insert(e)

	buf := make([]*entity.Entity, 0, 10)
	res := g.QueryBuf(
		geometry.Point64{X: 0, Y: 0, Z: 0},
		geometry.Point64{X: 30, Y: 30, Z: 30},
		buf,
	)

	if len(res) != 1 || res[0].ID != 1 {
		t.Fatalf("expected entity 1, got %v", res)
	}
}

func TestRemoval(t *testing.T) {
	g := NewGrid(3, 64, 64, 64, 100)
	e := newTestEntity(1, geometry.Point{X: 10, Y: 10, Z: 10}, geometry.Point{X: 20, Y: 20, Z: 20})
	g.Insert(e)
	g.Remove(e)

	res := g.QueryBuf(
		geometry.Point64{X: 0, Y: 0, Z: 0},
		geometry.Point64{X: 30, Y: 30, Z: 30},
		make([]*entity.Entity, 0, 10),
	)
	if len(res) != 0 {
		t.Fatal("entity still found after removal")
	}
}

func TestDeduplication(t *testing.T) {
	g := NewGrid(3, 64, 64, 64, 100)
	e := newTestEntity(1, geometry.Point{X: 0, Y: 0, Z: 0}, geometry.Point{X: 30, Y: 30, Z: 30})
	g.Insert(e)

	res := g.QueryBuf(
		geometry.Point64{X: 0, Y: 0, Z: 0},
		geometry.Point64{X: 60, Y: 60, Z: 60},
		make([]*entity.Entity, 0, 10),
	)
	if len(res) != 1 {
		t.Fatalf("expected unique entity, got %d", len(res))
	}
}

func TestGridBoundaries(t *testing.T) {
	g := NewGrid(3, 10, 10, 10, 100)
	e := newTestEntity(1, geometry.Point{X: 70, Y: 70, Z: 70}, geometry.Point{X: 80, Y: 80, Z: 80})
	g.Insert(e)

	res := g.QueryBuf(
		geometry.Point64{X: 75, Y: 75, Z: 75},
		geometry.Point64{X: 76, Y: 76, Z: 76},
		make([]*entity.Entity, 0, 10),
	)
	if len(res) != 0 {
		t.Fatal("entity should not be inserted/found out of grid bounds")
	}
}

func TestEmptyQuery(t *testing.T) {
	g := NewGrid(3, 64, 64, 64, 100)
	res := g.QueryBuf(
		geometry.Point64{X: 0, Y: 0, Z: 0},
		geometry.Point64{X: 10, Y: 10, Z: 10},
		make([]*entity.Entity, 0, 10),
	)
	if len(res) != 0 {
		t.Fatal("query on empty grid should return 0")
	}
}

func TestRemoveNonExistentEntity(t *testing.T) {
	g := NewGrid(3, 100, 100, 100, 100)
	e := newTestEntity(999, geometry.Point{X: 10, Y: 10, Z: 10}, geometry.Point{X: 11, Y: 11, Z: 11})

	g.Remove(e)
}

func TestInsertRemoveSequence(t *testing.T) {
	g := NewGrid(3, 100, 100, 100, 100)
	e1 := newTestEntity(1, geometry.Point{X: 10, Y: 10, Z: 10}, geometry.Point{X: 12, Y: 12, Z: 12})
	e2 := newTestEntity(2, geometry.Point{X: 10, Y: 10, Z: 10}, geometry.Point{X: 12, Y: 12, Z: 12})

	g.Insert(e1)
	g.Insert(e2)
	g.Remove(e1)

	buf := make([]*entity.Entity, 0, 10)
	res := g.QueryBuf(
		geometry.Point64{X: 0, Y: 0, Z: 0},
		geometry.Point64{X: 20, Y: 20, Z: 20},
		buf,
	)

	if len(res) != 1 || res[0].ID != 2 {
		t.Errorf("expected entity ID 2, got %+v", res)
	}
}

func BenchmarkGridInsert(b *testing.B) {
	g := NewGrid(3, 100, 100, 100, 1000)
	e := newTestEntity(1, geometry.Point{X: 10, Y: 10, Z: 10}, geometry.Point{X: 12, Y: 12, Z: 12})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Insert(e)
		g.Remove(e)
	}
}

func BenchmarkGridQuery(b *testing.B) {
	g := NewGrid(3, 100, 100, 100, 1000)
	e := newTestEntity(1, geometry.Point{X: 10, Y: 10, Z: 10}, geometry.Point{X: 12, Y: 12, Z: 12})
	g.Insert(e)

	buf := make([]*entity.Entity, 0, 10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = buf[:0]
		g.QueryBuf(
			geometry.Point64{X: 0, Y: 0, Z: 0},
			geometry.Point64{X: 50, Y: 50, Z: 50},
			buf,
		)
	}
}

func BenchmarkGridInsert_Giant(b *testing.B) {
	g := NewGrid(3, 100, 100, 100, 50000)
	e := newTestEntity(1, geometry.Point{X: 10, Y: 10, Z: 10}, geometry.Point{X: 60, Y: 60, Z: 60})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Insert(e)
		g.Remove(e)
	}
}

func BenchmarkGridQuery_Dense(b *testing.B) {
	g := NewGrid(3, 100, 100, 100, 10000)

	for id := uint64(1); id <= 100; id++ {
		e := newTestEntity(id, geometry.Point{X: 10, Y: 10, Z: 10}, geometry.Point{X: 15, Y: 15, Z: 15})
		g.Insert(e)
	}

	buf := make([]*entity.Entity, 0, 200)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = buf[:0]
		g.QueryBuf(
			geometry.Point64{X: 0, Y: 0, Z: 0},
			geometry.Point64{X: 30, Y: 30, Z: 30},
			buf,
		)
	}
}
