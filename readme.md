## Go Features
- Compiled Language 
- Typed 
- Garbage collection 
- Concurrency(Goroutines) 
- Platform agnostic (single code base builds to multiple target platforms)
- High Performant runtime (Low CPU - multicore usage, Low memory, High CPU throughput, smaller exectable?, smaller runtime?)
- Builtin libs/packages - Code format-styling `go fmt`, Test

## Program structure & Execution
- Package clause
- Import statement
- Actual code 

```
/* module this file/ our code belongs to */
package main

import "fmt"

/**
main is called by the runtime for any program
*/
func main() {
  fmt.Println("Hello world")
}
```
## CLI Options
Commands
```
$ go build main.go // compile - build(exectable)
$ go run main.go // compile - run on the fly

$ go fmt  // set editor to fmt on save
$ go test 
```
## Types - declarations - default values - Type conversion
Full syntax
````
var str1 string = "Hello world"
```
Type inferred syntax (preferred)

> Note: we have `const` in go. but that is used more like *enums* and for hardcoded values
> const ENV="prod"

```
### Types and its default values 

var integerNumber = 1;  // 0 
var floatNumber = 1.01; // 0
var string = "asf" // ""
var boolean = true // false
```
For Pointers, Structs, Map  - default value is `nil`

### Type conversion
To convert from one type to another use the type keyword `int, float64, bool, string`
```
var floatNumber = 3.14;
var integer = int(floatNumber)
```
## Read input from stdin


