package rules

import (
	"arrangio-core/entity"
	"arrangio-core/geometry"
)

// calculate attraction/distance scores to the closest object matching `Selector`
// scaled from 0.0 to 1.0

type ProximityRule struct {
	Target         Selector
	MaxDist        int64 // define rule's search scope and normalization factor
	RequireClosest bool
}

func (r *ProximityRule) Evaluate(subject *entity.Entity, ctx *RuleContext) float64 {
	anchor := subject.Footprint.Anchor

	searchMin := geometry.Point64{
		X: anchor.X - r.MaxDist,
		Y: anchor.Y - r.MaxDist,
		Z: anchor.Z - r.MaxDist,
	}
	searchMax := geometry.Point64{
		X: anchor.X + r.MaxDist,
		Y: anchor.Y + r.MaxDist,
		Z: anchor.Z + r.MaxDist,
	}

	// all entities from `searchMin` to `searchMax`
	neighbors := ctx.Grid.QueryBuf(searchMin, searchMax, ctx.Buffer)

	minDistSq := r.MaxDist * r.MaxDist
	found := false

	for _, neighbor := range neighbors {
		if subject.ID == neighbor.ID || !r.Target.Matches(neighbor) {
			continue
		}

		// calculate distance between `subject` and found object
		nAnchor := neighbor.Footprint.Anchor
		dx := anchor.X - nAnchor.X
		dy := anchor.Y - nAnchor.Y
		dz := anchor.Z - nAnchor.Z

		distSq := (dx * dx) + (dy * dy) + (dz * dz)
		// calculate distance between `subject` and the closest object with matching `Selector`
		if distSq < minDistSq {
			minDistSq = distSq
			found = true
		}
	}

	if !found {
		return 0.0
	}

	// return normalized score -- greater means closer
	return 1.0 - (float64(minDistSq) / float64(r.MaxDist*r.MaxDist))
}
