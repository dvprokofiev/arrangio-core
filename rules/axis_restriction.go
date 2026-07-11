package rules

import "arrangio-core/entity"

// define the type of comparison
type Operator uint8

const (
	OpEq  Operator = iota // ==
	OpNot                 // !=
	OpLt                  // <
	OpLe                  // <=
	OpGt                  // >
	OpGe                  // >=
)

type AxisRestrictionRule struct {
	Axis int64    // target axis value ('X', 'Y', 'Z')
	Op   Operator // comparison operator
	Ref  int64    // reference value to compare against
}

func (r *AxisRestrictionRule) Evaluate(subject *entity.Entity, ctx *RuleContext) float64 {
	var val int64
	switch r.Axis {
	case AxisX:
		val = subject.Footprint.Anchor.X
	case AxisY:
		val = subject.Footprint.Anchor.Y
	case AxisZ:
		val = subject.Footprint.Anchor.Z
	default:
		return 0.0
	}

	var matched bool
	switch r.Op {
	case OpEq:
		matched = val == r.Ref
	case OpNot:
		matched = val != r.Ref
	case OpLt:
		matched = val < r.Ref
	case OpLe:
		matched = val <= r.Ref
	case OpGt:
		matched = val > r.Ref
	case OpGe:
		matched = val >= r.Ref
	}

	if matched {
		return 1.0
	}

	// return smooth gradient penalty [0.0; 1.0] if rule is not matched
	// provide directional signal to solver, guiding mutation
	var diff int64
	switch r.Op {
	case OpEq:
		if val > r.Ref {
			diff = val - r.Ref
		} else {
			diff = r.Ref - val
		}
	case OpNot:
		return 0.5
	case OpLt, OpLe:
		diff = val - r.Ref
	case OpGt, OpGe:
		diff = r.Ref - val
	}

	if diff <= 0 {
		return 0.0
	}

	// divide by (diff+1) to protect against division-by-zero errors
	// non-linear fractional penalty, even small `diff` is crucial:
	// diff=0	->	return 1 (object is at the boundary)
	// diff=1	->	return 0.5 (object is one cell away from boundary)
	// ...
	return 1.0 / float64(diff+1)
}
