package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fdominas/sortowanie/pkg"
)

func main() {
	// error message i spis argumentow funkcji
	message := `poprawny argument z listy:
	1. bubble
	2. insert
	3. merge`

	// sprawdź czy podano argument
	if len(os.Args) < 2 {
		fmt.Println("Dodaj", message)
		return
	}

	// generowanie zbioru losowych wartości
	set := pkg.RandomSet()

	time_start := time.Now().UnixNano()

	// sprawdzenie czy podano odpowiedni argument przy uruchomieniu funkcji
	switch os.Args[1] {
	case "bubble": // podano bubble
		fmt.Println("Przed uporządkowniem: ", set)
		fmt.Println("Po uporządkowaniu: ", pkg.Bubble(set))
	case "insert":
		fmt.Println("Przed uporzadkowaniem: ", set)
		fmt.Println("Po uporzadkowaniu: ", pkg.Insert(set))
	case "merge":
		fmt.Println("Przed uporzadkowaniem: ", set)
		fmt.Print("Po uporzadkowaniu: ", pkg.Merge(set))
		fmt.Println("")
	default: // podano błędny
		fmt.Println("Wybierz", message)
	}

	time_final := time.Now().UnixNano() - time_start
	fmt.Println("Upłynęło czasu [nanosec]: ", time_final)
}
