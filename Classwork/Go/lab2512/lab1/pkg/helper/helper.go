package helpers

func GenerateSortedAndShuffledArrays(n int) ([]int, []int) {
	sorted := make([]int, n)
	for i := 0; i < n; i++ {
		sorted[i] = i + 1
	}
	shuffled := make([]int, n)
	copy(shuffled, sorted)

	for i := n - 1; i > 0; i-- {
		j := randInt(0, i+1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return sorted, shuffled
}

func randInt(min, max int) int {
	return min + pseudoRand(max-min)
}

var seed int = 1

func pseudoRand(n int) int {
	seed = (seed*9301 + 49297) % 233280
	return (seed % n)
}
