package main

import "fmt"

func hello() {

	fmt.Println("hello")

}

func world() {
	fmt.Println("world")
}

func makeEvenGenerator() func() int {

	i := 0
	return func() int {
		i += 2
		return i
	}
}

func visit(numbers []int, callback func(int)) {

	for _, n := range numbers {

		callback(n)

	}

}

func main() {

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) //2
	fmt.Println(nextEven()) //4
	fmt.Println(nextEven()) //6

	masEven := makeEvenGenerator()
	fmt.Println(masEven()) //2
	fmt.Println(masEven()) //4
	fmt.Println(masEven()) //6
	fmt.Println("callbacks->")
	visit([]int{1, 2, 3, 4}, func(n int) {

		fmt.Println(n)

	})
	defer hello()
	defer world()

	hello()
}
