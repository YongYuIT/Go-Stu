package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestTimeOut(test *testing.T) {
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)
}
