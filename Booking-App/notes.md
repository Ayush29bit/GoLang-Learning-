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


