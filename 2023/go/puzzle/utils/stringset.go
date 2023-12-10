package utils

// StringSet represents a set of strings
type StringSet struct {
	Set map[string]bool
}


func StringSetIntersection(sets ...StringSet) StringSet {
	if len(sets) == 0{
		return NewStringSet()
	}

	if len(sets) == 1 {
		return sets[0]
	}

	return sets[0].Intersect(sets[1:]...)
}


// NewStringSet returns a newly created string set
func NewStringSet(elems ...string) StringSet {
	res := StringSet{Set: make(map[string]bool)}
	for _, elem := range elems {
		res.Set[elem] = true
	}
	return res
}

// Add adds an element to the string set
func (s StringSet) Add(element string) {
	s.Set[element] = true
}

// Remove removes an element from the string set
func (s StringSet) Remove(element string) {
	if s.Has(element) {
		delete(s.Set, element)
	}
}

// Has checks if an element exists on the string set
func (s StringSet) Has(element string) bool {
	_, ok := s.Set[element]
	return ok
}

// Elements returns the elements of this set as a slice of strings
func (s StringSet) Elements() []string {

	var res []string
	for k := range s.Set {
		res = append(res, k)
	}
	return res
}

// Intersect returns the intersection of a set with other sets
func (s StringSet) Intersect(otherSets ...StringSet) StringSet {
	var res = NewStringSet(s.Elements()...)

	for _, otherSet := range otherSets {
		for k := range res.Set {
			if !otherSet.Has(k) {
				res.Remove(k)
			}
		}
	}

	return res
}

// Union merges a set with other sets, so the resulting values in the set
// wil be the values present in all sets
func (s StringSet) Union(otherSets ...StringSet) StringSet {
	var res = NewStringSet(s.Elements()...)

	for _, otherSet := range otherSets {
		for elem := range otherSet.Set {
			s.Add(elem)
		}
	}

	return res
}
