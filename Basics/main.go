package main

import "fmt"

func main() {
	// Lesson 1: Hello, World!
	greeting := "Hello, World!"
	fmt.Println(greeting)

	// Lesson 2: Variables
	var a int = 10
	var b int = 20
	var c int = a + b
	fmt.Println(c)

	// Lesson 3: Constants
	const d int = 30
	const e int = 40
	const f int = d + e
	fmt.Println(f)

	// Lesson 4: Data Types
	var g int = 50
	var h float64 = 60.5
	var i string = "Hello, World!"
	var j bool = true
	fmt.Println(g, h, i, j)

	// Lesson 5: Operators
	var k int = 10
	var l int = 20
	var m int = k + l
	var n int = k - l
	var o int = k * l
	var p int = k / l
	var q int = k % l
	fmt.Println(m, n, o, p, q)

	// Lesson 6: Control Structures
	var r int = 10
	if r > 5 {
		fmt.Println("r is greater than 5")
	} else {
		fmt.Println("r is less than or equal to 5")
	}

	// Lesson 7: Loops
	for s := 0; s < 5; s++ {
		fmt.Println(s)
	}

	// Lesson 8: Functions
	fmt.Println("Add funtion response:", add(10, 20))

	// Lesson 9: Arrays
	var t [5]int
	t[0] = 10
	t[1] = 20
	t[2] = 30
	t[3] = 40
	t[4] = 50
	fmt.Println(t)

	// Lesson 10: Slices
	var u []int
	u = append(u, 10)
	u = append(u, 20)
	u = append(u, 30)
	u = append(u, 40)
	u = append(u, 50)
	fmt.Println(u)

	// Lesson 11: Maps
	v := make(map[string]int)
	v["a"] = 10
	v["b"] = 20
	v["c"] = 30
	v["d"] = 40
	v["e"] = 50
	fmt.Println(v)

	// Lesson 12: Structs
	type Person struct {
		Name string
		Age  int
	}
	w := Person{Name: "John", Age: 30}
	fmt.Println("Lesson 12: Structs:", w)

	// Lesson 13: Pointers
	x := 10
	y := &x
	fmt.Println("Lesson 13: Pointers:", x, y, *y)

	// Lesson 14: Interfaces
	type Shape interface {
		Area() int
	}
	var shape Shape
	shape = Circle{Radius: 10}
	fmt.Println("Lesson 14: Interfaces: Area of Circle:", shape.Area())

	// Lesson 15: Packages
	//fmt.Println("Lesson 15: Packages: Add funtion response:", add(10, 20))

	// Lesson 16: Error Handling
	/*z := 10
	if z > 5 {
		panic("z is greater than 5")
	}*/

}

type Circle struct {
	Radius int
}

func (c Circle) Area() int {
	return 3 * c.Radius * c.Radius

}

// Lesson 8: Add Funtion
func add(i1, i2 int) int {
	return i1 + i2
}
