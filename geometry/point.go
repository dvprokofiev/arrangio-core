package geometry

type Point struct {
	X, Y, Z int16
}

// Point64 stores only global coordinates
type Point64 struct {
	X, Y, Z int64
}