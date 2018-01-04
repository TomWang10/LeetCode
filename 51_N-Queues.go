package main

import (
	"fmt"
	"math"
)

func CanPlace(table *[]int, row int, col int) (can bool) {
	// 列重复
	// 对角线重复，两个位置的坐标相减绝对值  [row - i] == [col - table[i]]
	for k, v := range *table {
		if v == col {
			return  false
		}
		if math.Abs(float64(row - k)) == math.Abs(float64(col - v)) {
			return false
		}
	}
	return true
}

func QueensHelper(result *[][]string, table *[]int, row int) {
	if row == len(*table) {
		// 添加到结果
		// 生成棋盘
		chessTable := make([]string, row)
		for i := 0; i < row; i++ {
			for j := 0; j < row; j++ {
				if (*table)[i] == j {
					chessTable[i] = chessTable[i] + "Q"
				} else {
					chessTable[i] = chessTable[i] + "."
				}
			}
		}
		*result = append(*result, chessTable)
		return
	}
	for col := 0; col < len(*table); col++ {
		if CanPlace(table, row, col) {
			// 放入
			(*table)[row] = col
			// 下一行
			QueensHelper(result, table, row+1)
			(*table)[row] = -10000
			}
		}
}

func solveNQueens(n int) (result [][]string) {
	table := make([]int, n)
	for i := 0; i < len(table); i++ {
		table[i] = -10000
	}
	QueensHelper(&result, &table, 0)
	return result
}

func main() {
	result := solveNQueens(4)
	for _, k := range result {
		for _, k2 := range k {
			fmt.Println(k2)
		}
	}
}
