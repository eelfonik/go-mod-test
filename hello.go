package main

import (
	"fmt"
	"go-mod/test"
)

func main() {
	// const has no type until given one
	const msg = "Hello earth! \n"
	fmt.Println(test.Response(msg))
}
