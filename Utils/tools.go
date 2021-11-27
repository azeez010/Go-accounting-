package utils

import (
	"fmt"
	"strconv"
	"errors"
)

var ErrNotNumber = errors.New("Not a Number")


func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func IsNum(v string) error {
	if _, err := strconv.Atoi(v); err == nil {
		// fmt.Printf("%q looks like a number.\n", v)
		return nil
	} else {
		return ErrNotNumber
	} 

}

func ToNumber(cost string, output *uint){	
	err := IsNum(cost)
	if err != nil {
		panic("Not a Number")
	} else {
		conv, _ := strconv.Atoi(cost)
		*output = uint(conv)
	}
}