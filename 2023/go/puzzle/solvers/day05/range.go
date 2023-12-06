package day05

type Range struct {
	DestinationStart int64
	SourceStart      int64
	Length           int64
}

type RangeGroup struct {
	Ranges []Range
}

func NewRangeGroup() *RangeGroup {
	return &RangeGroup{}
}

func (rg *RangeGroup) GetDestinationForSource(source int64) int64 {
	for _, r := range rg.Ranges {
		if source >= r.SourceStart && source < r.SourceStart+r.Length {
			return r.DestinationStart + (source - r.SourceStart)
		}
	}

	return source
}

func (rg *RangeGroup) AddRange(r Range) {
	rg.Ranges = append(rg.Ranges, r)
}
