package main

import (
	"context"
)

// TestHelloSvanImp servant implementation
type TestHelloSvanImp struct {
}

// Init servant init
func (imp *TestHelloSvanImp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destory
func (imp *TestHelloSvanImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *TestHelloSvanImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *TestHelloSvanImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
