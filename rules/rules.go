package rules

import (
	"arrangio-core/entity"
	"arrangio-core/grid"
)

// axis definitions for some rules
const (
	AxisX uint8 = iota
	AxisY
	AxisZ
)

type RuleContext struct {
	Grid *grid.Grid
	// pre-allocated buffer for `QueryBug` method
	Buffer []*entity.Entity
}

type Rule interface {
	Evaluate(subject *entity.Entity, ctx *RuleContext) float64
}
