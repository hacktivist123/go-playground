# go-playground

Where I try to play and understand the Go programming language

## Go Intro Notes

- Basic Variables

  - bool: a boolean value, either true or false
  -  string: a sequence of characters
  - int: a signed integer
  - float64: a floating-point number
  - byte: 8 bits of data

- The walrus operator, `:=`, declares a new variable and assigns a value to it in one line. Go can infer that mySkillIssues is an int because of the 42 value. 

- The limitation is that := can't be used outside of a function

- `package main` lets the Go compiler know that we want this code to compile and run as a standalone program, as opposed to being a library that's imported by other programs.

- Integers, uints, floats, and complex numbers all have type sizes.

- Whole numbers (No Decimal) - int  int8  int16  int32  int64

- Positive whole numbers (No Decimal) - uint uint8 uint16 uint32 uint64 uintptr

- Signed Decimal Numbers - float32 float64

- What's the Deal With the Sizes? The size (8, 16, 32, 64, 128, etc) represents how many bits in memory will be used to store the variable. The "default" int and uint types refer to their respective 32 or 64-bit sizes depending on the environment of the user.

The "standard" sizes that should be used unless you have a specific performance need (e.g. using less memory) are:

    int
    uint
    float64
    complex128

- Constants can be primitive types like strings, integers, booleans and floats. They can not be more complex types like slices, maps and structs

- Constants must be known at compile time. They are usually declared with a static value

- That said, you cannot declare a constant that can only be computed at run-time like you can in JavaScript.

- `fmt.Printf` - Prints a formatted string to standard out.
- `fmt.Sprintf()` - Returns the formatted string

- The %v variant prints the Go syntax representation of a value, it's a nice default.

    ```go
    s := fmt.Sprintf("I am %v years old", 10)
    // I am 10 years old
    ```
- string: `s := fmt.Sprintf("I am %s years old", "way too many")`

- Integer: `s := fmt.Sprintf("I am %d years old", 10)`

- float: `s := fmt.Sprintf("I am %f years old", 10.523)`

- 
