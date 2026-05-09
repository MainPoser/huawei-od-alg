package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	scanner.Scan()
	targetMon, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	mans := strings.Fields(scanner.Text())
	scanner.Scan()
	birthDayStrs := strings.Fields(scanner.Text())

	dumManBD := make(map[string]string)

	for idx, man := range mans {
		dumManBD[man] = birthDayStrs[idx]
	}
	count := 0
	for _, bir := range dumManBD {
		birInt, _ := strconv.Atoi(strings.SplitN(bir, "/", 3)[1])
		if birInt == targetMon {
			count++
		}
	}
	fmt.Println(count)
}
