package rules

import (
	"arrangio-core/entity"
	"arrangio-core/grid"
)

type ConstraintType int

const (
	HardConstraint ConstraintType = iota
	SoftConstraint
)

type Rule interface {
	Type() ConstraintType
	// importance of a rule -- from 0 to 1
	Weight() float64
	// function that returns basic mathematical value of rule violation
	Evaluate(g *grid.Grid, target *entity.Entity) float64
}
