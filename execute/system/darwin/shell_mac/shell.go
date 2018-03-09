package shell_mac

import (
	"os/exec"
	"bytes"
	"gs/utils/error"
)

func Exec_shell(s string) string {
	cmd := exec.Command("sh", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	error.Check(err, "")

	return out.String()
}
