package entity

import (
	"arrangio-core/geometry"
	"arrangio-core/tags"
)

type Entity struct {
	ID        uint64
	Tags      tags.Mask
	Props     []float64
	Footprint geometry.Footprint
}

func NewEntity(id uint64, tagMask tags.Mask, props []float64, footprint geometry.Footprint) *Entity {
	return &Entity{
		ID:        id,
		Tags:      tagMask,
		Props:     props,
		Footprint: footprint,
	}
}

func (e *Entity) HasTags(required tags.Mask) bool {
	return e.Tags.Has(required)
}
