package main

import (
	"fmt"
	"runtime"
)

func main() {
	printEnvironmentVars()
	printRuntimeFunctions()
}

func printEnvironmentVars() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("ROOT:", runtime.GOROOT())
}

func printRuntimeFunctions() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	fmt.Println("Go version:", runtime.Version())
}
