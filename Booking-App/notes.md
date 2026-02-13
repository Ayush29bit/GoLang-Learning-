## bASICS OF GOLANG
Everything in Go is organised in packages 
func is the keyword for declaring a function in GO


Go's module system is about dependency management and code organization. Let me break down what's happening and why.

When we run "go mod init something" it creates a go.mod file which manages dependency as it acts like main module 

when we rin go run main.go, it goes thorugh:-
1.Finds your go.mod

2.Resolves dependencies

3.Builds a dependency graph

4.Compiles packages (bottom-up)

5.Links everything into a single binary
