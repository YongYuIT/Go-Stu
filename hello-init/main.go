package main

import "./init_pkg"

func main() {
	//init_pkg包里面有两个init函数
	//实验证明这连个init函数是随机执行的
	init_pkg.DoAAA()
	init_pkg.DoBBB()
}
