package main

import "fmt"

/*
func main() {
	x := 42
	fmt.Println(x)
	// using the "&"" for getting pointer address
	fmt.Println(&x)

	// using "*" for getting the values from a pointer address
	y := &x
	fmt.Println(*y)
	fmt.Println(*&x)

	//note, the "%T" is used get the pointer type
	fmt.Printf("%v\t%T\n", y, y)
}

// hands on exersice 1
var coachela string = "The Weeknd"

func main() {
	fmt.Println(&coachela)
}

//hands on exersice 2
var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	//print the "value" stored in each variable
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//print the "type" of each variable
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

	//print the data stored at memory loctions
	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)

}

// hands on exercise 3

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walk.")
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	youngRun(d1)

	d2 := dog{"Padget"}
	youngRun(d2)

}
*/

// hands on exercise 4 value pointer semantics

type person struct {
	first string
}

func main() {
	p := person{
		first: "Jenny",
	}
	fmt.Println(p)
	p = changeName(p, "Jen")
	fmt.Println(p)

	changeNamePtr(&p, "JMo")
	fmt.Println(p)
}

func changeName(p person, s string) person {
	p.first = s
	return p
}

func changeNamePtr(p *person, s string) {
	p.first = s
}
