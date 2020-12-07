package day07_2020

type InnerBag struct {
	Color string
	Qtd   int
}

type BagTree struct {
	Adjacency    map[string][]InnerBag
	InversedTree map[string][]string
}

func NewBagTree() *BagTree {
	return &BagTree{
		Adjacency:    make(map[string][]InnerBag),
		InversedTree: make(map[string][]string),
	}
}

func (b *BagTree) AddBag(color string, innerBags ...InnerBag) {
	for _, innerBag := range innerBags {
		b.InversedTree[innerBag.Color] = append(b.InversedTree[innerBag.Color], color)
	}

	b.Adjacency[color] = append(b.Adjacency[color], innerBags...)
}

func (b *BagTree) BagsContainedBy(color string) []InnerBag {
	if _, ok := b.Adjacency[color]; ok {
		return b.Adjacency[color]
	}

	return nil
}

func (b *BagTree) BagContainsColor(bag string, targetColor string) bool {
	var innerBags []InnerBag
	var ok bool

	if innerBags, ok = b.Adjacency[bag]; !ok {
		return false
	}

	for _, bag := range innerBags {
		if bag.Color == targetColor {
			return true
		}

		contains := b.BagContainsColor(bag.Color, targetColor)
		if contains {
			return true
		}
	}

	return false
}

func (b *BagTree) BagsContaining(color string) []string {
	if _, ok := b.InversedTree[color]; ok {
		return b.InversedTree[color]
	}

	return nil
}

func (b *BagTree) TotalInnerBagsOf(color string) int {
	bags := b.BagsContainedBy(color)
	if len(bags) == 0 {
		return 0
	}

	res := 0
	for _, contained := range bags {
		containerSpace := b.TotalInnerBagsOf(contained.Color)
		bags := contained.Qtd + contained.Qtd*containerSpace
		res += bags
	}

	return res
}
