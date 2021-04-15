package test

import (
	"fmt"
	"testing"
)

func Test_simple(test *testing.T) {
	itest := []int{}
	itest1 := itest
	itest1 = append(itest1, 1)
	fmt.Println(itest)
	fmt.Println(itest1)

	itest2 := &itest
	*itest2 = append((*itest2), 2)
	fmt.Println(itest)
	fmt.Println(itest2)
}
