package utils

// IntPermutations generates all the permutations of integers using Heap's algorithm
func IntPermutations(arr []int) [][]int {

	var res [][]int

	var helper func([]int, int)

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}


// Dist gives all the combination of lists with the given size that have value as sum
func Dist(size int, value int) [][]int {
	if size == 1 {
		return [][]int{[]int{value}}
	}

	var s [][]int

	for i:= value; i >= 0 ; i-- {
		sub := Dist(size-1, value-i)
		for _, list := range sub {
			s = append(s, append([]int{i}, list...))
		}
	}

	return s

}