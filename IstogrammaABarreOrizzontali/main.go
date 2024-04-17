package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"unicode"
)

var b bool

// Sort slice x Cresc/Decresc Alfabetico/Non Alfabetico
func Sort(value []int, keys []rune, str []string) ([]int, []rune, []string) {
	boolean := false
	ns := true
	for boolean == false {
		ns = true
		for i := 0; i < (len(value))-1; i++ {
			if value[i] > value[i+1] {
				value[i], value[i+1] = value[i+1], value[i]
				keys[i], keys[i+1] = keys[i+1], keys[i]
				str[i], str[i+1] = str[i+1], str[i]
				ns = false
			}
		}
		if ns == true {
			boolean = true
		}
	}
	return value, keys, str
}

// Sort Alfabetico x Crescente A-Z / Z-A/Decrescente A-Z / Z-A
func SortAlf(value []int, str []string) []string {
	boolean := true
	count := 0
	var store []string
	var result []string
	var x int
	for _, i := range value {
		if boolean == true {
			x = i
			boolean = false
		}
		if x == i {
			store = append(store, str[count])
			x = i
		} else {
			sort.Strings(store)
			if b == true {
				store = Swap(store)
				b = false
			}
			result = append(result, store...)
			store = nil
			store = append(store, str[count])
			x = i
		}
		if len(value)-1 == count {
			result = append(result, str[count])
		}
		count++
	}
	return result
}

// Swap slice String
func Swap(str []string) []string {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return str
}

// Stampa
func Stampa(keys []string, mappa map[rune]int) {
	for _, k := range keys {
		x := ""
		h := rune(k[0])
		//h := RuneConverter(k) //Un tempo usavo un convertitore da stringa a runa fatto da me
		for i := 0; i < mappa[h]; i++ {
			x += "*"
		}
		Printf("%c: %s\n", h, x)
	}
}

// Input testo
func LeggiTesto(Testo string) string {
	Print("Inserisci testo (termina con 'CTRL+D' o '/'): ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "/" {
			break
		} else {
			Testo += scanner.Text()
		}
	}
	return Testo
}

// Calcolo numero Occorrenze
func Occorrenze(s string) map[rune]int {
	mappa := make(map[rune]int)
	for _, i := range s {
		if unicode.IsLetter(i) {
			mappa[i] += 1
		}
	}
	return mappa
}

// Default
func Default(mappa map[rune]int) {
	Println("Stampa : Default")
	for k := range mappa {
		x := ""
		for i := 0; i < mappa[k]; i++ {
			x += "*"
		}
		Printf("%c: %s\n", k, x)
	}
}

