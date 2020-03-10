package main

import (
	"bufio"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
)

func main() {
	wacher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("new wacher failed-->", err)
		return
	}
	defer wacher.Close()

	err = wacher.Add("/home/yong/Go-Stu20200302001/ETLLogMonitoringService/test_path/")
	if err != nil {
		fmt.Println("add path failed-->", err)
		return
	}

	go func() {
		for {
			select {
			case event := <-wacher.Events:
				fmt.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Println("Create file:", event.Name)
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("Write file:", event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					fmt.Println("Remove file:", event.Name)
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					fmt.Println("Rename file:", event.Name)
				}
				if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					fmt.Println("Chmod file:", event.Name)
				}
			case err = <-wacher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()

	inputReader := bufio.NewReader(os.Stdin)
	inputReader.ReadString('\n')
}
