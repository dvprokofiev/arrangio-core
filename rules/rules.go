package rules

import (
	"arrangio-core/entity"
	"arrangio-core/grid"
)

type RuleContext struct {
	Grid *grid.Grid
	// pre-allocated buffer for `QueryBug` method
	Buffer []*entity.Entity
}

type Rule interface {
	Evaluate(subject *entity.Entity, ctx *RuleContext) float64
}
