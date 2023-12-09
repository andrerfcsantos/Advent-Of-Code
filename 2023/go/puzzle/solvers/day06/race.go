package day06

type Race struct {
	Time           int
	RecordDistance int
}

func (r *Race) NumberOfWaysToWin() int {

	res := 0

	for timePressingButton := 1; timePressingButton < r.Time; timePressingButton++ {
		timeRacing := r.Time - timePressingButton
		speed := timePressingButton
		distance := speed * timeRacing
		if distance > r.RecordDistance {
			res++
		}

	}

	return res
}
