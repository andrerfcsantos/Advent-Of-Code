package day02

type Game struct {
	Id   int
	Sets []GameSet
}

func (g *Game) Power() int {
	var gameSet GameSet

	for _, set := range g.Sets {
		if gameSet.Red < set.Red {
			gameSet.Red = set.Red
		}
		if gameSet.Blue < set.Blue {
			gameSet.Blue = set.Blue
		}
		if gameSet.Green < set.Green {
			gameSet.Green = set.Green
		}
	}
	return gameSet.Power()
}

func (g *Game) IsValid() bool {
	for _, gameSet := range g.Sets {
		if !gameSet.IsValid() {
			return false
		}
	}
	return true

}

type GameSet struct {
	Red   int
	Blue  int
	Green int
}

func (g *GameSet) IsValid() bool {
	return g.Red <= 12 && g.Blue <= 14 && g.Green <= 13
}

func (g *GameSet) Power() int {
	return g.Red * g.Blue * g.Green
}
