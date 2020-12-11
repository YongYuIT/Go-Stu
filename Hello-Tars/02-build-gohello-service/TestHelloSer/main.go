package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"TestHelloMod/tars-protocol/TestHelloApp1"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(TestHelloSvanImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("TestHelloSvanImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(TestHelloApp1.TestHelloSvan)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".TestHelloSvanObj")

	// Run application
	tars.Run()
}
