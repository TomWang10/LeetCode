package main

import (
	"fmt"
)

/**
	1. 初始化结果棋盘
	2.
 */
func solveNQueens(n int) (result [][]string) {
	// 初始化
	chesstable := make([]string, n)
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			chesstable[row] = chesstable[row] + "."
		}
	}
	for i := 0; i < n; i++ {
		result = append(result, chesstable)
	}
}

func main() {
	chesstable := make([]string, 5)
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			chesstable[row] = chesstable[row] + "."
		}
	}
	fmt.Println(chesstable)
}
