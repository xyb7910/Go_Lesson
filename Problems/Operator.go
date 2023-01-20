package main

import "fmt"

func main() {
	var a, b int
	a, b = 1, 2
	c := a + b
	d := a - b
	e := a * b
	f := a / b
	g := a % b
	a++
	b--
	fmt.Println(c, d, e, f, g, a, b)

	var v1, v2 = 4, 5
	if v1 == v2 {
		fmt.Println("v1 == v2")
	} else {
		fmt.Println("v1 != v2")
	}
	if v1 < v2 {
		fmt.Println("v1 < v2")
	} else {
		fmt.Println("v1 !< v2")
	}

	if v1 > v2 {
		fmt.Println("v1 > v2")
	} else {
		fmt.Println("v1 !> v2")
	}

	var s1, s2 bool
	s1, s2 = true, false
	if s1 == s2 {
		fmt.Println("s1 == s2")
	} else {
		fmt.Println("s1 != s2")
	}
	if s1 != s2 {
		fmt.Println("s1 != s2")
	} else {
		fmt.Println("s1 == s2")
	}
	if s1 && s2 {
		fmt.Println("s1,s2 both true")
	} else {
		fmt.Println("s1,s2 another true")
	}
	if s1 || s2 {
		fmt.Println("s1,s2 another true")
	} else {
		fmt.Println("s1,s2 both true")
	}
}
