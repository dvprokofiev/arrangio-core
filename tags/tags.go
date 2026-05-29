package tags

import "sort"

type Mask struct {
	// tags which are used more often, that others
	FastBits uint64

	// use this if there are more than 64 tags
	DynamicIDs []uint16
}

func NewMask() Mask {
	return Mask{
		FastBits:   0,
		DynamicIDs: make([]uint16, 0, 4),
	}
}

// check if mask contains tags or group of tags
func (m Mask) Has(other Mask) bool {
	if (m.FastBits & other.FastBits) != other.FastBits {
		return false
	}

	if len(other.DynamicIDs) == 0 {
		return true
	}

	// use binary search
	for _, tagID := range other.DynamicIDs {
		// find the idx of the element in m.DynamicIDs which is
		// greater or equals some tagID in other.DynamicIDs
		idx := sort.Search(len(m.DynamicIDs), func(i int) bool {
			return m.DynamicIDs[i] >= tagID
		})
		// if the element not found or there's something else
		// which is not equal to tagID
		if idx == len(m.DynamicIDs) || m.DynamicIDs[idx] != tagID {
			return false
		}
	}

	return true
}

// add tag to the Mask
func (m Mask) With(tagID int) Mask {
	if tagID < 64 {
		m.FastBits |= (1 << uint64(tagID))
		return m
	}

	dynID := uint16(tagID)

	// search index in which tagID will be inserted
	idx := sort.Search(len(m.DynamicIDs), func(i int) bool {
		return m.DynamicIDs[i] >= dynID
	})

	// if all elements in m.DynamicIDs are smaller than tagID
	// or tagID must be inserted in between existent elements
	if idx == len(m.DynamicIDs) || m.DynamicIDs[idx] != dynID {
		// create independent slice to ensure branching works
		oldDyn := m.DynamicIDs
		m.DynamicIDs = make([]uint16, len(oldDyn)+1)

		copy(m.DynamicIDs[:idx], oldDyn[:idx])
		copy(m.DynamicIDs[idx+1:], oldDyn[idx:])
		m.DynamicIDs[idx] = dynID
	}

	return m
}
