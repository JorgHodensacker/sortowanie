// ----- funkcja sortowania bubble -----

package pkg

// sortowanie typu bubble zbioru wartości losowych (s []int)
// zwraca s([]int) - posortowany zbiór
func Bubble(s []int) []int {
	sorted := true
	for sorted {
		sorted = false
		for i := 1; i < len(s); i++ {
			if s[i-1] > s[i] {
				s[i-1], s[i] = s[i], s[i-1]
				sorted = true
			}
		}
	}
	return s
}
