package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"TestHelloApp/TestHelloServer/TestHelloApp"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(TestHelloServantImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("TestHelloServantImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(TestHelloApp.TestHelloServant)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".TestHelloServantObj")

	// Run application
	tars.Run()
}
