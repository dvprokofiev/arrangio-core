package rules

import (
	"arrangio-core/collision"
	"arrangio-core/entity"
)

type NoCollisionRule struct {
	Target Selector
}

// check if object overlaps with any other object nearby
func (r *NoCollisionRule) Evaluate(subject *entity.Entity, ctx *RuleContext) float64 {
	minBounds, maxBounds := subject.Footprint.WorldBounds()

	neighbors := ctx.Grid.QueryBuf(minBounds, maxBounds, ctx.Buffer)

	for _, neighbor := range neighbors {
		if subject.ID == neighbor.ID {
			continue
		}

		if !r.Target.Matches(neighbor) {
			continue
		}

		if collision.CheckCollision(&subject.Footprint, &neighbor.Footprint) {
			return 0.0
		}
	}

	return 1.0
}
