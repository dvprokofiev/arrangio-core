package geometry

import "sort"

type Shape interface {
	Contains(lx, ly, lz int16) bool
	Bounds() (min, max Point)
	ForEachPoint(fn func(p Point))
}

// simple shape - parallelepiped
type Box struct {
	W, H, D int16
}

func (b Box) Contains(lx, ly, lz int16) bool {
	return lx >= 0 && lx < b.W && ly >= 0 && ly < b.H && lz >= 0 && lz < b.D
}

func (b Box) Bounds() (min, max Point) {
	return Point{0, 0, 0}, Point{b.W, b.H, b.D}
}

func (b Box) ForEachPoint(fn func(p Point)) {
	for x := int16(0); x < b.W; x++ {
		for y := int16(0); y < b.H; y++ {
			for z := int16(0); z < b.D; z++ {
				fn(Point{x, y, z})
			}
		}
	}
}

// VoxelShape - specified by list of points
type VoxelShape struct {
	points    []Point
	minBounds Point
	maxBounds Point
}

func NewVoxelShape(points []Point) *VoxelShape {
	if len(points) == 0 {
		return &VoxelShape{}
	}

	// using sorted slice to store points
	sortedPoints := make([]Point, len(points))
	copy(sortedPoints, points)

	sort.Slice(sortedPoints, func(i, j int) bool {
		if sortedPoints[i].X != sortedPoints[j].X {
			return sortedPoints[i].X < sortedPoints[j].X
		}
		if sortedPoints[i].Y != sortedPoints[j].Y {
			return sortedPoints[i].Y < sortedPoints[j].Y
		}
		return sortedPoints[i].Z < sortedPoints[j].Z
	})

	// calculate bounds of the object
	min := sortedPoints[0]
	max := sortedPoints[0]
	for _, p := range sortedPoints {
		if p.X < min.X {
			min.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.Z < min.Z {
			min.Z = p.Z
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
		if p.Z > max.Z {
			max.Z = p.Z
		}
	}

	return &VoxelShape{
		points:    sortedPoints,
		minBounds: min,
		maxBounds: Point{max.X + 1, max.Y + 1, max.Z + 1},
	}
}

func (v *VoxelShape) Contains(lx, ly, lz int16) bool {
	n := len(v.points)
	if n == 0 {
		return false
	}

	idx := sort.Search(n, func(i int) bool {
		p := v.points[i]
		if p.X != lx {
			return p.X >= lx
		}
		if p.Y != ly {
			return p.Y >= ly
		}
		return p.Z >= lz
	})

	return idx < n && v.points[idx].X == lx && v.points[idx].Y == ly && v.points[idx].Z == lz
}

func (v *VoxelShape) Bounds() (min, max Point) {
	return v.minBounds, v.maxBounds
}

func (v *VoxelShape) ForEachPoint(fn func(p Point)) {
	for _, p := range v.points {
		fn(p)
	}
}
