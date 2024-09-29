/**
* default values
* int, float = 0;
* bool = false
* string = ''
 */
package main

import (
	"fmt"
	"reflect"
) // double quotes

func main() { // every go program should have main
	//var str string // str='' (default)
	//var str string = "hi how are you" // full form
	var str1 = "hi how are you" // type inferred - preffered
	str1 = "reassigned"         // can be reassinged
	str2 := "hi how are you"    // type inferred - go's way
	fmt.Println("aasdf")        // Exported functions, variables starts with Caps
	fmt.Println(str1)
	fmt.Println(reflect.TypeOf(str2))

	// default values
	var age int
	fmt.Println("default", age)

	// type conversion
	var float = 10.01
	var integer = int(float)
	fmt.Println(integer)

}
