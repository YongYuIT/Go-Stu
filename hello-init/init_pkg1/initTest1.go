package init_pkg1

import "fmt"
import "../init_pkg2"

func init() {
	fmt.Println("doInit111")
}

func Do111() {
	fmt.Println("hello 111")
	init_pkg2.Do222()
}
