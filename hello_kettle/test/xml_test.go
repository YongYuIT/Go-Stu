package test

import (
	"fmt"
	"github.com/beevik/etree"
	"testing"
)

func TestXmlRemoveEle(test *testing.T) {
	doc := etree.NewDocument()
	err := doc.ReadFromFile("test1.xml")
	if err != nil {
		fmt.Println("read doc err-->", err)
		return
	}
	rootEle := doc.FindElement("aaa")
	fmt.Println(rootEle.Child)
	fmt.Println(rootEle.ChildElements())
	rootEle.RemoveChild(rootEle.ChildElements()[0])
	doc.WriteToFile("test2.xml")
}
