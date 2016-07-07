package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type customInt int

func (c customInt) ToString() string {
	return strconv.Itoa(int(c))
}

func main() {
	var i customInt
	i = 25
	fmt.Println(i.ToString(), reflect.TypeOf(i.ToString()))
}
