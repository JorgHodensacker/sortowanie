// ----- funkcja do sortowania typu merge -----

package pkg

func Merge(s []int) []int {
	set := conquer(divide(s))
	return set
}

func conquer(set [][]int) []int {

	// iteruj dla kazdego elementu zbioru głównego
	for len(set) > 1 {

		// utwórz lewy i prawy podzbiór ze zbioru głównego
		// usuń uzyte elementy ze zbioru głównego
		leftTemp := set[0]
		rightTemp := set[1]
		set = set[2:]

		// utwórz zbiór tymczasowy i numerację
		var tempSet []int
		i, j := 0, 0

		// porównaj kazdy element z lewego i prawego podzbioru
		// umieść odpowiednie elemtenty na końcu zbioru tymczasowego
		for i < len(leftTemp) && j < len(rightTemp) {
			if leftTemp[i] < rightTemp[j] {
				tempSet = append(tempSet, leftTemp[i])
				i++
			} else {
				tempSet = append(tempSet, rightTemp[j])
				j++
			}
		}

		// dodaj pozostałe elementy z lewego podzbioru na końcu zbioru tymczasowego
		for i < len(leftTemp) {
			tempSet = append(tempSet, leftTemp[i])
			i++
		}

		// dodaj pozostałe elementy z prawego podzbioru na końcu zbioru tymczasowego
		for j < len(rightTemp) {
			tempSet = append(tempSet, rightTemp[j])
			j++
		}

		// dodaj zbiór tymczasowy do głównego zbioru
		set = append(set, tempSet)
	}

	// usuń nadmiar slice'ów
	s := set[0]
	return s
}

func divide(s []int) [][]int {
	// stworz slice set dla elementow niepodzielonych
	var set [][]int
	set = append(set, s)

	// stworz slice result dla zbioru koncowego
	var result [][]int

	// dla kazdego elementu ze zbioru
	for len(set) > 0 {
		// wez ostatni element ze zbioru set, przypisz do wartosci tmp,
		// usun ze zbioru set
		tmp := set[len(set)-1]
		set = set[:len(set)-1]

		// sprawdz czy dlugosc elementu tmp jest wieksza od 1,
		// jesli tak to dodaj do zbioru koncowego
		// jesli nie podziel na dwie rowne czesci
		if len(tmp) <= 1 {
			result = append(result, tmp)
		} else {
			split := len(tmp) / 2
			left := tmp[split:]
			right := tmp[:split]
			set = append(set, left, right)
		}
	}
	return result
}
