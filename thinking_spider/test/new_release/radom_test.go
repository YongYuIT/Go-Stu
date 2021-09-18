package new_release

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_get_radom(test *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(rand.Intn(3))
	}
}
