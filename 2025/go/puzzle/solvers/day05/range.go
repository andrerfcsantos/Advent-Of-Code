package day05

type Range struct {
	Start int
	End   int
}

func (r *Range) IsInRange(value int) bool {
	return value >= r.Start && value <= r.End
}
