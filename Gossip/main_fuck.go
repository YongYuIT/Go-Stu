package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strconv"
	"./gRPC"
	"bufio"
	"os"
	"strings"
)

type conf struct {
	Anchor_node string `yaml:"anchor_node"`
	Interval    int    `yaml:"interval"`
	Self_add    string `yaml:"self_add"`
}

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

	fmt.Println(c.Anchor_node)
	fmt.Println(strconv.Itoa(c.Interval))

	//start listen
	gRPC.Self_add = c.Self_add
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
