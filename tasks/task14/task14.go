package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу,
которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
Я использовал typeSwitch, и reflect для определения типа в случае, если не совпал ни один из предыдущих случаев
*/
func getType(value interface{}) string {
	switch v := value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	default:
		return reflect.TypeOf(v).String()
	}
}

func main() {
	var intValue int = 42
	var stringValue string = "Hello, Go!"
	var boolValue bool = true
	var channelValue chan int = make(chan int)

	fmt.Println(getType(intValue))
	fmt.Println(getType(stringValue))
	fmt.Println(getType(boolValue))
	fmt.Println(getType(channelValue))
}
