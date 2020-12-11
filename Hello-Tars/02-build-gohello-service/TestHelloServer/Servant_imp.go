package main

import (
	"context"
)

// TestHelloServantImp servant implementation
type TestHelloServantImp struct {
}

// Init servant init
func (imp *TestHelloServantImp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destory
func (imp *TestHelloServantImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *TestHelloServantImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	*c = a + b
	return 0, nil
}
func (imp *TestHelloServantImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	*c = a - b
	return 0, nil
}
