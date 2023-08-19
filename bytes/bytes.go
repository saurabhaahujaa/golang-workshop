package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	var start, stop int

	if args := os.Args[1:]; len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	}

	if start == 0 || stop == 0 {
		start, stop = 'A', 'Z'
	}

	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n",
		"literal", "dec", "hex", "encoded",
		strings.Repeat("-", 45))

	for n := start; n <= stop; n++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", n, string(n))
	}
	const text = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.
	
	Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.
		
	Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

	r, size := utf8.DecodeRuneInString("öykü")

	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("ykü")

	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("kü")

	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("ü")

	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	// for range loop automatically decodes the runes

	//   but it gives you the position of the current rune

	//   instead of its size.

	// for _, r := range text {}

	for i := 0; i < len(text); {

		r, size := utf8.DecodeRuneInString(text[i:])

		fmt.Printf("%c", r)

		i += size

	}
}
