/*
 * File: main_test.go
 * Project: 002_Space_Occupancy_Calculation
 * File Created: Saturday, 9th May 2026 3:06:52 pm
 * Author: tianyao (ty18710388929@163.com)
 * -----
 * Last Modified: Saturday, 9th May 2026 5:09:53 pm
 * Modified By: tianyao (ty18710388929@163.com>)
 * -----
 * Copyright <<projectCreationYear>> - 2026 tianyao, tianyao
 */

package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	// 仅在本地开发调试时使用
	if _, err := os.Stat("input.txt"); err == nil {
		f, _ := os.Open("input.txt")
		os.Stdin = f // 将标准输入重定向到文件
		defer f.Close()
	}
	tests := []struct {
		name string // description of this test case
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
