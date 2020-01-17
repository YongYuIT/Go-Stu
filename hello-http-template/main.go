package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	temp, err := template.ParseFiles("./view/fucktest.html")
	if err != nil {
		fmt.Println(err)
	}
	temp.Execute(os.Stdout, "fuck hhhhhhhhhhhhello")
}
