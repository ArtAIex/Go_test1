package test

import "fmt"

func Cout(a []int) {
	for _, elem := range a {
		fmt.Print(elem)
	}
}
