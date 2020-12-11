package main

import (
	"github.com/TarsCloud/TarsGo/tars"
	"hello_tars/TestApp"
)

type HelloImp struct {
}

//implete the Test interface
func (imp *HelloImp) Test() (int32, error) {
	return 0, nil
}

//implete the testHello interface

func (imp *HelloImp) TestHello(in string, out *string) (int32, error) {
	*out = in
	return 0, nil
}

func main() { //Init servant
	imp := new(HelloImp)                                    //New Imp
	app := new(TestApp.Hello)                               //New init the A Tars
	cfg := tars.GetServerConfig()                           //Get Config File Object
	app.AddServant(imp, cfg.App+"."+cfg.Server+".HelloObj") //Register Servant
	tars.Run()
}
