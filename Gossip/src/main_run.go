package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"./gRPC"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type conf struct {
	Anchor_node string `yaml:"anchor_node"`
	Interval    int    `yaml:"interval"`
}

//export self_add=peer1.yong.thinking.com
//go run main_run.go -conf ./conf.ymal

func main() {
	//read config
	conf_file_path := flag.String("conf", "./conf.ymal", "conf file path")
	flag.Parse()
	fmt.Println(*conf_file_path)
	yamlFile, err := ioutil.ReadFile(*conf_file_path)
	if err != nil {
		fmt.Println(err.Error())
	}

	var c *conf = new(conf)
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}

	self_add := os.Getenv("self_add")

	if self_add == "" {
		fmt.Println("env error")
		return
	}

	fmt.Println(c.Anchor_node)
	fmt.Println(strconv.Itoa(c.Interval))
	fmt.Println(self_add)

	//start listen
	gRPC.Self_add = self_add
	gRPC.Anchor_add = c.Anchor_node
	go gRPC.Listen(c.Interval)

	//stop
	for {
		inputReader := bufio.NewReader(os.Stdin)
		str, err := inputReader.ReadString('\n')
		if err == nil {
			if strings.Compare(str, "stop\n") == 0 {
				break
			}
		}
	}
	//kill coroutines
	gRPC.IsOut = true
}
