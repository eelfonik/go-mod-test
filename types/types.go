package types

import (
	"errors"
	"fmt"
)

// MyString is a named type that has a `Reverse` method
type MyString string

/*
N.B. if we try directly pass the basic type `string`,
we'll get an error, as a receiver of method cannot be basic/unnamed type
*/

// Reverse is a method with a value receiver that can reverse the order of named alias type `MyString`
func (myStr MyString) Reverse() string {
	s := string(myStr)
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

/*
struct(for data) & interface(for methods)
*/

type Actions interface {
	Response() (string, error)
}

// Alien is ... alien, and there're some alien you should not response their message
type Alien struct {
	Planet      string
	isResponded bool
}

// Animal is something we could pat
type Animal struct {
	Race        string
	isResponded bool
}

// implement Response method for Animal
func (sender *Animal) Response() (string, error) {
	if sender.isResponded {
		return "", errors.New("already responded")
	}
	sender.isResponded = true
	return sender.Race + " pat \n", nil
}

type alienDict struct {
	canResponse bool
	isResponded bool
}

var answerDict = map[string]alienDict{
	"prophet": {canResponse: false},
	"namic":   {canResponse: true},
}

// implement Response method for Alien
func (sender *Alien) Response() (string, error) {
	if sender.isResponded {
		return "", errors.New("already responded")
	}

	sender.isResponded = true
	// go has default zero value "" for string type
	if sender.Planet != "" {
		// the second value returns true if the key exits
		// otherwise will be false
		// see https://tour.golang.org/moretypes/22
		if alien, ok := answerDict[sender.Planet]; ok {
			if alien.canResponse {
				return "hello " + sender.Planet, nil
			}
			return "", errors.New(sender.Planet)
		}
		return "", errors.New(sender.Planet + ", we don't know them yet")
	}
	return "", errors.New("unknown alien")
}

func Response(sender Actions) {
	res, err := sender.Response()
	if err != nil {
		handler := errorHandler(func(msg string) {
			fmt.Println("don't answer " + msg)
		})
		handler.HandleSilence(err)
		// if in the func signature we defined return values with name
		// this return will return those values
		// it's called a "naked return"
		return
	}
	fmt.Println(res)
}

// A type can be a function signature
type errorHandler func(err string)

// A value pointer can be a function
func (fn errorHandler) HandleSilence(err error) {
	fn(err.Error())
}
