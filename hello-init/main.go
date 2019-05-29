package main

import "./init_pkg"
import "./init_pkg1"

func main() {
	//init_pkg包里面有两个init函数
	//实验证明这连个init函数是随机执行的
	init_pkg.DoAAA()
	init_pkg.DoBBB()
	//main包引入init_pkg1，init_pkg1引入init_pkg2
	//实验证明，init_pkg1先执行，init_pkg2后执行
	init_pkg1.Do111()
}
