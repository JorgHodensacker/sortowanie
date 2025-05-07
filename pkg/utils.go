// funkcje do obsługi sortowania

package pkg

import (
	"math/rand"
)

// tworzenie zbioru liczb o losowej liczbie encji i ich wartości (set)
// użycie funkcji randomValues z argumentem randomQuantity
// zwraca setR([]int) - gotowy zbiór o losowej wielkości
func RandomSet() []int {
	setR := RandomValues(RandomQuantity())
	return setR
}

// tworzenie losowej liczby, która określi wielkość zbioru
// zwraca number(int) - losowa wielkość
func RandomQuantity() int {
	for {
		number := rand.Intn(1000)
		if number > 15 {
			return number
		}
	}
}

// tworzenie zbioru o podanej wielkości z losowymi wartościami
// zwraca set([]int) - gotowy zbiór
func RandomValues(q int) []int {
	set := []int{}
	for i := 0; i < q; i++ {
		set = append(set, rand.Intn(1000))
	}
	return set
}
