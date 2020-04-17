package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var cekip, _ = exec.Command("ifconfig").Output()
	var pwd, _ = exec.Command("pwd").Output()
	var ls, _ = exec.Command("ls").Output()
	fmt.Println(string(cekip))
	fmt.Println(string(pwd))
	fmt.Println(string(ls))
}
