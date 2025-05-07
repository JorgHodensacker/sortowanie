// ----- funkcja do sortowania insert -----

package pkg

// sortowanie typu bubble zbioru wartości losowych (s []int)
// zwraca s([]int) - posortowany zbiór
func Insert(s []int) []int {
	for i := 1; i < len(s); i++ {
		j := i - 1
		tmp := s[i]
		for j >= 0 && s[j] > tmp {
			s[j+1], s[j] = s[j], s[j+1]
			j--
		}
	}
	return s
}
