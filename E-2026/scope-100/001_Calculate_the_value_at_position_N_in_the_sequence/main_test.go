/*
 * File: main_test.go
 * Project: 001_Calculate_the_value_at_position_N_in_the_sequence
 * File Created: Saturday, 9th May 2026 2:47:40 pm
 * Author: tianyao (ty18710388929@163.com)
 * -----
 * Last Modified: Saturday, 9th May 2026 2:51:19 pm
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
