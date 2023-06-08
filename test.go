package main

import (
	"fmt"
	"os/exec"
)

func main() {
	ls := exec.Command("ls", "-a", "..", "-l")
	by, err := ls.Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(by))
}
