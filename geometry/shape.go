package geometry

type Shape interface {
	Contains(lx, ly, lz int16) bool
	Bounds() (min, max Point)
}

// simple shape - parallelepiped
type Box struct {
	W, H, D int16
}

func (b Box) Contains(lx, ly, lz int16) bool {
	return lx >= 0 && lx < b.W && ly >= 0 && ly <= b.H && lz >= 0 && lz < b.D
}

func (b Box) Bounds() (min, max Point) {
	return Point{0, 0, 0}, Point{b.W, b.H, b.D}
}

// VoxelShape - specified by list of points
type VoxelShape struct {
	Points []Point
}

func (v *VoxelShape) Contains(lx, ly, lz int16) bool {
	for _, p := range v.Points {
		if p.X == lx && p.Y == ly && p.Z == lz {
			return true
		}
	}
	return false
}

func (v *VoxelShape) Bounds() (min, max Point) {
	if len(v.Points) == 0 {
		return Point{0, 0, 0}, Point{0, 0, 0}
	}
	min, max = v.Points[0], v.Points[0]
	for _, p := range v.Points {
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
	return min, max
}
