package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	// var books [4]string
	// var books [1 + 3]string
	var books [yearly]string

	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golden"
	books[2] = "Everythingship"
	books[3] += books[0] + " 2nd Edition"

	// Go compiler can catch indexing errors when constant is used
	// books[4] = "Neverland"

	// Go compiler cannot catch indexing errors when non-constant is used
	// i := 4
	// books[i] = "Neverland"

	// --------------------
	// PRINTING ARRAYS
	// --------------------

	// print the type
	fmt.Printf("books  : %T\n", books)

	// print the elements
	fmt.Println("books  :", books)

	// print the elements in quoted string
	fmt.Printf("books  : %q\n", books)

	// print the elements along with their types
	fmt.Printf("books  : %#v\n", books)
}
