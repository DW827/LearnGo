package main

import (
	"fmt"
)

func main() {
	nums := []int{-2,0,3,-5,2,-1}
	newnums := make([]int, len(nums))
	newnums[0] = nums[0]
	for i := 1; i < len(nums); i ++ {
		newnums[i] = newnums[i-1] + nums[i]
	}
	fmt.Printf("%#v", newnums)
}