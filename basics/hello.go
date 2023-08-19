package main

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

var a int

func main() {

	fmt.Println(a)
	var i int = 1
	var w float64 = 12.5

	fmt.Println(i, w)

	a := 30

	// taking a string variable
	Language := "Go Programming"

	fmt.Println("The Value of a is: ", a)
	fmt.Println("The Value of Language is: ", Language)
	var j, k = 1, 2

	var (
		name = "Sangam Biradar"

		occupation = "Developer Advocate"
	)

	fmt.Println(j, k)

	fmt.Printf("%s is a %s\n", name, occupation)
	fmt.Println(reflect.TypeOf(name))
	var file string
	_, file = path.Split("css/main.css")
	fmt.Println("file", file)
	var score int // already 0

	var (

		// related

		video string

		// Closely related

		duration int

		current int
	)
	fmt.Println(score, video, duration, current)
	speed := 100 // speed is int

	force := 2.5 // force is float64

	speed = speed * int(force)
	speed2 := float64(speed) * force
	fmt.Println(speed)
	fmt.Println(speed2)
	fmt.Println(os.Args)
	var s string

	s = "how are you?"

	s = `how are you?`

	fmt.Println(s)

	s = "<html> \n\t <body> \"hello\"</html>"

	s = `
	
<html>
	
<body> "hello" </body>
	
</html> `

	fmt.Println(s)

	fmt.Println("c:\\my\\dir\\file")

	fmt.Println(`c:\my\dir\file`)
	name2 := "संगम"

	fmt.Println(

		utf8.RuneCountInString(name2))
	fmt.Println((len(name)))
	msg := os.Args[0]

	l := len(msg)

	s2 := msg + strings.Repeat("!", l)
	s2 = strings.ToUpper(s2)
	fmt.Println(s2)
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)

	fmt.Println(EST, MST, PST)

	var brand = "Google"

	fmt.Printf("%q\n", brand)
	score = 21
	// score, valid := 3, true
	if score > 3 { // && valid
		fmt.Println("good")
	} else if score == 3 {
		fmt.Println("on the edge")
	} else if score == 2 {
		fmt.Println("meh...")
	} else {
		fmt.Println("low")
	}
	s = strconv.Itoa(42)

	fmt.Println(reflect.TypeOf(s), s)
	n, err := strconv.Atoi(os.Args[1])
	fmt.Println("Converted Number :", n)
	fmt.Println("Returned error value :", err)
	if err != nil {
		fmt.Printf("cannot convert %q \n", os.Args[1])
	}

	city := os.Args[1]
	switch city {
	case "Chandigarh":
		fmt.Println("India")
	case "tokyo":
		fmt.Println("japan")
	default:
		fmt.Println("None of known cities")
	}
	i = 10
	switch {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
	for _, v := range os.Args[1:] {

		fmt.Printf("%q\n", v)

	}
}
