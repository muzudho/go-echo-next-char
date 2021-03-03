package main

import (
	"bufio"
	"os"
	"strings"
)

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
var next = "BCDEFGHIJKLMNOPQRSTUVWXYZAbcdefghijklmnopqrstuvwxyza"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 標準入力を読込みます
	for scanner.Scan() {
		command := scanner.Text()

		if 0 < len(command) {
			one := command[0:1]

			index := strings.Index(alphabet, one)
			if index != -1 {
				next := next[index : index+1]
				os.Stdout.WriteString(next)
			}
		}
	}
}
