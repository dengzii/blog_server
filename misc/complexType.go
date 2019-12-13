package misc

import "fmt"

type Int int64

func (i Int) add(i2 int64) int64 {
	return int64(i) + i2
}

func typeVar() {
	a := Int(4)
	fmt.Println(a.add(2))
}

type error interface {
	Error() string
}

type DivideError struct {
	msg string
}

func (er DivideError) Error() string {
	return er.msg
}
func exception(a int32, b int32) (int32, error) {
	if b == 0 {
		return 0, DivideError{msg: "divide by 0."}
	}
	return a / b, nil
}

type Animal interface {
	eat(food string)
	weight() int
}
type Dog struct {
}

func (d Dog) eat(food string) {
	fmt.Println(food)
}

func (d Dog) weight() int {
	return 10
}

func interface_() {
	var ani Animal
	ani = new(Dog)
	fmt.Println("interface Animal")
	ani.eat("fish")
	fmt.Println(ani.weight())
}

func map_() {
	map1 := make(map[string]string)
	map1["a"] = "1"
	map1["b"] = "2"
	map1["c"] = "3"

	for key := range map1 {
		fmt.Println(map1[key], ":", key)
	}

	value, exist := map1["c"] // check if key exist
	fmt.Println(value, exist)
	delete(map1, "c") // delete element
	fmt.Println(map1)
}

func range_() {
	nums := []int{1, 2, 3, 4, 5, 6}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println("ele=", i, num)
	}
}

func slice() {
	var ind []int        // definition
	ind = make([]int, 3) // init

	s := []int32{1, 2, 3, 4, 5, 7} // definition and  init
	fmt.Println("slice=", ind)
	fmt.Println("slice2=", s)
	fmt.Println("slice2=", append(ind, 11, 22, 33)) // append ele to slice
	fmt.Println("slice3=", s[1:4])                  // slice
}

type Man struct {
	NAME string
	AGE  int
}

func struct_() {
	var bob = Man{NAME: "bob", AGE: 22}
	fmt.Println("struct=", Man{"Bob", 25})
	fmt.Println("man=", bob.NAME, bob.AGE)
}

func pointer2(a int32, b *int32) {
	a = 1
	*b = 2 // pointer reference
}

func pointer() {
	var p *int // a pointer
	fmt.Println(p)

	a := 1
	var p2 = &a
	fmt.Println(a)
	fmt.Println("pointer of a=", p2)

	*p2 = 2
	fmt.Println(a)
}

func main() {
	pointer()
	var a int32 = 3
	var b int32 = 4
	pointer2(a, &b) // b will change
	fmt.Println("pointer2=", a, b)
	struct_()
	typeVar()
	slice()
	range_()
	map_()
	interface_()
	fmt.Println(exception(11, 0))
}
