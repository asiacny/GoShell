package execute

import (
	"gs/utils/find"
	"fmt"
	"runtime"
	"gs/execute/system/darwin/shell_mac"
	"gs/execute/system/linux/shell_linux"
	"gs/execute/system/windows/cmd"
)

func Execute(shell string, args []string) {
	shell = find.Args(shell, args)
	if (shell != "null") {
		os := runtime.GOOS

		switch {
		case os == "darwin":
			result := shell_mac.Exec_shell(shell)
			fmt.Println(result)
		case os == "linux":
			result := shell_linux.Exec_shell(shell)
			fmt.Println(result)
		case os == "windows":
			result := cmd.Exec_shell(shell)
			fmt.Println(result)
		}
	}
}

func Execute_shell(shell string) {
	os := runtime.GOOS

	switch {
	case os == "darwin":
		result := shell_mac.Exec_shell(shell)
		fmt.Println(result)
	case os == "linux":
		result := shell_linux.Exec_shell(shell)
		fmt.Println(result)
	case os == "windows":
		result := cmd.Exec_shell(shell)
		fmt.Println(result)
	}
}
