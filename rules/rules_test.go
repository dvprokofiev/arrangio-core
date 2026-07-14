package rules_test

import (
	"math"
	"testing"

	"arrangio-core/entity"
	"arrangio-core/geometry"
	"arrangio-core/grid"
	"arrangio-core/rules"
	"arrangio-core/tags"
)

// TestEntity is a "blank" for `buildTestEntity` function
type TestEntity struct {
	ID      uint64
	Anchor  geometry.Point64
	Tags    []int // tags that would be assigned to this entity
	W, H, D int16 // 1x1x1 by default
}

// RuleTestCase describes the test case
type RuleTestCase struct {
	Name      string
	Subject   TestEntity
	Neighbors []TestEntity // slice of entities that would be placed to the grid
	Expected  float64
	Tolerance float64
}

// buildTestEntity builds an `Entity` from `TestEntity`
func buildTestEntity(cfg TestEntity) *entity.Entity {
	mask := tags.NewMask()
	for _, tagID := range cfg.Tags {
		mask = mask.With(tagID)
	}

	w, h, d := cfg.W, cfg.H, cfg.D
	if w == 0 {
		w = 1
	}
	if h == 0 {
		h = 1
	}
	if d == 0 {
		d = 1
	}

	return &entity.Entity{
		ID:   cfg.ID,
		Tags: mask,
		Footprint: geometry.Footprint{
			Anchor: cfg.Anchor,
			Shape:  geometry.Box{W: w, H: h, D: d},
		},
	}
}

func RunRuleTest(t *testing.T, rule rules.Rule, tc RuleTestCase) {
	t.Helper()
	t.Run(tc.Name, func(t *testing.T) {
		// initialize grid (shiftBits=3 -> cell size is 8)
		g := grid.NewGrid(3, 2000, 2000, 2000, 100)

		// placing neighbors
		for _, nCfg := range tc.Neighbors {
			neighbor := buildTestEntity(nCfg)
			g.Insert(neighbor)
		}

		// building subject from `TestEntity`
		subject := buildTestEntity(tc.Subject)
		g.Insert(subject)

		ctx := &rules.RuleContext{
			Grid:   g,
			Buffer: make([]*entity.Entity, 0, 50),
		}

		got := rule.Evaluate(subject, ctx)

		tol := tc.Tolerance
		if tol == 0 {
			tol = 1e-9
		}

		if math.Abs(got-tc.Expected) > tol {
			t.Errorf("\n[FAIL] %s\nEvaluate() = %f, want %f (delta: %e)", tc.Name, got, tc.Expected, tol)
		}
	})
}
