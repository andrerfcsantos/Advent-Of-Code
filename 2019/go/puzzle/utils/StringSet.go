package utils

// StringSet represents a set of strings
type StringSet struct {
	Set map[string]bool
}

// NewStringSet returns a newly created string set
func NewStringSet() StringSet {
	return StringSet{Set: make(map[string]bool)}
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
func (s StringSet) Elements(element string) []string {

	var res []string
	for k := range s.Set {
		res = append(res, k)
	}
	return res
}
