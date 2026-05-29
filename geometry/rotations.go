package geometry

// because we rotate only in 90--degree increments, the matrix cells
// only contain -1, 0, 1
type RotationMatrix [3][3]int16

// precomputed pool of all 24 valid 3D rotations
var rotationMatrices = computeRotations()

// rotates local point `p` forward into world space with matrix `m`
// standard matrix-vector multiplication
func applyForwardRotation(p Point, m RotationMatrix) Point {
	return Point{
		X: p.X*m[0][0] + p.Y*m[0][1] + p.Z*m[0][2],
		Y: p.X*m[1][0] + p.Y*m[1][1] + p.Z*m[1][2],
		Z: p.X*m[2][0] + p.Y*m[2][1] + p.Z*m[2][2],
	}
}

// calculates new bounding box for a rotated shape
func transformBounds(min, max Point, m RotationMatrix) (Point, Point) {
	// generate all 8 corner vertices of the original bounding box
	vertices := [8]Point{
		{min.X, min.Y, min.Z}, {max.X, min.Y, min.Z},
		{min.X, max.Y, min.Z}, {max.X, max.Y, min.Z},
		{min.X, min.Y, max.Z}, {max.X, min.Y, max.Z},
		{min.X, max.Y, max.Z}, {max.X, max.Y, max.Z},
	}

	resMin := Point{X: 32767, Y: 32767, Z: 32767}
	resMax := Point{X: -32768, Y: -32768, Z: -32768}

	for _, v := range vertices {
		rv := applyForwardRotation(v, m)
		if rv.X < resMin.X {
			resMin.X = rv.X
		}
		if rv.Y < resMin.Y {
			resMin.Y = rv.Y
		}
		if rv.Z < resMin.Z {
			resMin.Z = rv.Z
		}

		if rv.X > resMax.X {
			resMax.X = rv.X
		}
		if rv.Y > resMax.Y {
			resMax.Y = rv.Y
		}
		if rv.Z > resMax.Z {
			resMax.Z = rv.Z
		}
	}

	return resMin, resMax
}

func GetRotationMatrices() [24]RotationMatrix {
	return rotationMatrices
}

// precompute 24 rotations
func computeRotations() [24]RotationMatrix {
	var matrices [24]RotationMatrix
	idx := 0

	// 6 possible permutations of X, Y and Z
	permutations := [][]int{
		{0, 1, 2}, {0, 2, 1}, {1, 0, 2},
		{1, 2, 0}, {2, 0, 1}, {2, 1, 0},
	}

	// nested loops go through all permutations in two possible directions
	// going from 1 to -1 in order to `Contains()` method of decorator to be branchless
	for _, p := range permutations {
		for _, sx := range []int16{1, -1} {
			for _, sy := range []int16{1, -1} {
				for _, sz := range []int16{1, -1} {
					var m RotationMatrix
					m[0][p[0]] = sx
					m[1][p[1]] = sy
					m[2][p[2]] = sz

					// calculate matrix determinant
					det := m[0][0]*(m[1][1]*m[2][2]-m[1][2]*m[2][1]) -
						m[0][1]*(m[1][0]*m[2][2]-m[1][2]*m[2][0]) +
						m[0][2]*(m[1][0]*m[2][1]-m[1][1]*m[2][0])

					// det == 1 ensures a valid rotation
					if det == 1 && idx < 24 {
						matrices[idx] = m
						idx++
					}
				}
			}
		}
	}
	return matrices
}
