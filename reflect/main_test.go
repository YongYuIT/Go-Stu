package reflect

import (
	"testing"
	"strconv"
	"fmt"
	"reflect"
)

type student interface {
	getAge() int
	getName() string
}

type schoolchild struct {
	name        string
	age         int
	school_name string
}

func (this *schoolchild) getAge() int {
	return this.age
}

func (this *schoolchild) getName() string {
	return this.name
}

func SPrint(input interface{}) string {
	/*
	input是不定接口类型（即任意接口类型），input.(type)返回的就是特定接口类型
	例如，如果用input指向一个student接口类型，那么通过input是无法调到student接口函数的（因为interface{}实际上弱化了student接口的功能）
	这时候如果进行input.(type)，则返回的就是student类型的接口了（相当于将interface{}强化还原成student）
	*/
	switch  type_flg := input.(type) {
	case student:
		return type_flg.getName()
	case string:
		return type_flg
	case int:
		return strconv.Itoa(type_flg)
	default:
		return "???"
	}
}

func TestReflect(t *testing.T) {
	var yong student
	yong = &schoolchild{"yuyong", 10, "usc"}
	fmt.Println(SPrint(yong))

	//不定接口转特定接口不能用强转
	var test_ref interface{} = yong
	fmt.Println(test_ref.(student).getName())

	val := reflect.ValueOf(test_ref)
	//防止出现reflect: call of reflect.Value.NumField on ptr Value [recovered]
	val = val.Elem()
	for i := 0; i < val.NumField(); i++ {
		field_name := val.Type().Field(i).Name
		fmt.Print(field_name + " ")
		fmt.Print(val.Type().Field(i).Type.Name() + " ")
		fmt.Println(val.FieldByName(field_name).String())
	}
}
