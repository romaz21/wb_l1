package main

import (
	"fmt"
	"reflect"
)

func getType(value interface{}) string {
	switch reflect.ValueOf(value).Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "channel"
	default:
		return "unknown"
	}
}

func main() {
	var intValue int = 42
	var stringValue string = "Hello, World!"
	var boolValue bool = true
	var channelValue chan int

	fmt.Println(getType(intValue))
	fmt.Println(getType(stringValue))
	fmt.Println(getType(boolValue))
	fmt.Println(getType(channelValue))
}

