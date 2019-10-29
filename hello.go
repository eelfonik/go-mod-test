package main

import (
	"fmt"
	"go-mod/test"
)

func main() {
	const msg = "Hello earth! \n"

	// [...]T syntax is sugar for [n]T
	// It creates a fixed size array, but lets the compiler figure out how many elements are in it.
	// There's no such thing as constant array in go

	// here the msgs are an array (which will be [2]string for compiler)
	msgs := [...]string{"How are you going? \n", "We're in trouble \n"}

	// But here the msgs1 are a slice 🌚👻
	// the only difference is the `...` inside `[]`
	msgs1 := []string{"How are you going? \n", "We're in trouble \n"}

	/*
		A few differences between array/slice in go:
		slice is built on array.
		- the length of an array is part of its type ([4]int and [5]int is not compatible)
		so the type of array is always [n]T, but the type of a slice is []T
		- Zero value for int array is 0, whereas zero value for int slice is nil
	*/

	// map is just a classic map
	msgs2 := map[int]string{1: "How are you going? \n", 2: "We're in trouble \n"}

	fmt.Println(test.Response(msg))
	fmt.Println(test.Response(msgs))
	fmt.Println(test.Response(msgs1))
	fmt.Println(test.Response(msgs2))

	// for loop is the only way to iterate sth in go
	// https://stackoverflow.com/questions/48693172/how-do-i-iterate-over-an-array-of-a-custom-type-in-go
	for _, msg := range msgs2 {
		fmt.Printf(test.Response(msg))
	}
}
