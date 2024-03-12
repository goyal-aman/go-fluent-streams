package main

import (
	"fmt"
	"goyal-aman/go-fluent-streams/streams"
)

func main() {
	result := make([]int, 0)
	resultCollector := func(i *int) {
		if i != nil {
			result = append(result, *i)
		}
	}
	streams.Of([]int{1, 2, 3}).
		Map(func(i int) *int {
			retVal := i * 2
			return &retVal
		}).
		Filter(func(i int) bool {
			return i%2 == 0
		}).
		Collect(resultCollector)
	fmt.Println(result)
}
