package day07_2020

type InnerBag struct {
	Color string
	Qtd   int
}

type BagTree struct {
	Adjacency    map[string][]InnerBag
}

func NewBagTree() *BagTree {
	return &BagTree{
		Adjacency:    make(map[string][]InnerBag),
	}
}

// AddBag adds a bag to the bag tree
func (b *BagTree) AddBag(color string, innerBags ...InnerBag) {
	b.Adjacency[color] = append(b.Adjacency[color], innerBags...)
}

// BagsInside returns the list of bags contained directly by the bag of the color specified
func (b *BagTree) BagsInside(color string) []InnerBag {
	if _, ok := b.Adjacency[color]; ok {
		return b.Adjacency[color]
	}

	return nil
}

// BagContainsColor tells if a bag contains a bag of the target color, either directly or indirectly
func (b *BagTree) BagContainsColor(bag string, targetColor string) bool {
	var innerBags []InnerBag
	var ok bool

	if innerBags, ok = b.Adjacency[bag]; !ok {
		return false
	}

	for _, innerBag := range innerBags {
		if innerBag.Color == targetColor {
			return true
		}

		if b.BagContainsColor(innerBag.Color, targetColor) {
			return true
		}
	}

	return false
}

// BagsContaining gets the list of colors of the bags that contain bags with the specified color
func (b *BagTree) BagsContaining(color string) []string {
	var res []string

	for bag, _ := range b.Adjacency {
		if b.BagContainsColor(bag, color) {
			res = append(res, bag)
		}
	}

	return res
}

// NumberOfBagsContaining gets the number of color of bags that contain bags with the specified color
func (b *BagTree) NumberOfBagsContaining(color string) int {
	return len(b.BagsContaining(color))
}

// TotalInnerBagsOf returns the total number of inner bags of a bag with the specified color.
// The total returned refers to the total number of bags contained directly or indirectly by the
// bag of the specified color.
func (b *BagTree) TotalInnerBagsOf(color string) int {
	res := 0

	for _, innerBag := range b.BagsInside(color) {
		nBags := innerBag.Qtd + innerBag.Qtd * b.TotalInnerBagsOf(innerBag.Color)
		res += nBags
	}

	return res
}
