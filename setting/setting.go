package setting

import (
	"fmt"
	"gs/execute"
	"gs/utils/color"
	"gs/utils/file"
	"strings"
)

func Init(path string) {
	result := "[ls]\n" +
		"command = ls -all\n" +
		"introduce = ls -all 查看当前目录文件\n\n" +
		"" +
		"[ps]\n" +
		"command = ps -ef\n" +
		"introduce = ps -ef 查看当前进程\n\n" +
		"" +
		"[ll]\n" +
		"command = ll -h\n" +
		"introduce = ll -h 查看进程\n\n" +
		"" +
		"[kall]\n" +
		"command = ps -ef | grep {{name}} | awk '{print \"kill -9 \"$2}' | sh\n" +
		"introduce = introduce = Kill 全部相关进程，自定义参数 name 为服务名\n\n" +
		"" +
		"[pss]\n" +
		"command = ps {{key}}\n" +
		"introduce = ps 自定义参数\n\n" +
		""

	file.Output(result, path)

	fmt.Println(color.Blue("goshell setting success."))
}

func Help(list map[string]string) {
	execute.Execute_shell("clear")
	logo := `    ______     _____ __         ____
   / ____/___ / ___// /_  ___  / / /
  / / __/ __ \\__ \/ __ \/ _ \/ / /
 / /_/ / /_/ /__/ / / / /  __/ / /
 \____/\____/____/_/ /_/\___/_/_/`
	fmt.Print(color.Yellow(logo))
	fmt.Println(color.White("   V0.1\n"))
	fmt.Println(color.White(" A tool for quick execution of shell."))
	fmt.Println("")
	fmt.Println(color.Yellow(" +[ ABOUT ]-----------------------------------------------------+"))
	fmt.Println("")
	fmt.Println(color.Green("   + Autor   : "), color.White("lauixData"))
	fmt.Println(color.Green("   + Home    : "), color.White("https://GoShell.io"))
	fmt.Println(color.Green("   + Github  : "), color.White("https://github.com/lauixData/GoShell"))
	fmt.Println("")
	fmt.Println(color.Yellow(" +[ DEFAULT ARGUMENTS ]-----------------------------------------+"))
	fmt.Println("")
	fmt.Println(color.Cyan("   init,--init"), color.White("		Goshell Initialization"))
	fmt.Println(color.Cyan("   update,--update"), color.White("	Conf Update"))
	fmt.Println(color.Cyan("   check,--check"), color.White("	Conf Check"))
	fmt.Println(color.Cyan("   help,--help"), color.White("		Goshell Help"))
	fmt.Println("")
	fmt.Println(color.Yellow(" +[ CUSTOM ARGUMENTS ]------------------------------------------+"))
	fmt.Println("")
	for k, v := range list {
		arr := strings.Split(v, "&&")

		if arr[0] == "" {
			fmt.Println("  ", color.Cyan(k), "	", color.White(arr[1]))
		} else {
			fmt.Println("  ", color.Cyan(k), "	", color.White(arr[1]))
		}
	}
	fmt.Println("")
}
