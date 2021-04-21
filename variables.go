package main

import (
	"fmt"
	"strconv"
)

//Lowercase variable scope inside the package
//Uppercase variable scope outside the package

var xy int = 20
var (
	name   string  = "Vikas"
	age    int     = 30
	height float32 = 5.6
)

func main() {
	var x int = 10
	var y int
	y = 20
	z := 30
	fx := 34.54
	var (
		name1 string = "Vikas"
		age1  int    = 30
	)
	//shadowing
	xy := 43

	fmt.Printf("x:%d Type: %T\n", x, x)
	fmt.Printf("x:%v Type: %T\n", y, y)
	fmt.Printf("x:%v Type: %T\n", z, z)
	fmt.Printf("x:%v Type: %T\n", fx, fx)
	fmt.Printf("x:%v Type: %T\n", xy, xy)
	fmt.Printf("x:%v Type: %T\n", name, name)
	fmt.Printf("x:%v Type: %T\n", age, age)
	fmt.Printf("x:%v Type: %T\n", name1, name1)
	fmt.Printf("x:%v Type: %T\n", age1, age1)
	fmt.Printf("x:%v Type: %T\n", height, height)
	fmt.Printf("x:%v Type: %T\n", xy, xy)

	//Typecasting
	var i int = 20
	var j float32
	j = float32(i)
	var k string
	k = strconv.Itoa(i)
	fmt.Printf("x:%v Type: %T\n", j, j)
	fmt.Printf("x:%v Type: %T\n", k, k)

}
