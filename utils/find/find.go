package find

import (
	"fmt"
	"regexp"
	"strings"
	"gs/utils/color"
)

func Args(shell string, args []string) (result string) {

	reg := regexp.MustCompile(`{{\w*}}`)
	arg_list := reg.FindAllString(shell, -1)
	if (len(args)-2 == len(arg_list)) {
		for i := 0; i < len(arg_list); i++ {
			shell = strings.Replace(shell, arg_list[i], args[i+2], -1)
		}
		return shell
	} else {
		fmt.Println(color.Red("Please enter the correct parameters."))
		fmt.Println(color.Red(shell))
	}
	return "null"
}
