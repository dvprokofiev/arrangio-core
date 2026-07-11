package rules

import (
	"arrangio-core/entity"
	"arrangio-core/geometry"
	"math"
)

// Alignment rule helps to arrange objects in rows, ranks, etc...

type AlignmentRule struct {
	Target Selector
	Axis   int64
	Radius int64 // define search radius for this rule
}

func (r *AlignmentRule) Evaluate(subject *entity.Entity, ctx *RuleContext) float64 {
	anchor := subject.Footprint.Anchor

	searchMin := geometry.Point64{
		X: anchor.X - r.Radius,
		Y: anchor.Y - r.Radius,
		Z: anchor.Z - r.Radius,
	}

	searchMax := geometry.Point64{
		X: anchor.X + r.Radius,
		Y: anchor.Y + r.Radius,
		Z: anchor.Z + r.Radius,
	}

	neighbors := ctx.Grid.QueryBuf(searchMin, searchMax, ctx.Buffer)

	var minDiff int64 = math.MaxInt64
	var found bool

	var sVal int64
	switch r.Axis {
	case AxisX:
		sVal = anchor.X
	case AxisY:
		sVal = anchor.Y
	case AxisZ:
		sVal = anchor.Z
	default:
		return 0.0
	}

	for _, neighbor := range neighbors {
		// align only with objects within the `Radius` range and matching Selector
		if subject.ID == neighbor.ID || !r.Target.Matches(neighbor) {
			continue
		}

		var nVal int64
		switch r.Axis {
		case AxisX:
			nVal = neighbor.Footprint.Anchor.X
		case AxisY:
			nVal = neighbor.Footprint.Anchor.Y
		case AxisZ:
			nVal = neighbor.Footprint.Anchor.Z
		}

		diff := sVal - nVal
		if diff < 0 {
			diff = -diff
		}

		if diff < minDiff {
			minDiff = diff
			found = true
		}
	}

	if !found {
		return 1.0
	}

	if minDiff == 0 {
		return 1.0
	}

	return 1.0 / float64(minDiff+1)
}
