package main

import (
	"fmt"
	"gs/execute"
	"gs/setting"
	"gs/utils/color"
	"gs/utils/conf"
	"os"
)

const path = "/etc/goshell/shell.conf"

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		fmt.Println(color.Red("Please enter the parameters."))
		return
	}

	myConfig := new(conf.Config)

	if args[1] == "help" || args[1] == "--help" {
		myConfig.InitConfig(path)
		setting.Help(myConfig.MyNode)
	} else if args[1] == "init" || args[1] == "--init" {
		setting.Init(path)
	} else if args[1] == "update" || args[1] == "--update" {
		fmt.Println("update")
	} else if args[1] == "check" || args[1] == "--check" {
		fmt.Println("check")
	} else {
		myConfig.InitConfig(path)
		fmt.Println(args)

		if args == nil || len(args) > 2 {
			if args[2] == "--help" {
				shell := myConfig.Read(args[1], "command")
				introduce := myConfig.Read(args[1], "introduce")
				fmt.Println("")
				fmt.Println(color.Yellow(" Shell: "), color.White(shell))
				fmt.Println(color.Yellow(" Introduce: "), color.White(introduce))
				fmt.Println("")
				return
			}
		}

		shell := myConfig.Read(args[1], "command")
		if shell != "" {
			execute.Execute(shell, args)
		} else {
			fmt.Println(color.Red("Parameters do not exist."))
		}
	}
}
