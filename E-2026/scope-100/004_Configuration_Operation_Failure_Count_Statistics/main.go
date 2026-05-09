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
	reqLine := scanner.Text()
	commands := make([]string, 0)
	command := ""
	for _, c := range []rune(reqLine) {
		if c == ']' {
			commands = append(commands, command)
			command = ""
		} else if c != '[' {
			command = command + string(c)
		}
	}
	data := make(map[int]int)
	faildCount := 0
	for _, command := range commands {
		fmt.Println(command)
		fields := strings.Fields(command)
		op := fields[0]
		if op != "add_rule" && op != "mod_rule" && op != "del_rule" {
			faildCount++
			continue
		}

		infraData := make(map[string]int, 2)
		for i := 1; i < len(fields); i++ {
			as := strings.SplitN(fields[i], "=", 2)
			bs := strings.SplitN(fields[i], "=", 2)
			aVal, _ := strconv.Atoi(as[1])
			infraData[as[0]] = aVal
			bVal, _ := strconv.Atoi(bs[1])
			infraData[bs[0]] = bVal
		}

		if op == "add_rule" {
			if infraData["rule_id"] < 1 || infraData["rule_id"] > 9999 || infraData["rule_index"] < 1 || infraData["rule_index"] > 9999 || data[infraData["rule_id"]] != 0 {
				faildCount++
				continue
			}
			data[infraData["rule_id"]] = infraData["rule_index"]
		} else if op == "mod_rule" {
			if infraData["rule_id"] < 1 || infraData["rule_id"] > 9999 || infraData["rule_index"] < 1 || infraData["rule_index"] > 9999 || data[infraData["rule_id"]] == 0 {
				faildCount++
				continue
			}
			data[infraData["rule_id"]] = infraData["rule_index"]
		} else {
			if infraData["rule_id"] < 1 || infraData["rule_id"] > 9999 || data[infraData["rule_id"]] == 0 {
				faildCount++
				continue
			}
			delete(data, infraData["rule_id"])
		}
	}

	fmt.Println(faildCount)

}
