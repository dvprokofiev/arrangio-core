package tags

import (
	"testing"
)

func TestMask_FastBits(t *testing.T) {
	mask := NewMask()
	mask = mask.With(5)

	if mask.FastBits != (1 << 5) {
		t.Errorf("FastBits mismatched: expected %d, got %d", 1<<5, mask.FastBits)
	}
}

func TestMask_Has_FastBits(t *testing.T) {
	mask := NewMask().With(10)
	target := NewMask().With(10)

	if !mask.Has(target) {
		t.Error("Mask should have the bit that was set")
	}
}

func TestMask_DynamicTags_Insert(t *testing.T) {
	mask := NewMask().With(100).With(150)

	if len(mask.DynamicIDs) != 2 {
		t.Errorf("Expected 2 inserted dynamic tags, got %d", len(mask.DynamicIDs))
	}
}

func TestMask_DynamicTags_Sorting(t *testing.T) {
	mask := NewMask().With(150).With(100).With(125)
	expected := []uint16{100, 125, 150}

	for i, v := range mask.DynamicIDs {
		if v != expected[i] {
			t.Errorf("Sorting error at index %d: expected %d, got %d", i, expected[i], v)
		}
	}
}

func TestMask_Has_Negative(t *testing.T) {
	mask := NewMask().With(10)
	absent := NewMask().With(20)

	if mask.Has(absent) {
		t.Error("Mask should not have a bit that was never ever set")
	}
}
