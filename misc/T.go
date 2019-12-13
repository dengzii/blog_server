package misc

import (
	"fmt"
	"math"
	"time"
)

func varFunc() {

	f1 := func(a, b string) string {
		return a + b
	}
	fmt.Println("fun var ", f1("a", "b"))
}

func d(i int) {
	fmt.Print(i)
}

func deferStatement() {
	fmt.Println("defer...")
	defer d(1)
	d(2)
	defer fmt.Println("world")
	fmt.Println("hello")
}

func switchStatement() {
	fmt.Println("switch statement")
	i := 1
	switch i {
	case 1:
		fmt.Println(1)
		fallthrough // not break, go through
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func forStatement() {
	fmt.Println("for statement")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	j := 0
	for j < 7 { // while(j<7)
		j++
	}
	fmt.Println(j)
}

func typeConvert() {
	var a float64 = math.Pi
	var b float32 = float32(a)
	var c uint16 = uint16(b)
	fmt.Println(a, b, c)
}

var (
	MaxInt uint64 = 1<<64 - 1
	Enable bool   = false
) // declare variable out function

func variable2() {
	var b [2]int32
	b = [2]int32{2, 32}
	fmt.Println(b)
	fmt.Println("Max int=", MaxInt)
	fmt.Println(Enable)
}

func variable() {
	var a, b, c int                    // the same type
	var x, y string = "hello", "world" // init variable value
	z := 123.123                       // auto type
	const cons = "const"               // const can not use := dec

	fmt.Println(a, b, c, x, y, z, cons)
}

// a function
func fun1(a, b int) int {
	return a + b
}

// function with two return param
func swap(a, b int) (int, int) {
	return b, a
}

// function with named return param
func split(num int) (t, s int) {
	t = num / 10
	s = num - t*10
	return // t,s
}

func main() {
	fmt.Println("Hello World")
	fmt.Println("Time ", time.Now())
	fmt.Println(math.Pi)
	fmt.Println("1+2=", fun1(1, 2))

	a, b := swap(1, 2)
	fmt.Println("swap ", a, b)
	fmt.Println(split(42))
	variable()
	variable2()
	typeConvert()
	forStatement()
	switchStatement()
	deferStatement()
	varFunc()
}
