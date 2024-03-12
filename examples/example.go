package main

import (
	"fmt"
	"goyal-aman/go-fluent-streams/streams"
)

func main() {

	// result will store final result of funcational chaining
	result := make([]int, 0)

	streams.Of([]int{1, 2, 3}).
		Map(double).
		Filter(isEven).
		Collect(func(i *int) {
			if i != nil {
				result = append(result, *i)
			}
		})
	fmt.Println(result)
}
func double(i int) *int {
	retVal := i * 2
	return &retVal
}
func isEven(i int) bool {
	return i%2 == 0
}
