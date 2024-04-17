package main

import (
	"12_keychain/internal/command"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: keychain <command>")
		fmt.Println("Available commands:")
		fmt.Println("list - list all passwords in keychain")
		fmt.Println("save <password name> - save a password")
		fmt.Println("get <password name> - get a password")
		return
	}

	arg := flag.Arg(0)
	cmd := flag.NewFlagSet(arg, flag.ExitOnError)

	c, err := command.Create(arg, cmd)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := c.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
