package geometry

type Point struct {
	X, Y, Z int32
}

// Footprint describes the shape of the object through relative offsets
type Footprint struct {
	Offsets []Point
}
