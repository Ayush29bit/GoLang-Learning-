## bASICS OF GOLANG
Everything in Go is organised in packages 
func is the keyword for declaring a function in GO

Go's standard library provides different core packages for us to use 
(fmt is one of these core packages which we use by importing)

Packages can be though of containers of different functionalties which go gives us to usedddddd

Go's module system is about dependency management and code organization. Let me break down what's happening and why.

When we run "go mod init something" it creates a go.mod file which manages dependency as it acts like main module 

when we rin go run main.go, it goes thorugh:-
1.Finds your go.mod

2.Resolves dependencies

3.Builds a dependency graph

4.Compiles packages (bottom-up)

5.Links everything into a single binary

## Variables 

 var name string 
 var age int 
 var is Active bool

# var is the most basic way to declare a variable in go

the default value of var variables is a zero value( " ", 0, false), and can be intialised 

var age int = 18
var name string = "John Doe"

var name string = "John doe" -: This is a variable declared with an explicit variable type already specified by us 

var name = "John doe" -: This is a varible declared using type inference in go where go infers the type from value

When we declare a variable and we assign their value, it is no problem as Go can already infer the type of the variable

When we only declare the variable and wish to assign the value later, we need to declare the type of the variable



# Variable declaration
var name string
var age int
var isActive bool
These create variables with zero values (empty string "", 0, false)

go
func main() {
    name := "John"
    age := 25
    isActive := true
}
This is the most common way to declare variables in Go. The := operator declares and assigns in one step. You cannot use := outside of functions (package-level variables must use var).

## Printf function
This function is part of the fmt package
Printf=Print+format

It lets you add format specifier and then pass variable name as argument in the Printf function, thus allowing us to add varibles in string, control formatting and explicitly specify types

Verb	Meaning
%v	default format
%T	type of variable
%s	string
%d	integer (base 10)
%f	float
%t	boolean
%p	pointer address
%x	hex
%b	binary

### GO IS A STATICALLY TYPED LANGUAGE, THE DATA TYPE SHOULD BE TOLD WHEN DECLARING A VARIABLE 

## Input to app 
In order to take the input from a user, we use the scan() function from the same fmt package 

mportant Limitation

fmt.Scan()-:
Stops at whitespace
Not ideal for full sentences
Not ideal for robust input handling
If you want full line input, use Scanln or bufio.

## Pointers
Pointers in go are called special variables and are used to point to the value of other variable in the memory essentially serving as memory address

&name==points to the address in the memory where the value of name is stored 

## Arrays and Slices 

An array in Go:

Fixed size
Part of the type
Stored contiguously in memory

[x]int is how the type of an array is defined 

var a[3]int 
var b[4]int

a = b // False because they have different types 

#### Slices are abstraction of arrays in Go 
They are more flexible and powerful:
variable length or get an subarray of its own

var a[] int 
we dont define the size of the array and it becomes a slice
It uses append() method to add values into it 

var a[] int 
a = append(a,"apple")

## Loops in GO 
Go has only one loop keyword:

for
That’s it.
No while.
No do-while.
No foreach.

Standard for loop :

for i := 0; i < 5; i++ {
    fmt.Println(i)
}

While Style Loop:

i := 0
for i < 5 {
    fmt.Println(i)
    i++
}

Infine Loop:

for {
    fmt.Println("Running forever")
}

##### Looping Over Strings
str := "hello"

for index, char := range str {
    fmt.Println(index, char)
}

#### Looping Over Slices / Arrays (range)
nums := []int{10, 20, 30}

for index, value := range nums {
    fmt.Println(index, value)
}

#### Important 

for _, value := range nums {
    fmt.Println(value)
}
The _ means: “Ignore this value.”
Go forces you to use variables so _ discards them.

## Conditionals in GO

Go mainly gives you:
if
if-else
switch

##### Basic if Statement
age := 20

if age >= 18 {
    fmt.Println("Adult")
}

#### if-else
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}


#### else if
score := 85

if score >= 90 {
    fmt.Println("A")
} else if score >= 75 {
    fmt.Println("B")
} else {
    fmt.Println("C")
}

#### switch Statement

Go’s switch is more powerful than most languages.
Basic form:

day := "Monday"

switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("Weekend soon")
default:
    fmt.Println("Regular day")
}

