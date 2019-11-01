package main

import (
	"fmt"
	"go-mod/data"
	"go-mod/types"
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
	msgs2 := map[string]string{"index1": "How are you going? \n", "index2": "We're in trouble \n"}

	fmt.Println(data.Response(msg))
	fmt.Println(data.Response(msgs))

	// for loop is the only way to iterate sth in go
	// https://stackoverflow.com/questions/48693172/how-do-i-iterate-over-an-array-of-a-custom-type-in-go

	// iterate array
	for _, msg := range msgs {
		fmt.Printf(data.Response(msg))
	}

	fmt.Println(data.Response(msgs1))
	// iterate slice
	for _, msg := range msgs1 {
		fmt.Printf(data.Response(msg))
	}

	fmt.Println(data.Response(msgs2))
	// iterate map (map don't guarantee order)
	for _, msg := range msgs2 {
		fmt.Printf(data.Response(msg))
	}

	/*
		Using the types package
	*/
	// We define a varibale holds the custom defined `MyString` type
	reversable := types.MyString("shanghai")
	fmt.Println(reversable.Reverse())

	// Here came the interface usage
	// both of those implemented `Response` method of `Actions` interface
	dog := types.Animal{Race: "dog"}
	prophet := types.Alien{Planet: "prophet"}
	namic := types.Alien{Planet: "namic"}
	newAlien := types.Alien{Planet: "new"}
	mystery := types.Alien{}

	types.Response(&dog)
	types.Response(&prophet)
	types.Response(&namic)
	types.Response(&newAlien)
	types.Response(&mystery)

	// we won't pat a dog 2 times 😂
	types.Response(&dog)
}
