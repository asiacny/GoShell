package error

import (
	"fmt"
	"gs/utils/color"
)

func Check(e error, tips string) {
	if e != nil {
		panic(e)
		fmt.Println(color.Red(tips))
	}
}
