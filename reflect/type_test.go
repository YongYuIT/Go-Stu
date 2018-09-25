package reflect

import (
	"testing"
	"fmt"
)

type IStudent interface {
	getName() string
	setName(_name string)
}

type SchoolChild struct {
	name        string
	age         int
	school_name string
}

func (this *SchoolChild) getName() string {
	return this.name
}

func (this *SchoolChild) setName(_name string) {
	this.name = _name
}

type MiddleSchoolStudent struct {
	name               string
	age                int
	middle_school_name string
}

func (this MiddleSchoolStudent) getName() string {
	return this.name
}

func (this MiddleSchoolStudent) setName(_name string) {
	this.name = _name
}

func TestType(t *testing.T) {
	var yong IStudent = &SchoolChild{"yuyong", 10, "USC"}
	fmt.Println(yong.getName())

	var guo IStudent = MiddleSchoolStudent{"guoqing", 20, "HMYZ"}
	fmt.Println(guo.getName())
}

func TestCopy(t *testing.T) {

	yong := MiddleSchoolStudent{"yuyong", 10, "USC"}
	var i_yong IStudent = yong
	i_yong.setName("feifei")
	fmt.Println(yong.getName())

	guo := SchoolChild{"guoqing", 20, "HMYZ"}
	var i_guo IStudent = &guo
	i_guo.setName("benben")
	fmt.Println(guo.getName())

}
