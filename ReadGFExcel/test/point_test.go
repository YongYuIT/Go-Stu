package test

import (
	"../excel"
	"fmt"
	"testing"
)

func TestNextPoint(t *testing.T) {
	p1 := excel.Point{"AZ", "1"}
	p2 := p1.GetNextX()
	p3 := p2.GetNextX()
	p4 := p3.GetNextY()
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
}
