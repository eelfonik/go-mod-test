package test

import (
	"reflect"
	"strconv"
	"time"
)

// How to declare union type in go ??

// anything started with uppercase letter in a package is automatically exported
func Response(msg interface{}) string {
	/* For Variables
	Option 1:
	var currentTime int64
	currentTime = time.Now().Unix()
	Option 2:
	(type is infered on initialization, here will be int64)
	var currentTime = time.Now().Unix()
	Option 3:
	(the `:=` is a shorthand for variable declaration and initialization)
	currentTime := time.Now().Unix()
	*/

	currentTime := "replace the time"

	// the memory address (via using `&`) of variable `currentTime`
	p := &currentTime
	// println(p)

	// point the memory address(p) via `*` to the real time string
	*p = strconv.FormatInt(time.Now().Unix(), 10)
	// afterwards anywhere we use the variable `currentTime` will be pointed to this above string value

	// const has no type until given one
	// and the value must be decided in compile time
	const res = "hello alien "

	// function def inside function: assign the func to a local variable
	handleNonPrimitive := func(val interface{}) string {
		// as type switch only work for specific type, for array/slice/map, you need to use reflect
		// it's like the metadata of var
		v := reflect.ValueOf(val)
		switch v.Kind() {
		case reflect.Slice:
			// how to check the element type & value inside an slice ?
			return "weird thing here we got a slice"
		case reflect.Array:
			return "An array!!"
		case reflect.Map:
			return "The alien passing a map"
		default:
			return "nothing here"
		}
	}

	// use interface for type switch
	// see https://tour.golang.org/methods/16
	// note the `.(type)` here
	switch msgType := msg.(type) {
	case string:
		return msgType + res + currentTime
	case int64:
		return strconv.FormatInt(msgType, 10) + res + currentTime
	default:
		return handleNonPrimitive(msgType)
	}
}

func handleNonPrimitive(val interface{}) string {
	return "nothing outside"
}
