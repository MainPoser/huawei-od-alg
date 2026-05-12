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
	scanner.Scan()
	items := strings.Fields(scanner.Text())
	openStr := items[0]
	shutdownStr := items[1]

	openSplit := strings.SplitN(openStr, ":", 2)
	openHour, _ := strconv.Atoi(openSplit[0])
	openMin, _ := strconv.Atoi(openSplit[1])
	// 开机时间点，距离00:00多少分钟
	openPoint := openHour*60 + openMin

	shutdownSplit := strings.SplitN(shutdownStr, ":", 2)
	shutdownHour, _ := strconv.Atoi(shutdownSplit[0])
	shutdownMin, _ := strconv.Atoi(shutdownSplit[1])
	// 关机时间点，距离00:00多少分钟
	shutdownPoint := shutdownHour*60 + shutdownMin

	total := shutdownPoint - openPoint
	if total == 0 {
		fmt.Println(0)
		return
	}
	stage1 := 0
	stage1Point := [][]int{{12 * 60, 13*60 + 30}, {17*60 + 30, 18 * 60}}
	for _, stage := range stage1Point {
		stage1 = stage1 + chongdie(stage[0], stage[1], openPoint, shutdownPoint)
	}
	// 阶段二
	// 00:00-12:00
	// 13:30-17:30
	// 18:00-24:00
	stage2 := 0
	limit := 10 * 60
	stage2Point := [][]int{{0, 12 * 60}, {13*60 + 30, 17*60 + 30}, {18 * 60, 24 * 60}}
	for _, stage := range stage2Point {
		cur := stage2 + chongdie(stage[0], stage[1], openPoint, shutdownPoint)
		remain := limit - stage2
		if remain <= 0 {
			break
		}
		if remain < cur {
			stage2 = stage2 + remain
		} else {
			stage2 = stage2 + cur
		}
	}
	stage3 := total - stage1 - stage2

	fmt.Printf("%v %v %v", stage1, stage2, stage3)

}

func chongdie(as, ae, bs, be int) int {
	start := as
	if as < bs {
		start = bs
	}
	end := ae
	if ae > be {
		end = be
	}
	times := end - start
	if times > 0 {
		return times
	} else {
		return 0
	}
}
