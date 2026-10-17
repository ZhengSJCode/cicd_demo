package main

import (
	"flag"
	"fmt"
)

// version 由构建时通过 -ldflags "-X main.version=..." 注入，默认 dev。
var version = "dev"

// greet 根据名字返回问候语，空名字时返回 Hello, World!。
func greet(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return "Hello, " + name + "!"
}

func main() {
	showVersion := flag.Bool("version", false, "打印版本号并退出")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	name := ""
	if flag.NArg() > 0 {
		name = flag.Arg(0)
	}
	fmt.Println(greet(name))
}
