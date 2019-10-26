package test

import (
	"strconv"
	"time"
)

func Response(msg string) string {
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
	currentTime := strconv.FormatInt(time.Now().Unix(), 10)
	const res = "hello alien "
	return msg + res + currentTime
}
