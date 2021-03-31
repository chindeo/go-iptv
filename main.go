package main

import (
	"fmt"
	"os"
	"time"

	"github.com/snowlyg/go-iptv/sys"
)

// /www/AvProxy-x64 -HttpPort 8088 -Log ./log -ConfigUrl http://127.0.0.1/test/AvProxy.xml http://127.0.0.1/static/AvProxy.xml
func main() {
	cmdName := "./AvProxy-x64"
	args := []string{"-HttpPort", "8088", "-Log", "./log", "-ConfigUrl", "http://127.0.0.1/test/AvProxy.xml"}
	// cmdName := "ping"
	// args := []string{"-t", "www.baidu.com"}
	go func() {
		err := sys.CmdRun(cmdName, args)
		if err != nil {
			panic(err)
		}
		fmt.Printf("1- %+v\n", args)
	}()

	time.Sleep(2 * time.Second)

	go func() {
		for sys.IsRunning(cmdName) {
			if sys.IsRunning(cmdName) {
				sys.Stop(cmdName)
				os.Remove("./AvProxy.xml.tmp")
				os.Remove("./AvProxy.xml")
			}
		}
	}()

	for {
		if !sys.IsRunning(cmdName) {
			args = []string{"-HttpPort", "8088", "-Log", "./log", "-ConfigUrl", "http://127.0.0.1/static/AvProxy.xml"}
			// args = []string{"-t", "www.qq.com"}
			err := sys.CmdRun(cmdName, args)
			if err != nil {
				panic(err)
			}
			fmt.Printf("2-%+v\n", args)
		}
	}

}
