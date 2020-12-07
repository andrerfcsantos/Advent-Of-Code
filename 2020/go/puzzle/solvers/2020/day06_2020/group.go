package day06_2020

type Group struct {
	Answers map[rune]int
	Size    int
}

func NewGroup() *Group {
	return &Group{
		Answers: make(map[rune]int),
	}
}

func (g *Group) AddPersonAnswers(answers string) {
	for _, a := range answers {
		g.Answers[a]++
	}
	g.Size++
}

func (g *Group) NumberOfAnswers() int {
	return len(g.Answers)
}

func (g *Group) NumberOfAnswersFromEveryone() int {
	count := 0
	for _, v := range g.Answers {
		if v == g.Size {
			count++
		}
	}
	return count
}
