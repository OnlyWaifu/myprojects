package main

import (
	"fmt"
	"log"
	"os"
	"unicode"
)

const size = 32

func main() {
	var slovo string
	var sdvig int32
	fmt.Print("Введите слово: ")
	fmt.Fscan(os.Stdin, &slovo)

	fmt.Print("Введите ключ: ")
	fmt.Fscan(os.Stdin, &sdvig)

	r := []rune(slovo)
	for i, a := range r {
		r[i] = shift(a, sdvig)
	}
	fmt.Println(string(r))
}

func shift(r rune, sdvig int32) rune {
	if !unicode.Is(unicode.Cyrillic, r) {
		log.Fatalln("Неверный формат букв")
	}

	if unicode.IsLower(r) {
		r += sdvig
		if r > 'я' {
			r -= size
		}
		return r
	} else {
		r += sdvig
		if r > 'Я' {
			r -= size
		}
		return r
	}
	// r = (r-base-1+sdvig)%size + base + 1
}
