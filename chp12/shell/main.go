package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "echo test>test.txt")
	} else {
		cmd = exec.Command("/bin/sh", "-c", "echo test>test.txt")
	}

	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

}