// Alfabetico A-Z
func AlfabeticoAZ(mappa map[rune]int) {
	Println("Stampa : Alfabetico A-Z")
	keys := make([]string, 0, len(mappa))
	for k := range mappa {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	Stampa(keys, mappa)
}

// Alfabetico Z-A
func AlfebeticoZA(mappa map[rune]int) {
	Println("Stampa : Alfabetico Z-A")
	keys := make([]string, 0, len(mappa))
	for k := range mappa {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	keys = Swap(keys)
	Stampa(keys, mappa)
}

// Crescente
func Crescente(mappa map[rune]int) {
	Println("Stampa : Numero di occorrenze crescente")
	keys := make([]rune, 0, len(mappa))
	value := make([]int, 0, len(mappa))
	var str []string
	for k := range mappa {
		keys = append(keys, k)
		value = append(value, mappa[k])
		str = append(str, string(k))
	}
	value, keys, str = Sort(value, keys, str)
	Stampa(str, mappa)
}

// Decrescente
func Descrescente(mappa map[rune]int) {
	Println("Stampa : Numero di occorrenze decrescente")
	keys := make([]rune, 0, len(mappa))
	value := make([]int, 0, len(mappa))
	var str []string
	for k := range mappa {
		keys = append(keys, k)
		value = append(value, mappa[k])
		str = append(str, string(k))
	}
	value, keys, str = Sort(value, keys, str)
	str = Swap(str)
	Stampa(str, mappa)
}

// Crescente Alfabetico A-Z
func CrescenteAlfAZ(mappa map[rune]int) {
	Println("Stampa : Numero di occorrenze crescente alfabetico A-Z")
	keys := make([]rune, 0, len(mappa))
	value := make([]int, 0, len(mappa))
	str := make([]string, 0, len(mappa))
	b = false //x riutilizzare la func SortAlf()
	for k := range mappa {
		keys = append(keys, k)
		value = append(value, mappa[k])
		str = append(str, string(k))
	}
	value, keys, str = Sort(value, keys, str)
	str = SortAlf(value, str)
	Stampa(str, mappa)
}

// Crescente Alfabetico Z-A
func CrescenteAlfZA(mappa map[rune]int) {
	keys := make([]rune, 0, len(mappa))
	value := make([]int, 0, len(mappa))
	str := make([]string, 0, len(mappa))
	b = true //x riutilizzare la func SortAlf()
	for k := range mappa {
		keys = append(keys, k)
		value = append(value, mappa[k])
		str = append(str, string(k))
	}
	value, keys, str = Sort(value, keys, str)
	str = SortAlf(value, str)
	Stampa(str, mappa)
}

// Decrescente Alfabetico A-Z
func DecrescenteAlfAZ(mappa map[rune]int) {
	Println("Stampa : Numero di occorrenze decrescente alfabetico A-Z")
	keys := make([]rune, 0, len(mappa))
	value := make([]int, 0, len(mappa))
	str := make([]string, 0, len(mappa))
	b = true //x riutilizzare la func SortAlf()
	for k := range mappa {
		keys = append(keys, k)
		value = append(value, mappa[k])
		str = append(str, string(k))
	}
	value, keys, str = Sort(value, keys, str)
	str = SortAlf(value, str)
	Swap(str)
	Stampa(str, mappa)
}

// Decrescente Alfabetico Z-A
func DecrescenteAlfZA(mappa map[rune]int) {
	Println("Stampa : Numero di occorrenze decrescente alfabetico Z-A")
	keys := make([]rune, 0, len(mappa))
	value := make([]int, 0, len(mappa))
	str := make([]string, 0, len(mappa))
	b = false //x riutilizzare la func SortAlf()
	for k := range mappa {
		keys = append(keys, k)
		value = append(value, mappa[k])
		str = append(str, string(k))
	}
	value, keys, str = Sort(value, keys, str)
	str = SortAlf(value, str)
	Swap(str)
	Stampa(str, mappa)
}

// Lettera singola
func LetteraSingola(mappa map[rune]int) {
	Println("Stampa : Numero occorrenze lettara")
	Print("Inserisci lettera da stampare --> ")
	var in string
	Scan(&in)
	k := rune(in[0])
	//k := RuneConverter(in) //Un tempo usavo un convertitore da stringa a runa fatto da me
	if unicode.IsLetter(k) {
		if mappa[k] != 0 {
			x := ""
			for i := 0; i < mappa[k]; i++ {
				x += "*"
			}
			Printf("%c: %s\n", k, x)
		} else {
			Println("La lettera inserito non è presente nel testo!")
		}
	} else {
		Println("Hai inserito un input non valido!")
	}
}

// Main
func main() {
	Testo := ""
	Testo = LeggiTesto(Testo)
	mappa := Occorrenze(Testo)
	if len(mappa) == 0 {
		Println("Hai inserito non valido / troppo corto!")
	} else {
		Println("Come vuoi stampare?")
		var in int
		Println("Alfabetico A-Z: 1")
		Println("Alfabetico Z-A: 2")
		Println("Numero di occorrenze crescente : 3")
		Println("Numero di occorrenze decrescente : 4")
		Println("Numero di occorrenze crescente alfabetico A-Z: 5")
		Println("Numero di occorrenze crescente alfabetico Z-A: 6")
		Println("Numero di occorrenze decrescente alfabetico A-Z: 7")
		Println("Numero di occorrenze decrescente alfabetico Z-A: 8")
		Println("Numero occorrenze lettera: 9")
		Println("Altro : Default")
		Print("Input --> ")
		Scan(&in)
		switch in {
		case 1:
			AlfabeticoAZ(mappa)
		case 2:
			AlfebeticoZA(mappa)
		case 3:
			Crescente(mappa)
		case 4:
			Descrescente(mappa)
		case 5:
			CrescenteAlfAZ(mappa)
		case 6:
			CrescenteAlfZA(mappa)
		case 7:
			DecrescenteAlfAZ(mappa)
		case 8:
			DecrescenteAlfZA(mappa)
		case 9:
			LetteraSingola(mappa)
		default:
			Default(mappa)
		}
	}
}
