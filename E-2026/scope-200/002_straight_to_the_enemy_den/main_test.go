/*
 * File: main_test.go
 * Project: 002_straight_to_the_enemy_den
 * File Created: Monday, 11th May 2026 9:25:45 am
 * Author: tianyao (ty18710388929@163.com)
 * -----
 * Last Modified: Monday, 11th May 2026 9:26:11 am
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
